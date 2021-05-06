package message

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gauth/internal/cfg"
	"github.com/imakiri/gauth/internal/email"
	"github.com/imakiri/gauth/internal/telegram"
)

type Service interface {
	Send(addr Address, msg []byte) error
}

type service struct {
	email    email.Service
	telegram telegram.Service
}

func NewService(config cfg.ConfigInitMessenger) (Service, error) {
	var s = new(service)
	var err error

	if config.Email != nil {
		switch {
		case config.Email.Basic != nil:
			if s.email, err = email.NewServiceEmailBasic(*config.Email.Basic); err != nil {
				return nil, err
			}
		}
	}

	if config.Telegram != nil {
		switch {
		case config.Telegram.Basic != nil:
			//
		}
	}

	return s, err
}

func (s service) Send(addr Address, msg []byte) error {
	switch a := addr.(type) {
	case email.Address:
		return s.email.Send(a, msg)
	case telegram.Address:
		return s.telegram.Send(a, msg)
	default:
		return erres.TypeMismatch
	}
}
