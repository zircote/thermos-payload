package thermos_payload

import (
	"encoding/json"
	"fmt"
)

func DefaultThermosTaskConfig(name string) *ThermosTaskConfig {
	return &ThermosTaskConfig{
		Name:                name,
		Role:                "www-data",
		Environment:         "prod",
		Cluster:             "devcluster",
		EnableHooks:         false,
		Priority:            0,
		CronCollisionPolicy: "KILL_EXISTING",
		MaxTaskFailures: 1,
		Task: &Task{
			Name:             name,
			Resources:        &Resources{GPU: 0, CPU: 0, RAM: 0, Disk: 0},
			MaxFailures:      1,
			MaxConcurrency:   0,
			FinalizationWait: 30,
			Constraints:      []OrderConstraint{},
			Processes:        []*Process{},
		},
		Lifecycle: &Lifecycle{
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
					Endpoint:             "/health",
					ExpectedResponse:     "ok",
					ExpectedResponseCode: 0,
				},
			},
		},
		Service: true,
	}
}

func (t *ThermosTaskConfig) IsService() {
	t.Service = true
}

func (t *ThermosTaskConfig) NewHttpHealthChecker(config *HttpHealthCheckerConfig) HealthChecker {
	return HealthChecker{HTTP: config}
}

func (t *ThermosTaskConfig) NewShellHealthChecker(config *ShellHealthCheckerConfig) HealthChecker {
	return HealthChecker{Shell: config}
}
func (t *ThermosTaskConfig) String() string {
	s, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprint(err)
	}
	return string(s)
}

func (t *ThermosTaskConfig) FromString(data string) error {

	err := json.Unmarshal([]byte(data), t)
	if err != nil {
		return err
	}
	return nil
}
