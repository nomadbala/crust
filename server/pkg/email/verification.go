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

func SendVerificationEmail(receiver string) {
	verificationCode := GenerateVerificationCode()
	message := fmt.Sprintf("Your verification code: %s", verificationCode)

	err := resend.SendResendMessage(re)
}
