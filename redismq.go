package redismq

type RedisMQ struct {
	config *Config
	redis *redis.Client
}

func (r *RedisMQ) Connect() {

}

func New(config *Config) *RedisMQ {

	return &RedisMQ{config: config}
}
