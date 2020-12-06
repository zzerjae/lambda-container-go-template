# lambda-container-go-template

This is a minimal starting point for create a Lambda container go image.

## Usage

1. Create a new repository from lambda-container-go-template By Clicking "Use this template".

2. Clone the repository to your workspace.

3. Write some code for your program.  

4. Build the container image and run it to make sure the program works. 

```sh

# Build the Docker image
$ docker build -t hello-lambda-container-go .

# Run the Container
$ docker run -it -p 9000:8080  hello-lambda-container-go:latest

# Test a function invocation with cURL. Here, I am passing an JSON payload.
$ curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name":"world"}'

```

If you haven't modified the code, you'll see a response that says `{"message":"Hello, world"}.`

## so what now?

Bring in your own code and start writing programs.

When you're done writing the code, build your container image and upload it to the [AWS ECR Repository](https://aws.amazon.com/ko/ecr/).

Then Create a Lambda function with container image option in AWS Lambda Console.

If you prefer, you can start with the [SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html) created by AWS, or the [serverless](https://www.serverless.com/) framework popular in the community.

## references

- [ECR Public Gallery - AWS Lambda base images for Go](https://gallery.ecr.aws/lambda/go)

- [AWS Blog - new-for-aws-lambda-container-image-support](https://aws.amazon.com/blogs/aws/new-for-aws-lambda-container-image-support/)

- [AWS Docs - Building Lambda functions with Go](https://docs.aws.amazon.com/lambda/latest/dg/go-image.html)

- [GitHub - aws/aws-lambda-go](https://github.com/aws/aws-lambda-go)

- [GitHub - aws/aws-sam-cli-app-templates](https://github.com/aws/aws-sam-cli-app-templates/tree/master/go1.x-image/cookiecutter-aws-sam-hello-golang-lambda-image/%7B%7Bcookiecutter.project_name%7D%7D)

- [GitHub - serverless/serverless](https://github.com/serverless/serverless)