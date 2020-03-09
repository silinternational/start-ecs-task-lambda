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

func getRunTaskInput(config LambdaConfig) (ecs.RunTaskInput, error) {
	// optional ECS_CLUSTER
	cluster, _ := getConfigOrEnv(config.ECSCluster, "ECS_CLUSTER")

	taskDef, err := getConfigOrEnv(config.ECSTaskDefinition, "ECS_TASK_DEFINITION")
	if err != nil {
		return ecs.RunTaskInput{}, err
	}

	input := ecs.RunTaskInput{
		Cluster: &cluster,
		TaskDefinition: &taskDef,
	}
	return input, nil
}

func handler(config LambdaConfig) error {
	runTaskInput, err := getRunTaskInput(config)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("Starting ECS task.")

	c := ecs.ECS{}

	output, err := c.RunTask(&runTaskInput)
	if err != nil {
		err := errors.New("error starting ecs task: " + err.Error())
		log.Println(err.Error())
		return err
	}

	log.Printf("start ECS task output: %+v", output)
	log.Println("Completed starting ECS task.")

	return nil
}