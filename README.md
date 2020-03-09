# start-ecs-task-lambda
Golang lambda that starts an ECS task

## Config
```
 {
	"ecs_cluster": "abc123",
	"ecs_container_instance": "abc123",
	"ecs_task_definition": "abc123"
}
```

If any of these are missing or empty, the lambda will use the associated 
environment variable:
ECS_CLUSTER, ECS_CONTAINER_INSTANCE, ECS_TASK_DEFINITION.