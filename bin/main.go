package main

import (
	"fmt"

	email "github.com/jakoubek/emaillib"
)

func main() {
	emailer := email.NewClient(
		email.WithRelayhost("smtp.fastmail.com", 465),
		email.WithAuth("your.name@example.com", "123456789abcde", true),
		email.WithSender("Example Inc. Customer Service", "info@example.com"),
		email.WithDontSend(),
	)

	emailer.To("John Doe", "jd@example.com")
	emailer.CC("", "sales@example.com")

	emailer.Subject("Some message")

	fmt.Println(emailer.Debug())

	err := emailer.Sendmail("John Doe", "jd@example.com", "Message", "Body of message")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("OK!")
	}

}
