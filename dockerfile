FROM golang:1.23
RUN apt-get update && apt-get install -y librdkafka-dev
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o event-service ./services/event-service/main.go
EXPOSE 4000
CMD ["./event-service"]
