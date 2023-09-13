package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/BearTS/go-echo-template/mailer/pkg"
	"github.com/BearTS/go-echo-template/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var numWorkers int // Number of worker goroutines for parallel processing

var rootCmd = &cobra.Command{
	Use:   "mailer",
	Short: "Start the mailer service",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logger.GetInstance()
		err := godotenv.Load(".env")
		if err != nil {
			logger.Fatal(err)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())

		// Setup a signal channel to handle graceful shutdown
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			select {
			case sig := <-sigCh:
				logger.Infof("Received signal %v. Shutting down gracefully...", sig)
				cancel()
			case <-ctx.Done():
			}
		}()

		var mailer pkg.MailInstance

		mailer.Logger = logger
		smtp_username := os.Getenv("SMTP_USERNAME")
		smtp_password := os.Getenv("SMTP_PASSWORD")
		smtp_host := os.Getenv("SMTP_HOST")
		var smtp_port int
		fmt.Sscanf(os.Getenv("SMTP_PORT"), "%d", &smtp_port)

		mailer.SetCredentials(smtp_username, smtp_password)
		mailer.SetTransportDetails(smtp_host, smtp_port)

		var worker pkg.WorkerService
		worker.HostPort = os.Getenv("RABBITMQ_HOST_PORT")
		worker.QueueName = "mail"

		service := pkg.NewService(logger, &mailer, &worker)
		service.StartConsumer(ctx, numWorkers)

		// Wait for the context to be canceled (e.g., on receiving SIGINT or SIGTERM)
		<-ctx.Done()
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&numWorkers, "workers", "w", 1, "Number of worker goroutines")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
