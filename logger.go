package thermos_payload

type RotatePolicy struct {
	// Maximum size (in bytes) of an individual log file. (Default: 100 MiB)
	LogSize int64 `json:"log_size"`
	// The maximum number of backups to retain. (Default: 5)
	Backups int64 `json:"backups"`
}

type Logger struct {
	/*
	 * Logger objects specify a destination for Process logs which is, by default, file - a pair of stdout
	 * and stderr files. Its also possible to specify console to get logs output to the Process stdout and
	 * stderr streams, none to suppress any logs output or both to send logs to files and console streams.
	 */
	Destination string `json:"destination"`

	// The default Logger mode is standard which lets the stdout and stderr streams grow without bound.
	Mode string `json:"mode"`

	/*
	 * A RotatePolicy describes log rotation behavior for when mode is set to rotate and it is ignored otherwise.
	 * If rotate is Empty or RotatePolicy() when the mode is set to rotate the defaults below are used.
	 */
	Rotate *RotatePolicy `json:"rotate"`
}

func NewLogger(dest, mode string, log_size, backups int64) *Logger {
	return &Logger{
		Destination: dest,
		Mode:        mode,
		Rotate:      &RotatePolicy{LogSize: log_size, Backups: backups}}
}
