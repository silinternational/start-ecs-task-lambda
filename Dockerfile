FROM golang:latest

RUN apt-get update && apt-get install -y groff python python-pip && pip install awscli

RUN mkdir -p $GOPATH/src/start-ecs-task-lambda
WORKDIR $GOPATH/src/start-ecs-task-lambda

COPY ./main.go $GOPATH/src/start-ecs-task-lambda/
COPY ./codeship.sh $GOPATH/src/start-ecs-task-lambda/

RUN go get ./...