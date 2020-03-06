#!/usr/bin/env bash

# Exit script with error if any step fails.
set -e

# Echo out all commands for monitoring progress
set -x

# Build all the things
go build .

ls -al

aws s3 cp ./start-ecs-task-lambda s3://gtis-ops-lambdas/$1/start-ecs-task-lambda