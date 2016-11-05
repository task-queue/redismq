package broker

type Stub struct{}

func (c Stub) GetConsumerID(queue string) (int64, error) {
	return 0, nil
}

func (c Stub) Push(name string, body []byte) error {
	return nil
}

func (c Stub) Pop(name string) []byte {
	return []byte{}
}

func (c Stub) Ack(name string) {

}
