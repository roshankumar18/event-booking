# Event Booking System

## Overview

The **Event Booking System** is a microservices-based application that allows users to create events, book tickets, and receive notifications for bookings. It leverages Kafka for inter-service communication, JWT for authentication, and a modular architecture for scalability and maintainability.

## Features

- **User Roles**: Two types of users (Event Creators and Bookers), with role-based access control.
- **Event Management**: Event Creators can create and manage events, while Bookers can view and book tickets.
- **Booking System**: Bookers can reserve tickets, with real-time seat availability checks.
- **Notification System**: Real-time notifications are sent when bookings are confirmed, powered by Kafka.
- **Microservices**: Designed as a collection of services for better scalability and fault tolerance.

## Technologies

- **Golang**: Backend services written in Go.
- **Gin**: Web framework for building RESTful APIs.
- **Kafka**: Message broker used for notification handling and asynchronous communication.
- **PostgreSQL**: Database for storing user, event, and booking data.
- **JWT**: Secure token-based authentication.
- **Docker**: Containerization for easy deployment.

## Services

1. **User Service**: Handles user registration, login, and role-based access control.
2. **Event Service**: Manages events, including creation and seat availability.
3. **Booking Service**: Processes bookings, checks seat availability, and updates seat count.
4. **Notification Service**: Listens to booking events via Kafka and sends notifications to users.

## Setup Instructions

### Prerequisites

- Go 1.19+
- Docker
- Kafka
- PostgreSQL

### Local Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/roshankumar18/event-booking.git
   cd event-booking
   ```
