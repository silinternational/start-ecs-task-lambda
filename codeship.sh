#!/usr/bin/env bash

# Exit script with error if any step fails.
set -e

# Echo out all commands for monitoring progress
set -x

# Build all the things
go build .

ls -al

zip start-ecs-task-lambda.zip start-ecs-task-lambda

aws s3 cp ./start-ecs-task-lambda.zip s3://gtis-ops-lambdas/$1/start-ecs-task-lambda.zip

