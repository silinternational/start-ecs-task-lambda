# start-ecs-task-lambda
Golang lambda that runs a one-off ECS task

## Config
```
 {
	"ecs_cluster": "abc123",  
	"ecs_task_definition": "abc123"
}
```

If either of these are missing or empty, the lambda will use the associated 
environment variable:
ECS_CLUSTER,  ECS_TASK_DEFINITION.

(ecs_cluster is optional, but ecs_task_definition is required.)