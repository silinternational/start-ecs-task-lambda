package main

import (
	"errors"
	"log"
	"os"

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

func getConfigOrEnv(configVal, envKey string) (string, error) {
	if configVal != "" {
		return configVal, nil
	}

	value := os.Getenv(envKey)
	if value != "" {
		return configVal, nil
	}

	return "", errors.New("Missing required config/env var " + envKey)
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
	cluster, err := getConfigOrEnv(config.ECSCluster, "ECS_CLUSTER")
	if err != nil {
		return ecs.StartTaskInput{}, err
	}

	container, err := getConfigOrEnv(config.ECSContainerInstance, "ECS_CONTAINER_INSTANCE")
	if err != nil {
		return ecs.StartTaskInput{}, err
	}

	taskDef, err := getConfigOrEnv(config.ECSTaskDefinition, "ECS_TASK_DEFINITION")
	if err != nil {
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