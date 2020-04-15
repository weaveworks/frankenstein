package push

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/snappy"
	"github.com/prometheus/prometheus/prompb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cortexproject/cortex/pkg/distributor"
	"github.com/cortexproject/cortex/pkg/ingester/client"
)

func TestHandler_remoteWrite(t *testing.T) {
	req := createRemoteWriteRequest(t)
	resp := httptest.NewRecorder()

	handler := Handler(distributor.Config{MaxRecvMsgSize: 100000}, verifyWriteRequestHandler(t, client.API))
	handler.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestHandler_cortexWriteRequest(t *testing.T) {
	req := createCortexRemoteWriteRequest(t)
	resp := httptest.NewRecorder()

	handler := Handler(distributor.Config{MaxRecvMsgSize: 100000}, verifyWriteRequestHandler(t, client.RULE))
	handler.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func verifyWriteRequestHandler(t *testing.T, expectSource client.WriteRequest_SourceEnum) func(ctx context.Context, request *client.WriteRequest) (response *client.WriteResponse, err error) {
	t.Helper()
	return func(ctx context.Context, request *client.WriteRequest) (response *client.WriteResponse, err error) {
		assert.Equal(t, expectSource, request.Source)
		return &client.WriteResponse{}, nil
	}
}

func createRemoteWriteRequest(t *testing.T) *http.Request {
	t.Helper()
	input := prompb.WriteRequest{
		Timeseries: []prompb.TimeSeries{
			{
				Labels: []prompb.Label{
					{Name: "__name__", Value: "foo"},
				},
				Samples: []prompb.Sample{
					{Value: 1, Timestamp: time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC).UnixNano()},
				},
			},
		},
	}
	inoutBytes, err := input.Marshal()
	require.NoError(t, err)
	inoutBytes = snappy.Encode(nil, inoutBytes)
	req, err := http.NewRequest("POST", "http://localhost/", bytes.NewReader(inoutBytes))
	require.NoError(t, err)
	req.Header.Add("Content-Encoding", "snappy")
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("X-Prometheus-Remote-Write-Version", "0.1.0")
	return req
}

func createCortexRemoteWriteRequest(t *testing.T) *http.Request {
	t.Helper()
	ts := &client.TimeSeries{
		Labels: []client.LabelAdapter{
			{Name: "__name__", Value: "foo"},
		},
		Samples: []client.Sample{
			{Value: 1, TimestampMs: time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC).UnixNano()},
		},
	}
	input := client.WriteRequest{
		Timeseries: []client.PreallocTimeseries{
			client.PreallocTimeseries{ts},
		},
		Source: client.RULE,
	}
	inoutBytes, err := input.Marshal()
	require.NoError(t, err)
	inoutBytes = snappy.Encode(nil, inoutBytes)
	req, err := http.NewRequest("POST", "http://localhost/", bytes.NewReader(inoutBytes))
	require.NoError(t, err)
	req.Header.Add("Content-Encoding", "snappy")
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("X-Prometheus-Remote-Write-Version", "0.1.0")
	return req
}
