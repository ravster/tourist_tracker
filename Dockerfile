FROM golang:1.26.0-alpine3.23

RUN mkdir -p /app
WORKDIR /app
RUN apk add build-base

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

COPY . .

CMD ["tail", "-F", "/tmp/keep_docker_alive"]
