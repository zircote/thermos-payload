package thermos_payload

type HttpHealthCheckerConfig struct {
	ExpectedResponseCode int    `json:"expected_response_code"`
	Endpoint             string `json:"endpoint"`
	ExpectedResponse     string `json:"expected_response"`
}

type ShellHealthCheckerConfig struct {
	ShellCommand string `json:"shell_command,omitempty"`
}

type HealthChecker struct {
	HTTP  *HttpHealthCheckerConfig  `json:"http,omitempty"`
	Shell *ShellHealthCheckerConfig `json:"shell,omitempty"`
}

type HealthCheckConfig struct {
	InitialIntervalSecs    float64        `json:"initial_interval_secs"`
	HealthChecker          *HealthChecker `json:"health_checker,omitempty"`
	IntervalSecs           float64        `json:"interval_secs"`
	TimeoutSecs            float64        `json:"timeout_secs"`
	MaxConsecutiveFailures int            `json:"max_consecutive_failures"`
}

func DefaultHealthChecker() *HealthCheckConfig {
	return &HealthCheckConfig{
		IntervalSecs:           15,
		MaxConsecutiveFailures: 1,
		TimeoutSecs:            1,
		InitialIntervalSecs:    5,
		HealthChecker: &HealthChecker{
			HTTP: &HttpHealthCheckerConfig{
				Endpoint:         "/health",
				ExpectedResponse: "ok",
				ExpectedResponseCode: 0,
			},
		},
	}
}
