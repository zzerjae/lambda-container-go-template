package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func Handler(ctx context.Context, req Request) (Response, error) {
	return Response{Message: "Hello, " + req.Name}, nil
}

func main() {
	lambda.Start(Handler)
}
