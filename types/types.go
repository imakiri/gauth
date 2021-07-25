package types

const (
	LogPass ExternalMethod = "logpass"
)

type (
	//
	Key string

	// Result of a key validation
	Result struct {
		InternalID InternalIdentity
		Methods    struct {
			Internal []InternalMethod
			External []ExternalMethod
		}
	}

	// Authentication method. Ex.: password
	// Available methods are exposed through package constants
	InternalMethod string

	// Authentication method. Ex.: email, sms, telegram, etc.
	// Available methods are exposed through package constants
	ExternalMethod string

	// Internal subject ID. InternalID is one of the authentication result
	InternalIdentity string

	// External subject ID. ExternalID is needed to begin authentication.
	// It can be a login, a phone number, an email address, etc.
	ExternalIdentity string

	// Code is needed to complete authentication.
	// It can be a password or a temporal code received via email, sms, telegram, etc.
	Code string
)

type RequestInternal struct {
	Key        Key
	Method     InternalMethod
	ExternalID ExternalIdentity
	Code       Code
}

type RequestExternalBegin struct {
	Key        Key
	Method     ExternalMethod
	ExternalID ExternalIdentity
}

type RequestExternalComplete struct {
	Key  Key
	Code Code
}

func (m InternalMethod) AuthenticationMethod() string {
	return string(m)
}
func (m ExternalMethod) AuthenticationMethod() string {
	return string(m)
}

func (m InternalMethod) ContainsIn(ml []InternalMethod) bool {
	if ml == nil {
		return false
	}

	for i := range ml {
		if ml[i] == m {
			return true
		}
	}

	return false
}
func (m ExternalMethod) ContainsIn(ml []ExternalMethod) bool {
	if ml == nil {
		return false
	}

	for i := range ml {
		if ml[i] == m {
			return true
		}
	}

	return false
}

func (i ExternalIdentity) String() string {
	return string(i)
}
func (i InternalIdentity) String() string {
	return string(i)
}

func (c Code) String() string {
	return string(c)
}

