module github.com/roshankumar18/event-booking/services/notification-service

go 1.23.1

replace (
	github.com/roshankumar18/event-booking/pkg => ../../pkg
	github.com/roshankumar18/event-booking/utils => ../../utils
)

require (
	github.com/confluentinc/confluent-kafka-go v1.9.2
	github.com/roshankumar18/event-booking/utils v0.0.0
)

require (
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.25.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
