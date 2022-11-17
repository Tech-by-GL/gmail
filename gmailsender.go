package gmailsender

import (
	"net/smtp"
)

type (
	EmailContent struct {
		Auth struct {
			Username string
			Password string
		}
		Mail struct {
			Subject          string
			Body             string
			SendingAddresses []string
		}
	}
)

func NewGmailContent(mailUsername string, mailPassword string, mails []string, subject string, body string) (email EmailContent) {
	email.Auth.Username = mailUsername
	email.Auth.Password = mailPassword
	email.Mail.Subject = subject
	email.Mail.Body = body
	email.Mail.SendingAddresses = mails
	return email
}

/*
	Calling function NewGmailContent to get content
	this function contains (username, password, mails_sending, title, message_body)
 */
func SendEmail(content EmailContent) (err error) {
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	auth := smtp.PlainAuth("", content.Auth.Username, content.Auth.Password, host)
	message := []byte(content.Mail.Subject + content.Mail.Body)

	return smtp.SendMail(address, auth, content.Auth.Username, content.Mail.SendingAddresses, message)
}

