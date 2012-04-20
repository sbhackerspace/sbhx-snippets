// Steve Phillips / elimisteve
// 2011.09.21

package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	const LOGIN = "YOUR_USERNAME@gmail.com"
	const PASSWORD = "YOUR_PASSWORD"

	subject := "Subject: Test Email from Go!"
    content := "Hello from sendMail.go!"
	auth := smtp.PlainAuth("", LOGIN, PASSWORD, "smtp.gmail.com")
	recipients := []string{"elimisteve@gmail.com", "sbhackerspace@gmail.com"}

    err := smtp.SendMail("smtp.gmail.com:587", auth, LOGIN, recipients,
		[]byte(subject + "\n\n" + content))
	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
	}
}