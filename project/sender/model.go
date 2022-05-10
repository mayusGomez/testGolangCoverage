package sender

type Sender interface {
	Send(text string) error
}

type Email struct {
}
