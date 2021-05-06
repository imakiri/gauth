package message

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gauth/internal/email"
)

func isEmail(addr []string) bool {

	return false
}

func isTelegram(addr []string) bool {

	return false
}

type Address interface {
	Address() []string
}

func NewAddress(addr []string) (Address, error) {
	switch {
	case isEmail(addr):
		return email.NewAddressWithoutValidation(addr)
	case isTelegram(addr):
		return nil, nil
	default:
		return nil, erres.TypeMismatch
	}
}
