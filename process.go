package thermos_payload

type Process struct {
	Name        string  `json:"name"`
	Daemon      bool    `json:"daemon"`
	Ephemeral   bool    `json:"ephemeral"`
	MaxFailures int     `json:"max_failures"`
	MinDuration int     `json:"min_duration"`
	Cmdline     string  `json:"cmdline"`
	Final       bool    `json:"final"`
	Logger      *Logger `json:"logger,omitempty"`
}

func NewDaemonProcess(name, cmd string) *Process {
	p := DefaultProcess(name, cmd)
	p.Daemon = true
	return p
}

func NewEphemeralProcess(name, cmd string) *Process {
	p := DefaultProcess(name, cmd)
	p.Ephemeral = true
	return p
}

func NewFinalProcess(name, cmd string) *Process {
	p := DefaultProcess(name, cmd)
	p.Final = true
	return p
}

func DefaultProcess(name, cmd string) *Process {
	return &Process{
		Name:        name,
		Cmdline:     cmd,
		MaxFailures: 1,
		MinDuration: 5,
		Daemon:      false,
		Ephemeral:   false,
		Final:       false,
		Logger:      nil,
	}
}
