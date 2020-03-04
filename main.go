package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/ecs"
)

type LambdaConfig struct {
	ECSCluster string `json:"ecs_cluster"`
	ECSContainerInstance string `json:"ecs_container_instance"`
	ECSTaskDefinition string `json:"ecs_task_definition"`
}

func main() {
	lambda.Start(handler)
}

// Ignoring the following input values ...
//  EnableECSManagedTags
//  Group
//  NetworkConfiguration
//  Overrides
//  PropagateTags
//  ReferenceId
//  StartedBy
//  Tags
func getStartTaskInput(config LambdaConfig) (ecs.StartTaskInput, error) {
	cluster := config.ECSCluster
	if cluster == "" {
		err := errors.New("missing required config: ecs_cluster")
		return ecs.StartTaskInput{}, err
	}

	container := config.ECSContainerInstance
	if container == "" {
		err := errors.New("missing required config: ecs_container_instance")
		return ecs.StartTaskInput{}, err
	}

	taskDef := config.ECSTaskDefinition
	if taskDef == "" {
		err := errors.New("missing required config: ecs_task_definition")
		return ecs.StartTaskInput{}, err
	}

	input := ecs.StartTaskInput{
		Cluster: &cluster,
		ContainerInstances: []*string{&container},
		TaskDefinition: &taskDef,
	}
	return input, nil
}

func handler(config LambdaConfig) error {
	startTaskInput, err := getStartTaskInput(config)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("Starting ECS task.")

	c := ecs.ECS{}

	output, err := c.StartTask(&startTaskInput)
	if err != nil {
		err := errors.New("error starting ecs task: " + err.Error())
		log.Println(err.Error())
		return err
	}

	log.Printf("start ECS task output: %+v", output)
	log.Println("Completed starting ECS task.")

	return nil
}