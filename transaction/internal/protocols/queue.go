package protocols

type Connection interface {
	Close()
}

type Publisher interface {
	Publish(subject string, data []byte) error
}

type Subscriber interface {
	Subscribe(subject string, callback func(data []byte)) error
}
