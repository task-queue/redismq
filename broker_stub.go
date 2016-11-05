package redismq

type Stub struct{}

func (c Stub) Push(name string, body []byte) error {
	return nil
}

func (c Stub) InitConsumer(queue string) []byte {
	return nil
}

func (c Stub) Pop() []byte {
	return []byte{}
}

func (c Stub) Ack() {

}
