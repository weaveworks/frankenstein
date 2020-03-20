package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cortexproject/cortex/integration/e2ecortex"
)

type ServiceType string

const (
	Distributor   = ServiceType("distributor")
	Ingester      = ServiceType("ingester")
	Querier       = ServiceType("querier")
	QueryFrontend = ServiceType("query-frontend")
	TableManager  = ServiceType("table-manager")
)

var (
	// Service-specific metrics prefixes which shouldn't be used by any other service.
	serviceMetricsPrefixes = map[ServiceType][]string{
		Distributor:   []string{},
		Ingester:      []string{},
		Querier:       []string{},
		QueryFrontend: []string{"cortex_frontend", "cortex_query_frontend"},
		TableManager:  []string{},
	}
)

func assertServiceMetricsPrefixes(t *testing.T, serviceType ServiceType, service *e2ecortex.CortexService) {
	metrics, err := service.Metrics()
	require.NoError(t, err)

	// Build the list of blacklisted metrics prefixes for this specific service.
	blacklist := getBlacklistedMetricsPrefixesByService(serviceType)

	// Ensure no metric name matches the blacklisted prefixes.
	for _, metricLine := range strings.Split(metrics, "\n") {
		metricLine = strings.TrimSpace(metricLine)
		if metricLine == "" || strings.HasPrefix(metricLine, "#") {
			continue
		}

		for _, prefix := range blacklist {
			assert.NotRegexp(t, "^"+prefix, metricLine, "service: %s", service.Name())
		}
	}
}

func getBlacklistedMetricsPrefixesByService(serviceType ServiceType) []string {
	blacklist := []string{}

	// Add any service-specific metrics prefix excluding the service itself.
	for t, prefixes := range serviceMetricsPrefixes {
		if t == serviceType {
			continue
		}

		blacklist = append(blacklist, prefixes...)
	}

	return blacklist
}
