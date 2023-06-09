package mail

import (
	"log"
	"net/smtp"
	"fmt"
	"html/template"
	"bytes"
)

func SendAuthMail(email, token, hashedPassword, random string) {
	from := "yourMail@gmail.com"
	password := "yourPassword"
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	fmt.Println(hashedPassword)
	subject := "Doğrulama Kodu"
	link := fmt.Sprintf("<a href=\"/accept/register/%s/%s\">buraya</a>", token, random)
	body := fmt.Sprintf("<p>Merhaba,</p><p>Doğrulama kodunuz aşağıdadır:</p><p><b>%s</b></p><p>Lütfen kayıt işleminizi tamamlamak için %s tıklayın.</p>", random, link)

	tmpl, err := template.New("email").Parse(body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	htmlBody := tpl.String()

	message := []byte("From: " + from + "\r\n"+"To: " + email + "\r\n" +"Subject: " + subject + "\r\n" +"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +htmlBody + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("E-posta başarıyla gönderildi.")
}
