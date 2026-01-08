package mailer

import (
	"bytes"
	"study-stack/internal/shared/env"

	"github.com/resend/resend-go/v2"
)

func SendVerificationEmail(receiverEmail string, token string) error {
	data := map[string]string{
		"Link": env.Config.BACKEND_URL + "/users/verify?t=" + token,
	}

	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, "email_verification", data); err != nil {
		return err
	}

	params := &resend.SendEmailRequest{
		From:    senderEmail,
		To:      []string{receiverEmail},
		Subject: "Verify Your Email",
		Html:    buf.String(),
	}

	_, err := resendEmailClient.Emails.Send(params)
	return err
}
