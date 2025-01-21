FROM golang:1.21-alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o ./dist/redis-clone main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/dist/redis-clone /app/dist/redis-clone

CMD ["/app/dist/redis-clone"]