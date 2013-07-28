// sendmail project main.go
package main

import (
	"flag"
	"fmt"
	"net/smtp"
	"strings"
)

/*
 *  user : example@example.com login smtp server user
 *  password: xxxxx login smtp server password
 *  host: smtp.example.com:port   smtp.163.com:25
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		fmt.Println("err input,please input:")
		fmt.Println("sendmail [mail@xx.com;mail@yy.com] [subject]")
		return
	}
	to := flag.Arg(0)
	subject := flag.Arg(1)

	//fmt.Println(to)
	//fmt.Println(subject)
	//return
	user := "test@test.com"
	password := "test"
	host := "stmp.test.com:25"

	body := `
    <html>
    <body>
    <h3>
    "Auto send email by golang"
    </h3>
    </body>
    </html>
    `
	fmt.Println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}

}
