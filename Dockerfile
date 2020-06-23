FROM golang:1.14-alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download -x

COPY . .

RUN go build -o echoSwaggerTest

FROM alpine:3.11

WORKDIR /app

COPY --from=builder build/echoSwaggerTest .

EXPOSE 7777

CMD ["/app/echoSwaggerTest"]