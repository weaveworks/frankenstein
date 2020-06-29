package queryrange

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStepAlign(t *testing.T) {
	for i, tc := range []struct {
		input, expected *PrometheusRequest
	}{
		{
			input: &PrometheusRequest{
				Start: 0,
				End:   100,
				Step:  10,
			},
			expected: &PrometheusRequest{
				Start: 0,
				End:   100,
				Step:  10,
			},
		},

		{
			input: &PrometheusRequest{
				Start: 2,
				End:   102,
				Step:  10,
			},
			expected: &PrometheusRequest{
				Start: 0,
				End:   100,
				Step:  10,
			},
		},
		{
			input: &PrometheusRequest{
				Start: 0,
				End:   1000,
				Step:  100,
			},
			expected: &PrometheusRequest{
				Start: 0,
				End:   1000,
				Step:  100,
			},
		},
		{
			input: &PrometheusRequest{
				Start: 2,
				End:   1002,
				Step:  100,
			},
			expected: &PrometheusRequest{
				Start: 2,
				End:   1002,
				Step:  100,
			},
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var result *PrometheusRequest
			s := stepAlign{
				next: HandlerFunc(func(_ context.Context, req Request) (Response, error) {
					result = req.(*PrometheusRequest)
					return nil, nil
				}),
				maxStepAlignmentMs: 99,
			}
			_, err := s.Do(context.Background(), tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}
