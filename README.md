# Thermos Payload

Helper library to provide validation and tools to create Apache Aurora Thermos centric Payloads

```golang

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
	
```

Result:

```JSON
 {
   "task": {
     "name": "hello_world",
     "processes": [
       {
         "name": "hello",
         "daemon": false,
         "ephemeral": false,
         "max_failures": 1,
         "min_duration": 5,
         "cmdline": "\n\twhile true; do\n\t   echo hello world\n\t   # Privde a port {{thermos.ports[health]}}\n\t   sleep 10\n\tdone",
         "final": false
       }
     ],
     "finalization_wait": 30,
     "max_failures": 1,
     "max_concurrency": 0,
     "resources": {
       "gpu": 0,
       "disk": 128,
       "ram": 128,
       "cpu": 0.5
     },
     "constraints": [
       {
         "order": [
           "hello"
         ]
       }
     ]
   },
   "cluster": "devcluster",
   "role": "www-data",
   "environment": "prod",
   "name": "hello_world",
   "health_check_config": {
     "initial_interval_secs": 15,
     "health_checker": {
       "http": {
         "expected_response_code": 0,
         "endpoint": "/health",
         "expected_response": "ok"
       }
     },
     "interval_secs": 10,
     "timeout_secs": 1,
     "max_consecutive_failures": 1
   },
   "service": true,
   "max_task_failures": 1,
   "cron_collision_policy": "KILL_EXISTING",
   "enable_hooks": false,
   "lifecycle": {
     "http": {
       "graceful_shutdown_endpoint": "/quitquitquit",
       "port": "health",
       "shutdown_endpoint": "/abortabortabort"
     }
   },
   "priority": 0,
   "production": false
 }


 ```