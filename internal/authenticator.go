package internal

import "github.com/imakiri/gauth/types"

type authenticatorService struct {
	key types.Key
}

func (a authenticatorService) Authenticate(key types.Key) (bool, types.Result, error) {
	panic("implement me")
}

func NewAuthenticator() (types.Authenticator, error) {
	var a = new(authenticatorService)
	panic("implement me")
	return a, nil
}