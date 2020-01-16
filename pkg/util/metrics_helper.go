package util

import (
	dto "github.com/prometheus/client_model/go"
)

// MetricFamiliesPerUser is a collection of metrics gathered via calling Gatherer.Gather() method on different
// gatherers, one per user.
// First key = userID, second key = metric name.
// Value = slice of gathered values with the same metric name.
type MetricFamiliesPerUser map[string]map[string][]*dto.MetricFamily

func NewMetricFamiliersPerUser() MetricFamiliesPerUser {
	return MetricFamiliesPerUser{}
}

// AddGatheredDataForUser adds user-specific output of Gatherer.Gather method.
func (d MetricFamiliesPerUser) AddGatheredDataForUser(userID string, metrics []*dto.MetricFamily) {
	// first, create new map which maps metric names to a slice of MetricFamily instances.
	// That makes it easier to do searches later.
	perMetricName := map[string][]*dto.MetricFamily{}

	for _, m := range metrics {
		if m.Name == nil {
			continue
		}
		perMetricName[*m.Name] = append(perMetricName[*m.Name], m)
	}

	d[userID] = perMetricName
}

// SumCountersAcrossAllUsers returns sum(counter).
func (d MetricFamiliesPerUser) SumCountersAcrossAllUsers(counter string) float64 {
	result := float64(0)
	for _, perMetric := range d {
		result += sumCounters(perMetric[counter])
	}
	return result
}

// SumCountersPerUser returns sum(counter) by (userID), where userID will be the map key.
func (d MetricFamiliesPerUser) SumCountersPerUser(counter string) map[string]float64 {
	result := map[string]float64{}
	for user, perMetric := range d {
		v := sumCounters(perMetric[counter])
		result[user] = v
	}
	return result
}

func sumCounters(mfs []*dto.MetricFamily) float64 {
	result := float64(0)
	for _, mf := range mfs {
		for _, m := range mf.Metric {
			// This works even if m is nil, m.Counter is nil or m.Counter.Value is nil (it returns 0 in those cases)
			result += m.GetCounter().GetValue()
		}
	}
	return result
}
