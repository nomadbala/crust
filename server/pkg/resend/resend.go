package resend

import (
	"fmt"
	"github.com/nomadbala/crust/server/pkg/log"
	"github.com/resend/resend-go/v2"
	"go.uber.org/zap"
	"os"
)

var client *resend.Client

func ConfigureResendClient() {
	client = resend.NewClient(os.Getenv("RESEND_API_KEY_2"))
}

func SendResendMessage(receiver, message string) error {
	params := &resend.SendEmailRequest{
		From:    "kiteo@kiteo.app",
		To:      []string{receiver},
		Html:    fmt.Sprintf("<strong>%s</strong>", message),
		Subject: "Hello from Crust",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		log.Logger.Error("error occurred while sending message via resend", zap.Error(err))
		return err
	}

	log.Logger.Info("message successfully delivered", zap.Any("sent_id", sent.Id))

	return nil
}
