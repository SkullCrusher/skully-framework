package email

import (
	"fmt"
	"net/smtp"
)

// The constant mime.
var mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

/*
	SendEmail
	Send a webhook message to discord.

	@params {string} username  - AWS IAM account username
	@params {string} password  - AWS IAM account password
	@params {string} recipient - The email we want to send the message to.
	@params {string} from      - Who is sending the message.
	@params {string} subject   - The email subject.
	@params {string} body      - The body of the email.

	@returns {error} error - The error that happened during sending.
*/
func SendEmail(username string, password string, recipient string, from string, subject string, body string) error {

	var auth = smtp.PlainAuth("", username, password, "email-smtp.us-east-1.amazonaws.com")

	// Format the message.
	formattedMessage := fmt.Sprintf("To: %v\r\nSubject: %v\r\n%v\r\n%v\r\n", recipient, subject, mime, body)

	// Format the content.
	to  := []string{recipient}
	msg := []byte(formattedMessage)

	// Send out the message.
	err := smtp.SendMail("email-smtp.us-east-1.amazonaws.com:25", auth, from, to, msg)

	// If we had an error, return it.
	if err != nil {
		return err
	}

	return nil
}
