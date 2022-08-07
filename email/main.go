package main

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func main() {
	e := email.NewEmail()
	e.From = "Jcoder <jeekang@qq.com>"
	e.To = []string{"jeekang@qq.com"}
	e.Bcc = []string{"jeekang@qq.com"}
	e.Cc = []string{"jeekang@qq.com"}
	e.Subject = "test golang email"
	e.Text = []byte("test golang email")
	e.HTML = []byte("<h1>test golang email</h1>")
	e.AttachFile("text.txt")

	//e := &email.Email{
	//	From:    "Jcoder <jeekang@qq.com>",
	//	To:      []string{"jeekang@qq.com"},
	//	Subject: "test golang email",
	//	Text:    []byte("test golang email"),
	//	HTML:    []byte("<h1>test golang email</h1>"),
	//	Headers: textproto.MIMEHeader{},
	//}
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "jeekang@qq.com", "AuthC0d1", "smtp.qq.com"))
	if err != nil {
		panic(err)
	}
}
