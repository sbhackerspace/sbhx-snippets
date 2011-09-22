// Steve Phillips / elimisteve
// 2011.09.21

package main

import "smtp"

func main() {
	const LOGIN = "YOUR_USERNAME@gmail.com"
	const PASSWORD = "YOUR_PASSWORD"

	subject := "Subject: Test Email from Go!\n\n"
    content := "Hello from sendMail.go!"

    smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", LOGIN,
		PASSWORD, "smtp.gmail.com"), LOGIN, []string{"elimisteve@gmail.com"},
		[]byte(subject + content))
}