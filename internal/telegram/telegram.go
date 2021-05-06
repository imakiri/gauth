package telegram

type Address []string

func (e Address) Address() []string {
	return e
}

type Service interface {
	Send(addr Address, msg []byte) error
}
