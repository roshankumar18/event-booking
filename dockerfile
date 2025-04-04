FROM golang:1.23
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY ../../pkg .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux  go build -o user-service ./services/user-service/main.go
EXPOSE 4000
CMD ["./user-service"]
# 