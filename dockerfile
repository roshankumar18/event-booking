FROM golang:1.23
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
EXPOSE 4000
CMD ["go" "run" "./services/event-service/main.go"]