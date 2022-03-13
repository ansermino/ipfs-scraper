FROM golang:1.17-alpine AS builder

RUN apk add build-base
WORKDIR /code
COPY go.mod .
COPY go.sum .
RUN go mod download

# build the binary
ADD . .
RUN env GOOS=linux GOARCH=amd64 go build -o /api ./main.go

# final stage
FROM alpine:3.14.0

RUN apk add build-base

RUN apk add curl
COPY --from=builder /api /

RUN chmod +x /api
ENTRYPOINT ["/api"]