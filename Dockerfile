FROM golang:1.17-alpine

LABEL maintainer="antoniomerino99@gmail.com"

RUN mkdir -p /app/test \
    && adduser --disabled-password gestor_recibos \
    && chown -R gestor_recibos:gestor_recibos /app/test

WORKDIR /app/test

USER gestor_recibos

COPY go.mod /app/test
    
RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT ["task", "test"]