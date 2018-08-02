package main

import (
	"bytes"
	"html/template"
	"net/smtp"
)

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func CreateNewRequest(to []string, subject string, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (req *Request) SendEmail(auth smtp.Auth, user string) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := req.subject
	msg := []byte(subject + mime + "\n" + req.body)
	address := "smtp.gmail.com:587"

	if err := smtp.SendMail(address, auth, user, req.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (req *Request) ParseTemplate(templateFile string, data interface{}) error {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	req.body = buf.String()
	return nil
}
