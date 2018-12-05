package thermos_payload

import (
	"reflect"
	"testing"
)

func TestNewDaemonProcess(t *testing.T) {
	type args struct {
		name string
		cmd  string
	}
	tests := []struct {
		name string
		args args
		want *Process
	}{
		{
			"daemon",
			args{"test_daemon", "exit 0",},
			&Process{
				Name:        "test_daemon",
				Cmdline:     "exit 0",
				MinDuration: 5,
				MaxFailures: 1,
				Daemon:      true,
				Ephemeral:   false,
				Final:       false,
				Logger:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDaemonProcess(tt.args.name, tt.args.cmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDaemonProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEphemeralProcess(t *testing.T) {
	type args struct {
		name string
		cmd  string
	}
	tests := []struct {
		name string
		args args
		want *Process
	}{
		{
			"daemon",
			args{"test_daemon", "exit 0",},
			&Process{
				Name:        "test_daemon",
				Cmdline:     "exit 0",
				MinDuration: 5,
				MaxFailures: 1,
				Daemon:      false,
				Ephemeral:   true,
				Final:       false,
				Logger:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEphemeralProcess(tt.args.name, tt.args.cmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEphemeralProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFinalProcess(t *testing.T) {
	type args struct {
		name string
		cmd  string
	}
	tests := []struct {
		name string
		args args
		want *Process
	}{
		{
			"daemon",
			args{"test_daemon", "exit 0",},
			&Process{
				Name:        "test_daemon",
				Cmdline:     "exit 0",
				MinDuration: 5,
				MaxFailures: 1,
				Daemon:      false,
				Ephemeral:   false,
				Final:       true,
				Logger:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFinalProcess(tt.args.name, tt.args.cmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFinalProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProcess(t *testing.T) {
	type args struct {
		name string
		cmd  string
	}
	tests := []struct {
		name string
		args args
		want *Process
	}{
		{
			"daemon",
			args{"test_daemon", "exit 0",},
			&Process{
				Name:        "test_daemon",
				Cmdline:     "exit 0",
				MinDuration: 5,
				MaxFailures: 1,
				Daemon:      false,
				Ephemeral:   false,
				Final:       false,
				Logger:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultProcess(tt.args.name, tt.args.cmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

