package email

import (
	"fmt"
	"github.com/nomadbala/crust/server/pkg/resend"
	"math/rand"
	"time"
)

func GenerateVerificationCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func SendVerificationEmail(receiver string) error {
	verificationCode := GenerateVerificationCode()

	err := resend.SendResendMessage(receiver, verificationCode)
	if err != nil {
		return err
	}

	return nil
}
