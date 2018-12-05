package thermos_payload

import (
	"reflect"
	"testing"
)

func TestDefaultHealthChecker(t *testing.T) {
	tests := []struct {
		name string
		want *HealthCheckConfig
	}{
		{
			"default",
			&HealthCheckConfig{
				InitialIntervalSecs: 5,
				IntervalSecs: 15,
				TimeoutSecs: 1,
				MaxConsecutiveFailures: 1,
				HealthChecker: &HealthChecker{
					HTTP: &HttpHealthCheckerConfig{
						ExpectedResponse: "ok",
						Endpoint: "/health",
						ExpectedResponseCode: 0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultHealthChecker(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultHealthChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
