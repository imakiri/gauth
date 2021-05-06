package email

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gauth/internal/cfg"
	"net/smtp"
)

func isEmail(addr string) bool {

	// TODO: Email validator

	return true
}

type Address struct {
	addr            []string
	properlyCreated bool
}

func (e Address) Address() []string {
	return e.addr
}

type Service interface {
	Send(email Address, msg []byte) error
}

func NewAddress(addr []string) (Address, error) {
	var e Address
	e.addr = make([]string, len(addr))

	for i := range addr {
		if isEmail(addr[i]) {
			e.addr = append(e.addr, addr[i])
		} else {
			return e, erres.TypeMismatch
		}
	}

	e.properlyCreated = true
	return e, nil
}

func NewAddressWithoutValidation(addr []string) (Address, error) {
	var e Address
	e.addr = addr
	e.properlyCreated = true

	return e, nil
}

type ServiceEmailBasic struct {
	addr string
	auth smtp.Auth
	from string
}

func (e ServiceEmailBasic) Send(addr Address, msg []byte) error {
	var err error

	err = smtp.SendMail(e.addr, e.auth, e.from, addr.Address(), msg)
	if err != nil {
		return erres.InternalServiceError.Extend().
			AddRoute("message").AddRoute("ServiceEmailBasic").AddRoute("Send").AddDescription(err.Error())
	}

	return err
}

func NewServiceEmailBasic(config cfg.ConfigInitMessengerEmailBasic) (*ServiceEmailBasic, error) {
	var err error
	if !isEmail(config.Address) {
		return nil, erres.TypeMismatch.Extend().AddRoute("message").AddRoute("internal").AddRoute("NewServiceEmailBasic").
			AddDescription("incorrect email address").AddDescription(config.Address)
	}

	var email ServiceEmailBasic

	email.auth = smtp.PlainAuth(config.Identity, config.User, config.Password, config.Address)
	email.from = config.From

	return &email, err
}
