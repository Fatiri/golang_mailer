package helper

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

const (
	SMTPServer = "smtp.gmail.com"
)

func SendMail(Dest []string, Subject, bodyMessage string, user map[int]string) {

	msg := "From: " + user[1] + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", user[1], user[2], SMTPServer),
		user[1], Dest, []byte(msg))

	if err != nil {

		fmt.Printf("smtp error: %s", err)
		return
	}

	fmt.Println("Mail sent successfully!")
}

func WriteEmail(dest []string, contentType string, subject string, bodyMessage string, user map[int]string) string {

	header := make(map[string]string)
	header["From"] = user[1]

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}
