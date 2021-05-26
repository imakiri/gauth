package gauth

type (
	Key string

	SubjectGroupName  string
	ObjectGroupName   string
	FunctionGroupName string

	SubjectID  string
	ObjectID   string
	FunctionID string

	// AuthenticationMethod indicates what Authentication method to use.
	// Ex.: password, email, sms, telegram, etc.
	// Available methods are exposed through package constants
	AuthenticationMethod string

	// AuthenticationIdentity is one of Authentication pair.
	// It can be a login, a phone number, an email address, etc.
	AuthenticationIdentity string

	// AuthenticationCode is one of Authentication pair.
	// It can be a password or a temporal code that you can receive with email, with sms, by direct telegram message, etc.
	AuthenticationCode string
)

type Registrar interface {
	RegistrationBegin(keys []Key, method AuthenticationMethod, identity AuthenticationIdentity) ([]Key, error)
	RegistrationComplete(keys []Key, code AuthenticationCode) ([]Key, error)
}

type Authenticator interface {
	AuthenticationBegin(keys []Key, method AuthenticationMethod, identity AuthenticationIdentity) ([]Key, error)
	AuthenticationComplete(keys []Key, code AuthenticationCode) ([]Key, error)
}

type Authorizer interface {
	Authorize(keys []Key, oid ObjectID, fid FunctionID) error
}

type SubjectManager interface {
	Create(keys []Key, id SubjectID, name SubjectGroupName) (Key, error)
	Modify(keys []Key, id SubjectID, newName SubjectGroupName) error
	Delete(keys []Key, id SubjectID) error
}

type ObjectManager interface {
	Create(keys []Key, id ObjectID, name ObjectGroupName) (Key, error)
	Modify(keys []Key, id ObjectID, newName ObjectGroupName) error
	Delete(keys []Key, id ObjectID) error
}

type FunctionManager interface {
	Create(keys []Key, id FunctionID, name FunctionGroupName) (Key, error)
	Modify(keys []Key, id FunctionID, newName FunctionGroupName) error
	Delete(keys []Key, id FunctionID) error
}

type SubjectGroupManager interface {
	Create(keys []Key, name SubjectGroupName) (Key, error)
	Modify(keys []Key, name SubjectGroupName) error
	Delete(keys []Key, name SubjectGroupName) error
	AddSubjects(keys []Key, name SubjectGroupName, subjects []SubjectID) (int, error)
	RemoveSubjects(keys []Key, name SubjectGroupName, subjects []SubjectID) (int, error)
}

type ObjectGroupManager interface {
	Create(keys []Key, name ObjectGroupName) (Key, error)
	Modify(keys []Key, name ObjectGroupName) error
	Delete(keys []Key, name ObjectGroupName) error
	AddObject(keys []Key, name ObjectGroupName, objects []ObjectID) (int, error)
	RemoveObject(keys []Key, name ObjectGroupName, objects []ObjectID) (int, error)
}

type FunctionGroupManager interface {
	Create(keys []Key, name FunctionGroupName) (Key, error)
	Modify(keys []Key, name FunctionGroupName) error
	Delete(keys []Key, name FunctionGroupName) error
	AddFunction(keys []Key, name FunctionGroupName, functions []FunctionID) (int, error)
	RemoveFunction(keys []Key, name FunctionGroupName, functions []FunctionID) (int, error)
}

func (s SubjectGroupName) Subject() string {
	return string(s)
}
func (s SubjectID) Subject() string {
	return string(s)
}
func (o ObjectGroupName) Object() string {
	return string(o)
}
func (o ObjectID) Object() string {
	return string(o)
}
func (o FunctionGroupName) Function() string {
	return string(o)
}
func (o FunctionID) Function() string {
	return string(o)
}

type SubjectIdentity interface {
	Subject() string
}
type ObjectIdentity interface {
	Object() string
}
type FunctionIdentity interface {
	Function() string
}

type MSOF struct {
	Method   AuthenticationMethod
	Subject  SubjectIdentity
	Object   ObjectIdentity
	Function FunctionIdentity
}

type PermissionManager interface {
	Allow(keys []Key, msof MSOF) error
	Restrict(keys []Key, msof MSOF) error
	Withdraw(keys []Key, msof MSOF) error
	ListAll(keys []Key) ([]MSOF, error)
}

type Managers interface {
	Permission() PermissionManager
	FunctionGroup() FunctionGroupManager
	ObjectGroup() ObjectGroupManager
	SubjectGroup() SubjectGroupManager
	Function() FunctionManager
	Object() ObjectManager
	Subject() SubjectManager
}

type Service interface {
	Registrar() Registrar
	Authenticator() Authenticator
	Authorizer() Authorizer
	Permission() PermissionManager
}

func NewService() (Service, error) {
	var s Service

	return s, nil
}
