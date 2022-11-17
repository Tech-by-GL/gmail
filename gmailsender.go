package main

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

type Person struct {
	ID        string
	BalanceID string
	Title     string
	ClerkName string
}

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
	mime := "MIME-version: 1.0\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(content.Mail.Subject + mime + content.Mail.Body)

	return smtp.SendMail(address, auth, content.Auth.Username, content.Mail.SendingAddresses, message)
}

// func main() {
// 	var (
// 		t   *template.Template
// 		err error
// 		p   []Person
// 	)

// 	t, err = template.ParseFiles("./index.html")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	p = append(p, Person{
// 		ID:        "123",
// 		BalanceID: "ABC",
// 		Title:     "Đóng Tiền",
// 		ClerkName: "Minh Tam",
// 	})
// 	p = append(p, Person{
// 		ID:        "1234",
// 		BalanceID: "ABC",
// 		Title:     "Đóng Tiền",
// 		ClerkName: "Minh Tam",
// 	})
// 	p = append(p, Person{
// 		ID:        "12345",
// 		BalanceID: "ABC",
// 		Title:     "Đóng Tiền",
// 		ClerkName: "Minh Tam",
// 	})
// 	p = append(p, Person{
// 		ID:        "123456",
// 		BalanceID: "ABC",
// 		Title:     "Đóng Tiền",
// 		ClerkName: "Minh Tam",
// 	})
// 	p = append(p, Person{
// 		ID:        "1234567",
// 		BalanceID: "ABC",
// 		Title:     "Đóng Tiền",
// 		ClerkName: "Minh Tam",
// 	})
// 	p = append(p, Person{
// 		ID:        "1234567",
// 		BalanceID: "ABC",
// 		Title:     "Đóng Tiền",
// 		ClerkName: "Minh Tam",
// 	})
// 	p = append(p, Person{
// 		ID:        "12345678",
// 		BalanceID: "ABC",
// 		Title:     "Đóng Tiền",
// 		ClerkName: "Minh Tam",
// 	})
// 	buff := new(bytes.Buffer)
// 	t.Execute(buff, p)
// 	log.Println(buff.String())
// 	// email := NewGmailContent("minhtamtesting@gmail.com", "hoyautrhublylaut", []string{"minhtam30102002@gmail.com"}, "Subject: Hi Trung Kien\n", "Minh Tâm xin thông báo một số thứ"+buff.String())
// 	// err = SendEmail(email)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// 	return
// 	// }
// 	log.Println("Succesfully")
// }
