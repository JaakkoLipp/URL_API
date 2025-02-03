# Use an official Go runtime as a build stage
FROM golang:1.23.2 AS builder

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o url-shortener main.go

# Use a minimal base image for the final container
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/url-shortener .

EXPOSE 8080
CMD ["./url-shortener"]
