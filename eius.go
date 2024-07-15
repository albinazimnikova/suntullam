package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@example.com")
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", "Links to Other Websites")
	m.SetBody("text/html", `
		<p>Here are some useful links:</p>
		<ul>
			<li><a href="https://www.example.com">example</a></li>
			<li><a href="https://www.example.com">example</a></li>
			<li><a href="https://www.example.com">example</a></li>
		</ul>`)

	d := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-password")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
