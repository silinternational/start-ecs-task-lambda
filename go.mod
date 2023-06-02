module github.com/silinternational/start-ecs-task-lambda

// Ensure this version tracks with Dockerfile
go 1.20

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/aws/aws-sdk-go v1.44.274
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect
