package thermos_payload

type HttpLifecycleConfig struct {
	GracefulShutdownEndpoint string `json:"graceful_shutdown_endpoint,omitempty"`
	Port                     string `json:"port,omitempty"`
	ShutdownEndpoint         string `json:"shutdown_endpoint,omitempty"`
	GracefulShutdownWaitSecs int    `json:"graceful_shutdown_wait_secs,omitempty"`
	ShutdownWaitSecs         int    `json:"shutdown_wait_secs,omitempty"`
}

type Lifecycle struct {
	HTTP *HttpLifecycleConfig `json:"http"`
}

func DefaultLifecycleConfig() *Lifecycle {
	return &Lifecycle{
		HTTP: &HttpLifecycleConfig{
			GracefulShutdownEndpoint: "/quitquitquit",
			Port:                     "health",
			ShutdownEndpoint:         "/abortabortabort",
			GracefulShutdownWaitSecs: 5,
			ShutdownWaitSecs:         5,
		},
	}
}

func (t *ThermosTaskConfig) NewHttLifecycleConfig(config *HttpLifecycleConfig) Lifecycle {
	return Lifecycle{HTTP: config}
}
