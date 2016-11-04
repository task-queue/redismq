package redismq

type Queue struct {
	Name string
	client *RedisMQ
}

func (q Queue) Publish(body []byte) {
	q.client.Publish(q.Name, body)
}
