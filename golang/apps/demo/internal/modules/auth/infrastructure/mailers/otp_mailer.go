package mailers

import (
	"fmt"
	"strconv"
	"time"

	mail2 "github.com/thumbrise/demo/golang-demo/internal/modules/shared/mail"
	"github.com/wneessen/go-mail"
)

type OTPMailer struct {
	config mail2.Config
}

func NewOTPMail(config mail2.Config) *OTPMailer {
	return &OTPMailer{
		config: config,
	}
}

func (m *OTPMailer) Send(email, otp string, expiredAt time.Time) error {
	message := mail.NewMsg()
	if err := message.From(m.config.From); err != nil {
		return err
	}

	if err := message.To(email); err != nil {
		return err
	}

	message.Subject("Auth OTP")

	expiredAtStr := expiredAt.Format(time.RFC3339)
	body := fmt.Sprintf("OTP: %s\nEXPIRED AT: %s\n", otp, expiredAtStr)
	message.SetBodyString(mail.TypeTextPlain, body)

	port, err := strconv.Atoi(m.config.Port)
	if err != nil {
		return err
	}

	client, err := mail.NewClient(
		m.config.Host,
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthNoAuth),
		mail.WithTLSPolicy(mail.NoTLS),
	)
	if err != nil {
		return err
	}

	if err := client.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
