package thermos_payload

import (
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	type args struct {
		dest     string
		mode     string
		log_size int64
		backups  int64
	}
	tests := []struct {
		name    string
		args    args
		want    *Logger
		wantErr bool
	}{
		{
			"default",
			args{"test", "file", 5, 5},
			&Logger{"test", "file", &RotatePolicy{5, 5}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLogger(tt.args.dest, tt.args.mode, tt.args.log_size, tt.args.backups)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

