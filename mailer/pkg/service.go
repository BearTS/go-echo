package pkg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BearTS/go-echo/pkg/logger"
	"github.com/pkg/errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Service struct {
	Logger logger.Logger
	Mailer *MailInstance
	Worker *WorkerService
}

type WorkerService struct {
	HostPort  string
	QueueName string
}

func NewService(logger logger.Logger, mailer *MailInstance, worker *WorkerService) *Service {
	return &Service{
		Logger: logger,
		Mailer: mailer,
		Worker: worker,
	}
}

func (svc *Service) StartConsumer(ctx context.Context, numWorkers int) {
	conn, err := amqp.Dial(svc.Worker.HostPort)
	if err != nil {
		svc.Logger.Fatal(errors.Wrap(err, "failed to connect to RabbitMQ"))
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		svc.Logger.Fatal(errors.Wrap(err, "failed to open a channel"))
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		svc.Worker.QueueName, // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		svc.Logger.Fatal(errors.Wrap(err, "failed to declare a queue"))
		return
	}

	ch.Qos(
		numWorkers, // Set the maximum number of unacknowledged messages to the number of workers.
		0,
		false,
	)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		svc.Logger.Fatal(errors.Wrap(err, "failed to register a consumer"))
		return
	}

	svc.Logger.Info(" [*] Waiting for messages. To exit press CTRL+C")

	// Start worker goroutines to process messages concurrently
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for delivery := range msgs {
				var msg Message
				if err := json.Unmarshal(delivery.Body, &msg); err != nil {
					svc.Logger.Error(errors.Wrap(err, "failed to unmarshal message"))
					continue
				}
				svc.Logger.Info(fmt.Sprintf("Worker %d received a message: %s", workerID, msg))
				if err := svc.Mailer.SendEmail(msg); err != nil {
					svc.Logger.Error(errors.Wrap(err, "failed to send email"))
					continue
				}
				svc.Logger.Info(fmt.Sprintf("Worker %d finished processing message: %s", workerID, msg))
				delivery.Ack(false) // Acknowledge the message
			}
		}(i)
	}

	// Block and wait for signals to gracefully shutdown the consumer
	<-ctx.Done()
}
