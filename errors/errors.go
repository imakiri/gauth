package errors

import "github.com/imakiri/erres"

const (
	FailedToDeleteInternalRecord   erres.Error = "failed to delete identity"
	ValidationFailed               erres.Error = "validation failed"
	InternalValidatorError         erres.Error = "internal validator error"
	PreparationFailed              erres.Error = "preparations failed"
	FailedToCreateInternalData     erres.Error = "failed to create internal data"
	FailedToGenerateNewID          erres.Error = "failed to generate new externalID"
	FailedToDecodeTheKey           erres.Error = "failed to decode the key"
	FailedToEncodeTheKey           erres.Error = "failed to encode the key"
	MethodIdentityPairAlreadyExist erres.Error = "method-externalID pair already exist"
	MethodAlreadyPresentInTheKey   erres.Error = "method already present in the key"
	InvalidMethod                  erres.Error = "invalid authentication method"
	FailedToGenerateRegKey         erres.Error = "failed to generate new key"
	FailedToCheckRegKey            erres.Error = "failed to check new key"
	NoValidRegKeyPresent           erres.Error = "no valid registration key present"
	UnableToRegister               erres.Error = "unable to register"
)
