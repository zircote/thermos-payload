package thermos_payload

import (
	"reflect"
	"testing"
)

func TestNewThermosJob(t *testing.T) {
	type args struct {
		name    string
		role    string
		env     string
		cluster string
		task    *Task
	}
	tests := []struct {
		name string
		args args
		want *ThermosTaskConfig
	}{
		{
			"default",
			args{
				"hello",
				"www-data",
				"prod",
				"devcluster",
				&Task{},
			},
			&ThermosTaskConfig{
				Name:        "hello",
				Role:        "www-data",
				Environment: "prod",
				Cluster:     "devcluster",
				CronCollisionPolicy: "KILL_EXISTING",
				Task: &Task{
					Name:             "hello",
					Resources:        &Resources{0,0,0,0},
					MaxFailures:      1,
					MaxConcurrency:   0,
					FinalizationWait: 30,
					Constraints:      []OrderConstraint{},
					Processes: []*Process{},
				},
				EnableHooks:       false,
				MaxTaskFailures:   1,
				Lifecycle:         &Lifecycle{
					HTTP: &HttpLifecycleConfig{
						GracefulShutdownEndpoint: "/quitquitquit",
						Port:                     "health",
						ShutdownEndpoint:         "/abortabortabort",
					},
				},
				HealthCheckConfig: &HealthCheckConfig{
					IntervalSecs:           10,
					MaxConsecutiveFailures: 1,
					TimeoutSecs:            1,
					InitialIntervalSecs:    15,
					HealthChecker: &HealthChecker{
						HTTP: &HttpHealthCheckerConfig{
							Endpoint:         "/health",
							ExpectedResponse: "ok",
							ExpectedResponseCode: 0,
						},
					},
				},
				Service:           true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultThermosTaskConfig(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewThermosJob() = %v\n\n\n, want %v", got, tt.want)
			}
		})
	}
}
