FROM golang:1.23
WORKDIR /app
COPY go.* ./
RUN go mod tidy 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux  go build -o user-service main.go
EXPOSE 4000
CMD ["./user-service"]
