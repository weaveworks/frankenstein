package ingester

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRate(t *testing.T) {
	ticks := []struct {
		events int
		want   float64
	}{
		{60, 1},
		{30, 0.9},
		{0, 0.72},
		{60, 0.776},
		{0, 0.6208},
		{0, 0.49664},
		{0, 0.397312},
		{0, 0.3178496},
		{0, 0.25427968},
		{0, 0.203423744},
		{0, 0.1627389952},
	}
	r := newEWMARate(0.2, time.Minute)

	for i, tick := range ticks {
		for e := 0; e < tick.events; e++ {
			r.inc()
		}
		r.tick()
		// We cannot do double comparison, because double operations on different
		// platforms may actually produce results that differ slightly.
		// There are multiple issues about this in Go's github, eg: 18354 or 20319.
		require.InDelta(t, tick.want, r.rate(), 0.0000000001, "unexpected rate %d", i)
	}

	r = newEWMARate(0.2, time.Minute)

	for i, tick := range ticks {
		r.add(int64(tick.events))
		r.tick()
		if r.rate() != tick.want {
			t.Fatalf("%d. unexpected rate: want %v, got %v", i, tick.want, r.rate())
		}
	}
}
