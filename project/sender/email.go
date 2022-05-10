package sender

import "log"

func NewEmail() Sender {
	log.Println("Email configured")
	return &Email{}
}

func (email *Email) Send(text string) error {
	log.Println("Email sent")
	return nil
}
