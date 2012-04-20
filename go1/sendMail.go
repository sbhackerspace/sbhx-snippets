// Steve Phillips / elimisteve
// 2011.09.21

package main

import (
	"fmt"
	"os"
	"net/smtp"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: sendMail.go YOUR_USERNAME@gmail.com YOUR_PASSWORD\n")
		return
	}
	LOGIN := os.Args[1]
	PASSWORD := os.Args[2]

	subject := "Subject: Test Email from Go!"
    content := "Hello from sendMail.go!"
	auth := smtp.PlainAuth("", LOGIN, PASSWORD, "smtp.gmail.com")
	recipients := []string{"elimisteve@gmail.com", "sbhackerspace@gmail.com"}

    err := smtp.SendMail("smtp.gmail.com:587", auth, LOGIN, recipients,
		[]byte(subject + "\n\n" + content))
	if err == nil { 
		fmt.Printf("Email sent successfully to %v!\n",
			strings.Join(recipients, ", "))
	} else {
		fmt.Printf("Error sending email: %v\n", err)
	}
}