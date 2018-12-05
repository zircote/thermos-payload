package thermos_payload

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTask(t *testing.T) {
	type args struct {
		name  string
		procs []*Process
	}
	tests := []struct {
		name string
		args args
		want *Task
	}{
		{
			"test_task",
			args{
				"test_task",
				[]*Process{
					{"test", false, false, 0, 5, "echo 1", false, nil},
					{"test2", false, false, 0, 5, "echo 2", false, nil},
				},
			},
			&Task{
				Name: "test_task",
				Processes: []*Process{
					{"test", false, false, 0, 5, "echo 1", false, nil},
					{"test2", false, false, 0, 5, "echo 2", false, nil},
				},
				FinalizationWait: 30,
				MaxFailures:      1,
				MaxConcurrency:   0,
				Resources:        &Resources{},
				Constraints:      []OrderConstraint{{Order: []string{"test", "test2"}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttc := DefaultThermosTaskConfig(tt.args.name)
			ttc.Task.AddProcess(tt.args.procs...)
			if !reflect.DeepEqual(ttc.Task, tt.want) {
				t.Errorf("NewTask() = %v, want %v", ttc.Task, tt.want)
			}
		})
	}
}

func TestNewThermosService(t *testing.T) {
	type args struct {
		name    string
	}
	tests := []struct {
		name string
		args args
		want *ThermosTaskConfig
	}{
		{
			"service_test",
			args{
				"st",
			},
			&ThermosTaskConfig{Service: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got :=  DefaultThermosTaskConfig(tt.args.name)
			got.IsService()
			if got.Service != tt.want.Service {
				t.Errorf("NewThermosService() = %v, want %v", got.Service, tt.want.Service)
			}
		})
	}
}

func TestNewSequentialTask(t *testing.T) {
	type args struct {
		name  string
		procs []*Process
	}
	tests := []struct {
		name string
		args args
		want *Task
	}{
		{
			"test_task",
			args{
				"test_task",
				[]*Process{
					{"test1", false, false, 0, 5, "echo 1", false, nil},
					{"test2", false, false, 0, 5, "echo 2", false, nil},
				},
			},
			&Task{
				Name: "test_task",
				Constraints: []OrderConstraint{{Order: []string{"test1", "test2"}},
				},
				Processes: []*Process{
					{"test1", false, false, 0, 5, "echo 1", false, nil},
					{"test2", false, false, 0, 5, "echo 2", false, nil},
				},
				Resources:        &Resources{},
				MaxFailures:      1,
				MaxConcurrency:   0,
				FinalizationWait: 30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultThermosTaskConfig(tt.args.name)
			got.IsService()
			got.Task.AddProcess(tt.args.procs...)
			if !reflect.DeepEqual(got.Task, tt.want) {
				t.Errorf("NewSequentialTask() = %v, want %v", got, tt.want)
			}
		})
	}
}



func ExampleDefaultThermosTaskConfig_output() {


	thermosTaskConfig := DefaultThermosTaskConfig(
		"hello_world",
	)
	// 0.5 CPU 128M RAM 128M Disk 0 GPU
	thermosTaskConfig.Task.Resources = &Resources{CPU: .5, Disk:128, RAM:128}
	thermosTaskConfig.Task.AddProcess(DefaultProcess("hello", `
	while true; do
	   echo hello world
	   # Privde a port {{thermos.ports[health]}}
	   sleep 10
	done`))
	j := thermosTaskConfig.String()
	fmt.Println(j)
	// Output:  {"task":{"name":"hello_world","processes":[{"name":"hello","daemon":false,"ephemeral":false,"max_failures":1,"min_duration":5,"cmdline":"\n\twhile true; do\n\t   echo hello world\n\t   # Privde a port {{thermos.ports[health]}}\n\t   sleep 10\n\tdone","final":false}],"finalization_wait":30,"max_failures":1,"max_concurrency":0,"resources":{"gpu":0,"disk":128,"ram":128,"cpu":0.5},"constraints":[{"order":["hello"]}]},"cluster":"devcluster","role":"www-data","environment":"prod","name":"hello_world","health_check_config":{"initial_interval_secs":15,"health_checker":{"http":{"expected_response_code":0,"endpoint":"/health","expected_response":"ok"}},"interval_secs":10,"timeout_secs":1,"max_consecutive_failures":1},"service":true,"max_task_failures":1,"cron_collision_policy":"KILL_EXISTING","enable_hooks":false,"lifecycle":{"http":{"graceful_shutdown_endpoint":"/quitquitquit","port":"health","shutdown_endpoint":"/abortabortabort"}},"priority":0,"production":false}
}



func TestOrderConstraint_SetOrder(t *testing.T) {
	type fields struct {
		Order []string
	}
	type args struct {
		p []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			"test",
			fields{Order: []string{"test", "value"}},
			args{[]string{"one", "two", "three"}},
			[]string{"one", "two", "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderConstraint{
				Order: tt.fields.Order,
			}
			o.SetOrder(tt.args.p)
			if got := o.GetValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderConstraint.SetOrder(); OrderConstraint.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderConstraint_AddProcess(t *testing.T) {

	type args struct {
		process []*Process
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test",
			args{[]*Process{
				&Process{"test", false, false, 0, 5, "echo 1", false, nil},
				&Process{"value", false, false, 0, 5, "echo 2", false, nil},
			},
			},
			[]string{"test", "value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderConstraint{}
			for _, proc := range tt.args.process {
				o.AddProcess(proc)
			}
			if got := o.GetValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderConstraint().AddProcess = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOrderConstraint(t *testing.T) {
	tests := []struct {
		name string
		want *OrderConstraint
	}{
		{
			"default",
			&OrderConstraint{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderConstraint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderConstraint() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestOrderConstraint_GetValue(t *testing.T) {
	type fields struct {
		Order []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{{
		"default key",
		fields{Order: []string{"test1", "test2"}},
		[]string{"test1", "test2"},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &OrderConstraint{
				Order: tt.fields.Order,
			}
			if got := c.GetValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderConstraint.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}