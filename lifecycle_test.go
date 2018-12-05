package thermos_payload

import (
	"reflect"
	"testing"
)

func TestDefaultLifecycleConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Lifecycle
	}{
		{
			"default",
			&Lifecycle{
				HTTP: &HttpLifecycleConfig{
					GracefulShutdownEndpoint: "/quitquitquit",
					ShutdownEndpoint: "/abortabortabort",
					Port: "health",
					GracefulShutdownWaitSecs: 5,
					ShutdownWaitSecs: 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultLifecycleConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultLifecycleConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
