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

func SendVerificationEmail(receiver, code string) error {
	err := resend.SendResendMessage(receiver, code)
	if err != nil {
		return err
	}

	return nil
}
