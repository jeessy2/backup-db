package notice

// SendMessage interface
type SendMessage interface {
	CanBeSend() bool
	SendMessage(title, message string) error
}
