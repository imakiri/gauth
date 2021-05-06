package gauth

// Key is what a user need in order to pass an Authorization.
// Key can be a random cookie-related string or a JWT itself
type Key string

// Method indicates what Authentication method to use.
// Ex.: password, email, sms, telegram, etc.
// Available methods are exposed through package constants
type Method string

// GroupName defines base subject permissions
type GroupName string

// ObjectID is an object ID to which subject are trying get access
type ObjectID string

// FunctionName is the name of a called by subject function which has something to do with a given object(s)
type FunctionName string

// Identity is one of Authentication pair.
// It can be a login, a phone number, an email address, etc.
type Identity string

// Code is one of Authentication pair.
// It can be a password or a temporal code that you can receive with email, with sms, by direct telegram message, etc.
type Code string

type Registrar interface {
	Begin(keys []Key, groupName GroupName, method Method, identity Identity) ([]Key, error)
	Complete(keys []Key, code Code) ([]Key, error)
}

type Authenticator interface {
	Begin(keys []Key, method Method, identity Identity) ([]Key, error)
	Complete(keys []Key, code Code) ([]Key, error)
}

type Authorizer interface {
	Authorize(keys []Key, objectID ObjectID, functionName FunctionName) error
}

type GroupManager interface {
	Create(keys []Key)
	Modify(keys []Key)
	Delete(keys []Key)
}
