package adapters

import "fmt"

type EmailMessageSender struct {
}

func (m *EmailMessageSender) Send(content string, contact string) {
	fmt.Println("Sent message ", contact, " to contact ", contact)
}
