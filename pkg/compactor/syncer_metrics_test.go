package compactor

import (
	"bytes"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

func TestSyncerMetrics(t *testing.T) {
	reg := prometheus.NewPedanticRegistry()

	sm := newSyncerMetrics(reg)
	sm.gatherThanosSyncerMetrics(generateTestData(12345))
	sm.gatherThanosSyncerMetrics(generateTestData(76543))
	sm.gatherThanosSyncerMetrics(generateTestData(22222))
	// total base = 111110

	err := testutil.GatherAndCompare(reg, bytes.NewBufferString(`
			# HELP cortex_compactor_sync_meta_total TSDB Syncer: Total number of sync meta operations.
			# TYPE cortex_compactor_sync_meta_total counter
			cortex_compactor_sync_meta_total 111110

			# HELP cortex_compactor_sync_meta_failures_total TSDB Syncer: Total number of failed sync meta operations.
			# TYPE cortex_compactor_sync_meta_failures_total counter
			cortex_compactor_sync_meta_failures_total 222220

			# HELP cortex_compactor_sync_meta_duration_seconds TSDB Syncer: Time it took to sync meta files.
			# TYPE cortex_compactor_sync_meta_duration_seconds histogram
			# Observed values: 3.7035, 22.9629, 6.6666 (seconds)
			cortex_compactor_sync_meta_duration_seconds_bucket{le="0.01"} 0
			cortex_compactor_sync_meta_duration_seconds_bucket{le="0.1"} 0
			cortex_compactor_sync_meta_duration_seconds_bucket{le="0.3"} 0
			cortex_compactor_sync_meta_duration_seconds_bucket{le="0.6"} 0
			cortex_compactor_sync_meta_duration_seconds_bucket{le="1"} 0
			cortex_compactor_sync_meta_duration_seconds_bucket{le="3"} 0
			cortex_compactor_sync_meta_duration_seconds_bucket{le="6"} 1
			cortex_compactor_sync_meta_duration_seconds_bucket{le="9"} 2
			cortex_compactor_sync_meta_duration_seconds_bucket{le="20"} 2
			cortex_compactor_sync_meta_duration_seconds_bucket{le="30"} 3
			cortex_compactor_sync_meta_duration_seconds_bucket{le="60"} 3
			cortex_compactor_sync_meta_duration_seconds_bucket{le="90"} 3
			cortex_compactor_sync_meta_duration_seconds_bucket{le="120"} 3
			cortex_compactor_sync_meta_duration_seconds_bucket{le="240"} 3
			cortex_compactor_sync_meta_duration_seconds_bucket{le="360"} 3
			cortex_compactor_sync_meta_duration_seconds_bucket{le="720"} 3
			cortex_compactor_sync_meta_duration_seconds_bucket{le="+Inf"} 3
			# rounding error
			cortex_compactor_sync_meta_duration_seconds_sum 33.333000000000006
			cortex_compactor_sync_meta_duration_seconds_count 3

			# HELP cortex_compactor_garbage_collected_blocks_total TSDB Syncer: Total number of deleted blocks by compactor.
			# TYPE cortex_compactor_garbage_collected_blocks_total counter
			cortex_compactor_garbage_collected_blocks_total 444440

			# HELP cortex_compactor_garbage_collection_total TSDB Syncer: Total number of garbage collection operations.
			# TYPE cortex_compactor_garbage_collection_total counter
			cortex_compactor_garbage_collection_total 555550

			# HELP cortex_compactor_garbage_collection_failures_total TSDB Syncer: Total number of failed garbage collection operations.
			# TYPE cortex_compactor_garbage_collection_failures_total counter
			cortex_compactor_garbage_collection_failures_total 666660

			# HELP cortex_compactor_garbage_collection_duration_seconds TSDB Syncer: Time it took to perform garbage collection iteration.
			# TYPE cortex_compactor_garbage_collection_duration_seconds histogram
			# Observed values: 8.6415, 53.5801, 15.5554
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="0.01"} 0
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="0.1"} 0
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="0.3"} 0
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="0.6"} 0
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="1"} 0
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="3"} 0
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="6"} 0
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="9"} 1
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="20"} 2
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="30"} 2
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="60"} 3
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="90"} 3
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="120"} 3
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="240"} 3
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="360"} 3
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="720"} 3
			cortex_compactor_garbage_collection_duration_seconds_bucket{le="+Inf"} 3
			cortex_compactor_garbage_collection_duration_seconds_sum 77.777
			cortex_compactor_garbage_collection_duration_seconds_count 3

			# HELP cortex_compactor_group_compactions_total TSDB Syncer: Total number of group compaction attempts that resulted in a new block.
			# TYPE cortex_compactor_group_compactions_total counter
			cortex_compactor_group_compactions_total{group="aaa"} 888880
			cortex_compactor_group_compactions_total{group="bbb"} 999990
			cortex_compactor_group_compactions_total{group="ccc"} 1.1111e+06

			# HELP cortex_compactor_group_compaction_runs_started_total TSDB Syncer: Total number of group compaction attempts.
			# TYPE cortex_compactor_group_compaction_runs_started_total counter
			cortex_compactor_group_compaction_runs_started_total{group="aaa"} 1.22221e+06
			cortex_compactor_group_compaction_runs_started_total{group="bbb"} 1.33332e+06
			cortex_compactor_group_compaction_runs_started_total{group="ccc"} 1.44443e+06

			# HELP cortex_compactor_group_compaction_runs_completed_total TSDB Syncer: Total number of group completed compaction runs. This also includes compactor group runs that resulted with no compaction.
			# TYPE cortex_compactor_group_compaction_runs_completed_total counter
			cortex_compactor_group_compaction_runs_completed_total{group="aaa"} 1.55554e+06
			cortex_compactor_group_compaction_runs_completed_total{group="bbb"} 1.66665e+06
			cortex_compactor_group_compaction_runs_completed_total{group="ccc"} 1.77776e+06

			# HELP cortex_compactor_group_compactions_failures_total TSDB Syncer: Total number of failed group compactions.
			# TYPE cortex_compactor_group_compactions_failures_total counter
			cortex_compactor_group_compactions_failures_total{group="aaa"} 1.88887e+06
			cortex_compactor_group_compactions_failures_total{group="bbb"} 1.99998e+06
			cortex_compactor_group_compactions_failures_total{group="ccc"} 2.11109e+06

			# HELP cortex_compactor_group_vertical_compactions_total TSDB Syncer: Total number of group compaction attempts that resulted in a new block based on overlapping blocks.
			# TYPE cortex_compactor_group_vertical_compactions_total counter
			cortex_compactor_group_vertical_compactions_total{group="aaa"} 2.2222e+06
			cortex_compactor_group_vertical_compactions_total{group="bbb"} 2.33331e+06
			cortex_compactor_group_vertical_compactions_total{group="ccc"} 2.44442e+06
	`))
	require.NoError(t, err)
}

func generateTestData(base float64) *prometheus.Registry {
	r := prometheus.NewRegistry()
	m := newTesSyncerMetrics(r)
	m.syncMetas.Add(1 * base)
	m.syncMetaFailures.Add(2 * base)
	m.syncMetaDuration.Observe(3 * base / 10000)
	m.garbageCollectedBlocks.Add(4 * base)
	m.garbageCollections.Add(5 * base)
	m.garbageCollectionFailures.Add(6 * base)
	m.garbageCollectionDuration.Observe(7 * base / 10000)
	m.compactions.WithLabelValues("aaa").Add(8 * base)
	m.compactions.WithLabelValues("bbb").Add(9 * base)
	m.compactions.WithLabelValues("ccc").Add(10 * base)
	m.compactionRunsStarted.WithLabelValues("aaa").Add(11 * base)
	m.compactionRunsStarted.WithLabelValues("bbb").Add(12 * base)
	m.compactionRunsStarted.WithLabelValues("ccc").Add(13 * base)
	m.compactionRunsCompleted.WithLabelValues("aaa").Add(14 * base)
	m.compactionRunsCompleted.WithLabelValues("bbb").Add(15 * base)
	m.compactionRunsCompleted.WithLabelValues("ccc").Add(16 * base)
	m.compactionFailures.WithLabelValues("aaa").Add(17 * base)
	m.compactionFailures.WithLabelValues("bbb").Add(18 * base)
	m.compactionFailures.WithLabelValues("ccc").Add(19 * base)
	m.verticalCompactions.WithLabelValues("aaa").Add(20 * base)
	m.verticalCompactions.WithLabelValues("bbb").Add(21 * base)
	m.verticalCompactions.WithLabelValues("ccc").Add(22 * base)
	return r
}

// directly copied from Thanos (and renamed syncerMetrics to testSyncerMetrics to avoid conflict)
type testSyncerMetrics struct {
	syncMetas                 prometheus.Counter
	syncMetaFailures          prometheus.Counter
	syncMetaDuration          prometheus.Histogram
	garbageCollectedBlocks    prometheus.Counter
	garbageCollections        prometheus.Counter
	garbageCollectionFailures prometheus.Counter
	garbageCollectionDuration prometheus.Histogram
	compactions               *prometheus.CounterVec
	compactionRunsStarted     *prometheus.CounterVec
	compactionRunsCompleted   *prometheus.CounterVec
	compactionFailures        *prometheus.CounterVec
	verticalCompactions       *prometheus.CounterVec
}

func newTesSyncerMetrics(reg prometheus.Registerer) *testSyncerMetrics {
	var m testSyncerMetrics

	m.syncMetas = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "thanos_compact_sync_meta_total",
		Help: "Total number of sync meta operations.",
	})
	m.syncMetaFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "thanos_compact_sync_meta_failures_total",
		Help: "Total number of failed sync meta operations.",
	})
	m.syncMetaDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "thanos_compact_sync_meta_duration_seconds",
		Help:    "Time it took to sync meta files.",
		Buckets: []float64{0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120, 240, 360, 720},
	})

	m.garbageCollectedBlocks = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "thanos_compact_garbage_collected_blocks_total",
		Help: "Total number of deleted blocks by compactor.",
	})
	m.garbageCollections = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "thanos_compact_garbage_collection_total",
		Help: "Total number of garbage collection operations.",
	})
	m.garbageCollectionFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "thanos_compact_garbage_collection_failures_total",
		Help: "Total number of failed garbage collection operations.",
	})
	m.garbageCollectionDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "thanos_compact_garbage_collection_duration_seconds",
		Help:    "Time it took to perform garbage collection iteration.",
		Buckets: []float64{0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120, 240, 360, 720},
	})

	m.compactions = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "thanos_compact_group_compactions_total",
		Help: "Total number of group compaction attempts that resulted in a new block.",
	}, []string{"group"})
	m.compactionRunsStarted = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "thanos_compact_group_compaction_runs_started_total",
		Help: "Total number of group compaction attempts.",
	}, []string{"group"})
	m.compactionRunsCompleted = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "thanos_compact_group_compaction_runs_completed_total",
		Help: "Total number of group completed compaction runs. This also includes compactor group runs that resulted with no compaction.",
	}, []string{"group"})
	m.compactionFailures = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "thanos_compact_group_compactions_failures_total",
		Help: "Total number of failed group compactions.",
	}, []string{"group"})
	m.verticalCompactions = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "thanos_compact_group_vertical_compactions_total",
		Help: "Total number of group compaction attempts that resulted in a new block based on overlapping blocks.",
	}, []string{"group"})

	if reg != nil {
		reg.MustRegister(
			m.syncMetas,
			m.syncMetaFailures,
			m.syncMetaDuration,
			m.garbageCollectedBlocks,
			m.garbageCollections,
			m.garbageCollectionFailures,
			m.garbageCollectionDuration,
			m.compactions,
			m.compactionRunsStarted,
			m.compactionRunsCompleted,
			m.compactionFailures,
			m.verticalCompactions,
		)
	}
	return &m
}
