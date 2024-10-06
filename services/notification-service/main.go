package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type NotificationMessage struct {
	UserID     uint   `json:"user_id"`
	EventID    uint   `json:"event_id"`
	SeatsTaken int    `json:"seats_taken"`
	Timestamp  int64  `json:"timestamp"`
	Email      string `json:"email"`
}

func sendEmail(to string, body string) {
	from := "youremail@example.com"
	password := "your-email-password"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: Booking Confirmation\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return
	}

	fmt.Println("Email sent successfully to:", to)
}

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "notification-service",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Printf("Failed to create Kafka consumer: %s\n", err)
		return
	}

	defer consumer.Close()

	consumer.SubscribeTopics([]string{"booking-notifications"}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			//unmarshal message
			var notificationMessage NotificationMessage
			err := json.Unmarshal(msg.Value, &notificationMessage)
			if err != nil {
				fmt.Printf("Failed to unmarshal message: %v\n", err)
				continue
			}
			// make a template for email
			emailBody := notificationMessage.generateBody()
			sendEmail(notificationMessage.Email, emailBody)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func (notification *NotificationMessage) generateBody() string {
	eventTime := time.Unix(notification.Timestamp, 0).Format("02 Jan 2006 15:04:05")

	message := fmt.Sprintf(
		"Hello,\n\nYour booking was successful!\n\nDetails:\n"+
			"User ID: %d\nEvent ID: %d\nSeats Taken: %d\nEvent Date & Time: %s\n\nThank you for booking!\n",
		notification.UserID,
		notification.EventID,
		notification.SeatsTaken,
		eventTime,
	)

	return message
}
