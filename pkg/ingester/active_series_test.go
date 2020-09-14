package ingester

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cortexproject/cortex/pkg/ingester/client"
)

func copyFn(l labels.Labels) labels.Labels { return l }

func TestActiveSeries_UpdateSeries(t *testing.T) {
	ls1 := []labels.Label{{Name: "a", Value: "1"}}
	ls2 := []labels.Label{{Name: "a", Value: "2"}}

	c := NewActiveSeries()
	assert.Equal(t, 0, c.Active())

	c.UpdateSeries(ls1, time.Now(), copyFn)
	assert.Equal(t, 1, c.Active())

	c.UpdateSeries(ls1, time.Now(), copyFn)
	assert.Equal(t, 1, c.Active())

	c.UpdateSeries(ls2, time.Now(), copyFn)
	assert.Equal(t, 2, c.Active())
}

func TestActiveSeries_ShouldCorrectlyHandleFingerprintCollisions(t *testing.T) {
	metric := labels.NewBuilder(labels.FromStrings("__name__", "logs"))
	ls1 := metric.Set("_", "ypfajYg2lsv").Labels()
	ls2 := metric.Set("_", "KiqbryhzUpn").Labels()

	require.True(t, client.Fingerprint(ls1) == client.Fingerprint(ls2))

	c := NewActiveSeries()
	c.UpdateSeries(ls1, time.Now(), copyFn)
	c.UpdateSeries(ls2, time.Now(), copyFn)

	assert.Equal(t, 2, c.Active())
}

func TestActiveSeries_Purge(t *testing.T) {
	series := [][]labels.Label{
		{{Name: "a", Value: "1"}},
		{{Name: "a", Value: "2"}},
		// The two following series have the same Fingerprint
		{{Name: "_", Value: "ypfajYg2lsv"}, {Name: "__name__", Value: "logs"}},
		{{Name: "_", Value: "KiqbryhzUpn"}, {Name: "__name__", Value: "logs"}},
	}

	// Run the same test for increasing TTL values
	for ttl := 0; ttl < len(series); ttl++ {
		c := NewActiveSeries()

		for i := 0; i < len(series); i++ {
			c.UpdateSeries(series[i], time.Unix(int64(i), 0), copyFn)
		}

		c.Purge(time.Unix(int64(ttl+1), 0))
		// call purge twice, just to hit "quick" path. It doesn't really do anything.
		c.Purge(time.Unix(int64(ttl+1), 0))

		exp := len(series) - (ttl + 1)
		assert.Equal(t, exp, c.Active())
	}
}

func TestActiveSeries_PurgeOpt(t *testing.T) {
	metric := labels.NewBuilder(labels.FromStrings("__name__", "logs"))
	ls1 := metric.Set("_", "ypfajYg2lsv").Labels()
	ls2 := metric.Set("_", "KiqbryhzUpn").Labels()

	c := NewActiveSeries()

	now := time.Now()
	c.UpdateSeries(ls1, now.Add(-2*time.Minute), copyFn)
	c.UpdateSeries(ls2, now, copyFn)
	c.Purge(now)

	assert.Equal(t, 1, c.Active())

	c.UpdateSeries(ls1, now.Add(-1*time.Minute), copyFn)
	c.UpdateSeries(ls2, now, copyFn)
	c.Purge(now)

	assert.Equal(t, 1, c.Active())

	// This will *not* update the series, since there is already newer timestamp.
	c.UpdateSeries(ls2, now.Add(-1*time.Minute), copyFn)
	c.Purge(now)

	assert.Equal(t, 1, c.Active())
}

var activeSeriesTestGoroutines = []int{50, 100, 500}

func BenchmarkActiveSeriesTest_single_label(b *testing.B) {
	const seriesCount = 1e5

	series := make([]labels.Labels, seriesCount)

	for s := 0; s < len(series); s++ {
		series[s] = labels.Labels{
			{Name: "a", Value: strconv.Itoa(s)},
		}
	}

	for _, num := range activeSeriesTestGoroutines {
		b.Run(fmt.Sprintf("%d", num), func(b *testing.B) {
			benchmarkActiveSeriesConcurrency(b, series, num)
		})
	}
}

func benchmarkActiveSeriesConcurrency(b *testing.B, series []labels.Labels, goroutines int) {
	c := NewActiveSeries()

	r := rand.New(rand.NewSource(123456789))
	ch := make(chan int, b.N)
	for i := 0; i < b.N; i++ {
		ch <- r.Intn(len(series))
	}
	close(ch)

	wg := &sync.WaitGroup{}
	start := make(chan struct{})

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start

			now := time.Now()

			for ix := range ch {
				now = now.Add(time.Duration(ix) * time.Millisecond)

				c.UpdateSeries(series[ix], now, copyFn)
			}
		}()
	}

	b.ResetTimer()
	close(start)
	wg.Wait()
}

func BenchmarkActiveSeries_UpdateSeries(b *testing.B) {
	const numSeries = 10000

	c := NewActiveSeries()

	// Prepare series
	series := [numSeries]labels.Labels{}
	for s := 0; s < numSeries; s++ {
		series[s] = labels.Labels{{Name: "a", Value: strconv.Itoa(s)}}
	}

	now := time.Now()

	b.ResetTimer()
	for i := 0; i < int(math.Ceil(float64(b.N)/numSeries)); i++ {
		for s := 0; s < numSeries; s++ {
			c.UpdateSeries(series[s], now, copyFn)
		}
	}
}

func BenchmarkActiveSeries_Purge_once(b *testing.B) {
	benchmarkPurge(b, false)
}

func BenchmarkActiveSeries_Purge_twice(b *testing.B) {
	benchmarkPurge(b, true)
}

func benchmarkPurge(b *testing.B, twice bool) {
	const numSeries = 10000
	const numExpiresSeries = numSeries / 25

	now := time.Now()
	c := NewActiveSeries()

	series := [numSeries]labels.Labels{}
	for s := 0; s < numSeries; s++ {
		series[s] = labels.Labels{{Name: "a", Value: strconv.Itoa(s)}}
	}

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		// Prepare series
		for ix, s := range series {
			if ix < numExpiresSeries {
				c.UpdateSeries(s, now.Add(-time.Minute), copyFn)
			} else {
				c.UpdateSeries(s, now, copyFn)
			}
		}

		assert.Equal(b, numSeries, c.Active())
		b.StartTimer()

		// Purge everything
		c.Purge(now)
		assert.Equal(b, numSeries-numExpiresSeries, c.Active())

		if twice {
			c.Purge(now)
			assert.Equal(b, numSeries-numExpiresSeries, c.Active())
		}
	}
}
