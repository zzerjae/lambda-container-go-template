FROM golang:1.14-alpine AS builder

WORKDIR /go/src

COPY go.mod main.go ./

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../bin

FROM public.ecr.aws/lambda/go:1

COPY --from=builder /go/bin/ /var/task/

CMD [ "main" ]