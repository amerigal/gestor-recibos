FROM golang:1.17-alpine

LABEL maintainer="antoniomerino99@gmail.com"

RUN mkdir -p /app/test \
    && adduser --disabled-password gestor_recibos

WORKDIR /app/test

USER gestor_recibos
    
RUN go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT ["task", "test"]