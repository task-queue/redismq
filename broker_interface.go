package redismq

type IBroker interface {
	Push(string, []byte) error

	InitConsumer(string) error
	Pop() ([]byte, error)
	Ack()
}
