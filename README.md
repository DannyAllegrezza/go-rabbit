# go-rabbit 🐇

This project is a personal playground for testing out some basic abstractions over the awesome [rabbitmq/amqp091-go](https://github.com/rabbitmq/amqp091-go) package.

## Prerequisites 
- Go 1.9
- RabbitMQ running on your machine on default port `5672`
  - TIP: Easiest way to run RabbitMQ locally is through a container

## Running the app

### 1. Start the RabbitMQ server

`$ docker-compose up -d`

Once this is up and running, you can explore the RabbitMQ management UI at http://localhost:15672 and logging in with the default credentials (use `guest` as the username and password).
