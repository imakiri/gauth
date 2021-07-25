package key

import (
	"github.com/imakiri/gauth/types"
)

type keyService struct {

}

func (s keyService) Encode(key types.Key, method types.AuthenticationMethod, identity types.InternalIdentity) (types.Key, error) {
	panic("implement me")
}

func (s keyService) Decode(key types.Key) (*types.Result, bool, error) {
	panic("implement me")
}

func (s keyService) EncodeRegistrationKey(key types.Key, method types.ExternalMethod, identity types.ExternalIdentity) (types.Key, error) {
	panic("implement me")
}

func (s keyService) DecodeRegistrationKey(key types.Key) (ok bool, k types.Key, method types.ExternalMethod, identity types.ExternalIdentity, err error) {
	panic("implement me")
}

func (s keyService) EncodeAuthenticationKey(key types.Key, method types.ExternalMethod, identity types.ExternalIdentity) (types.Key, error) {
	panic("implement me")
}

func (s keyService) DecodeAuthenticationKey(key types.Key) (ok bool, k types.Key, method types.ExternalMethod, identity types.ExternalIdentity, err error) {
	panic("implement me")
}

func NewKeyService() (*keyService, error) {
	var ks = new(keyService)
	return ks, nil
}

