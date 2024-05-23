# Dockerfile for Go app
FROM golang:1.22.1 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /backend
EXPOSE 8080
CMD ["/backend"]
