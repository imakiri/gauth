package types

type AuthenticationMethod interface {
	AuthenticationMethod() string
}

type Authenticator interface {
	Authenticate(key Key) (bool, Result, error)
}

type Manager interface {
	InternalRegistration(i RequestInternal) (Key, error)
	InternalAuthentication(i RequestInternal) (Key, error)
	InternalDeletion(i RequestInternal) (Key, error)

	ExternalBeginRegistration(eb RequestExternalBegin) (Key, error)
	ExternalBeginAuthentication(eb RequestExternalBegin) (Key, error)
	ExternalBeginDeletion(eb RequestExternalBegin) (Key, error)

	ExternalCompleteRegistration(ec RequestExternalComplete) (Key, error)
	ExternalCompleteAuthentication(ec RequestExternalComplete) (Key, error)
	ExternalCompleteDeletion(ec RequestExternalComplete) (Key, error)
}

// Key service is about key encoding and decoding in a way
// that only this service could be able to perform this operations
type KeyService interface {
	//
	Encode(key Key, method AuthenticationMethod, identity InternalIdentity) (Key, error)

	// The Bool is true if the Result is valid and can be used.
	// Not-nil error means an internal error has occurred
	Decode(key Key) (*Result, bool, error)

	// Returns modified key with an temporal registration information
	// Not-nil error means an internal error has occurred
	EncodeRegistrationKey(key Key, method ExternalMethod, identity ExternalIdentity) (Key, error)

	// Temporal registration information will be removed from the key if decoding is successful.
	// The Bool is true if there is a useful result
	// Not-nil error means an internal error has occurred
	DecodeRegistrationKey(key Key) (ok bool, k Key, method ExternalMethod, identity ExternalIdentity, err error)

	// Returns modified key with an temporal authentication information
	// Not-nil error means an internal error has occurred
	EncodeAuthenticationKey(key Key, method ExternalMethod, identity ExternalIdentity) (Key, error)

	// Temporal authentication information will be removed from the key if decoding is successful.
	// The Bool is true if there is a useful result
	// Not-nil error means an internal error has occurred
	DecodeAuthenticationKey(key Key) (ok bool, k Key, method ExternalMethod, identity ExternalIdentity, err error)
}

type ExternalValidator interface {
	BeginValidation(identity ExternalIdentity) error
	CompleteValidation(identity ExternalIdentity, code Code) (bool, error)
}

type InternalValidator interface {
	Validate(identity ExternalIdentity, code Code) (bool, error)
}

type DataService interface {
	NewID() (InternalIdentity, error)
	DeleteID(identity InternalIdentity) error
	InternalCreate(method InternalMethod, internalID InternalIdentity, externalID ExternalIdentity, code Code) (bool, error)
	InternalDelete(method InternalMethod, internalID InternalIdentity) error
}

