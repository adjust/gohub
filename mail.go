package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strconv"
)

type MailSettings struct {
	User     string
	Password string
	Host     string
	From     string
	To       string
	Subject  string
	Port     int
}

func SendErrorReport(opts *MailSettings, body string) {
	auth := smtp.PlainAuth("", opts.User, opts.Password, opts.Host)
	messageBody := fmt.Sprintf("Subject: %s\r\n\r\n%s", opts.Subject, body)

	err := smtp.SendMail(
		opts.Host+":"+strconv.Itoa(opts.Port),
		auth,
		opts.From,
		[]string{opts.To},
		[]byte(messageBody),
	)
	if err != nil {
		log.Printf("[ERROR] Could not send email: %s\n", err)
	}
}
