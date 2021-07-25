package internal

import (
	"github.com/imakiri/gauth/errors"
	"github.com/imakiri/gauth/types"
	"github.com/imakiri/gorum/cfg"
)

type managerService struct {
	key                types.KeyService
	data               types.DataService
	internalValidators map[types.InternalMethod]types.InternalValidator
	externalValidators map[types.ExternalMethod]types.ExternalValidator
}

func (s managerService) prepareInternal(i types.RequestInternal) (types.InternalValidator, *types.Result, bool, error) {
	var internal, okk = s.internalValidators[i.Method]
	if !okk {
		return nil, nil, false, errors.InvalidMethod
	}

	var result, ok, err = s.key.Decode(i.Key)
	if err != nil {
		return nil, nil, false, errors.FailedToDecodeTheKey.ExtendAndLink(0, err)
	}
	if ok && i.Method.ContainsIn(result.Methods.Internal) {
		return nil, nil, false, errors.MethodAlreadyPresentInTheKey
	}

	return internal, result, ok, nil
}

func (s managerService) prepareExternalBegin(eb types.RequestExternalBegin) (error) {
	panic("implement me")
}

func (s managerService) prepareExternalComplete(ec types.RequestExternalComplete) (error) {
	panic("implement me")
}

func (s managerService) InternalRegistration(i types.RequestInternal) (types.Key, error) {
	var validator, result, needNewID, err = s.prepareInternal(i)
	if err != nil {
		return i.Key, errors.PreparationFailed.ExtendAndLink(0, err)
	}

	var ok bool
	ok, err = validator.Validate(i.ExternalID, i.Code)
	if err != nil {
		return i.Key, errors.InternalValidatorError.ExtendAndLink(0, err)
	}
	if !ok {
		return i.Key, errors.ValidationFailed
	}
	if !needNewID {
		result.InternalID, err = s.data.NewID()
		if err != nil {
			return i.Key, errors.FailedToGenerateNewID.ExtendAndLink(0, err)
		}
	}

	ok, err = s.data.InternalCreate(i.Method, result.InternalID, i.ExternalID, i.Code)
	if err != nil {
		return i.Key, errors.FailedToCreateInternalData.ExtendAndLink(0, err)
	}
	if !ok {
		return i.Key, errors.MethodIdentityPairAlreadyExist
	}

	var k types.Key
	k, err = s.key.Encode(i.Key, i.Method, result.InternalID)
	if err != nil {
		err = errors.FailedToEncodeTheKey.ExtendAndLink(0, err)
		var e = s.data.InternalDelete(i.Method, result.InternalID)
		if e != nil {
			err = errors.FailedToDeleteInternalRecord.ExtendAndLink(0, err).SetDescription(e.Error())
		}
		return i.Key, err
	}

	return k, nil
}

func (s managerService) InternalAuthentication(i types.RequestInternal) (types.Key, error) {
	//var validator, result, needNewID, err = s.prepareInternal(i)
	//if err != nil {
	//	return i.key, errors.PreparationFailed.ExtendAndLink(0, err)
	//}

	panic("implement me")
}

func (s managerService) InternalDeletion(i types.RequestInternal) (types.Key, error) {
	//var validator, result, needNewID, err = s.prepareInternal(i)
	//if err != nil {
	//	return i.key, errors.PreparationFailed.ExtendAndLink(0, err)
	//}

	panic("implement me")
}

func (s managerService) ExternalBeginRegistration(eb types.RequestExternalBegin) (types.Key, error) {
	panic("implement me")
}

func (s managerService) ExternalBeginAuthentication(eb types.RequestExternalBegin) (types.Key, error) {
	panic("implement me")
}

func (s managerService) ExternalBeginDeletion(eb types.RequestExternalBegin) (types.Key, error) {
	panic("implement me")
}

func (s managerService) ExternalCompleteRegistration(ec types.RequestExternalComplete) (types.Key, error) {
	panic("implement me")
}

func (s managerService) ExternalCompleteAuthentication(ec types.RequestExternalComplete) (types.Key, error) {
	panic("implement me")
}

func (s managerService) ExternalCompleteDeletion(ec types.RequestExternalComplete) (types.Key, error) {
	panic("implement me")
}

func NewManagerService(key types.KeyService, data types.DataService, config cfg.Service) (*managerService, error) {
	var service = new(managerService)
	service.key = key
	service.data = data

	//service.config = config
	//
	//var err error
	//var c *cfg.ConfigDatabasePostgres
	//var s *cfg.SecretDatabasePostgres
	//
	//if c, err = service.config.GetPostgresConfig(name); err != nil {
	//	return nil, err
	//}
	//if s, err = service.config.GetPostgresSecret(name); err != nil {
	//	return nil, err
	//}
	//
	//var conn *postgres.Connection
	//if conn, err = postgres.NewConnection(nil, c, s); err != nil || conn == nil {
	//	return nil, err
	//}
	//
	//service.connection = *conn
	return service, nil
}

