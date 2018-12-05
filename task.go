package thermos_payload

type Resources struct {
	GPU  float64 `json:"gpu"`
	Disk uint64     `json:"disk"`
	RAM  uint64     `json:"ram"`
	CPU  float64 `json:"cpu"`
}

// OrderConstraint only used in the Task.Constraints
type OrderConstraint struct {
	Order []string `json:"order"`
}

func (o *OrderConstraint) init() {
	o.Order = make([]string, 0)
}

func (o *OrderConstraint) SetOrder(p []string) {
	o.Order = p
}

func (o *OrderConstraint) AddProcess(process *Process) {
	if process.Final {
		return
	}
	o.Order = append(o.Order, process.Name)
}

func NewOrderConstraint() *OrderConstraint {
	return &OrderConstraint{}
}

func (c *OrderConstraint) GetValue() interface{} {
	return c.Order
}

type Task struct {
	Name             string            `json:"name"`
	Processes        []*Process        `json:"processes"`
	FinalizationWait int               `json:"finalization_wait,omitempty"`
	MaxFailures      int               `json:"max_failures"`
	MaxConcurrency   int               `json:"max_concurrency"`
	Resources        *Resources        `json:"resources"`
	Constraints      []OrderConstraint `json:"constraints"`
}

type ThermosTaskConfig struct {
	Task                *Task              `json:"task"`
	Cluster             string             `json:"cluster"`
	Role                string             `json:"role"`
	Environment         string             `json:"environment"`
	Name                string             `json:"name"`
	HealthCheckConfig   *HealthCheckConfig `json:"health_check_config,omitempty"`
	Service             bool               `json:"service"`
	MaxTaskFailures     int                `json:"max_task_failures"`
	CronCollisionPolicy string             `json:"cron_collision_policy"`
	EnableHooks         bool               `json:"enable_hooks"`
	Lifecycle           *Lifecycle         `json:"lifecycle,omitempty"`
	Priority            int                `json:"priority"`
	Production          bool               `json:"production"`
}

func (t *Task) AddProcess(p ... *Process) {
	if len(t.Constraints) == 0 {
		t.AddConstraint(OrderConstraint{Order: []string{}})
	}
	for _, proc := range p {
		t.Processes = append(t.Processes, proc)
		t.Constraints[0].AddProcess(proc)
	}
}

func (t *Task) AddConstraint(c OrderConstraint) {
	t.Constraints = []OrderConstraint{c}
}
