# Build stage
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN apk add build-base
RUN go mod vendor
RUN go build -v -o go-dating-app ./app/app.go

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/go-dating-app .

EXPOSE 8080
ENTRYPOINT [ "/app/go-dating-app" ]