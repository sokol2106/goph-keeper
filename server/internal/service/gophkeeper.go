package service

import (
	"encoding/json"
	"github.com/google/uuid"
	"server/internal/model"
)

type Storage interface {
	SelectUser(user model.User) (model.UserResponse, error)
	InsertUser(user model.User) (model.UserResponse, error)

	InsertDataText(data model.DataText) (model.DataTextResponse, error)
	SelectDataText(data model.DataText) (model.DataTextResponse, error)
	DeleteDataText(data model.DataText) error

	InsertDataBinary(data model.DataBinary) (model.DataBinaryResponse, error)
	SelectDataBinary(data model.DataBinary) (model.DataBinaryResponse, error)
	DeleteDataBinary(data model.DataBinary) error

	InsertDataCard(model.DataCreditCard) (model.DataCreditCardResponse, error)
	SelectDataCard(model.DataCreditCard) (model.DataCreditCardResponse, error)
	DeleteDataCard(model.DataCreditCard) error
}

type GophKeeper struct {
	str              Storage
	srvAuthorization *Authorization
}

func NewGophKeeper(str Storage) GophKeeper {
	return GophKeeper{
		str:              str,
		srvAuthorization: NewAuthorization(),
	}

}

func (gk *GophKeeper) GetServiceAuthorization() *Authorization {
	return gk.srvAuthorization
}

func (gk *GophKeeper) RegisterUser(body string) ([]byte, string, error) {
	var strUser model.User
	err := json.Unmarshal([]byte(body), &strUser)
	if err != nil {
		return nil, "", err
	}

	strUser.EncryptionKey, err = GenerateEncryptionKey()
	if err != nil {
		return nil, "", err
	}

	result, err := gk.str.InsertUser(strUser)
	if err != nil {
		return nil, "", err
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, "", err
	}

	token, err := gk.srvAuthorization.NewUserToken(strUser.Login, result.PrivateUserKey.String(), result.EncryptionKey)
	if err != nil {
		return nil, "", err
	}

	return resultBytes, token, nil
}

func (gk *GophKeeper) AuthorizationUser(body string) ([]byte, string, error) {
	var strUser model.User
	err := json.Unmarshal([]byte(body), &strUser)
	if err != nil {
		return nil, "", err
	}

	result, err := gk.str.SelectUser(strUser)
	if err != nil {
		return nil, "", err
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, "", err
	}

	token, err := gk.srvAuthorization.NewUserToken(strUser.Login, result.PrivateUserKey.String(), result.EncryptionKey)
	if err != nil {
		return nil, "", err
	}
	return resultBytes, token, nil
}

func (gk *GophKeeper) LogoutUser() error {
	return nil
}

func (gk *GophKeeper) InsertDataText(body []byte, privateUserKey uuid.UUID) ([]byte, error) {
	var data model.DataText
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	data.PrivateUserKey = privateUserKey
	result, err := gk.str.InsertDataText(data)
	if err != nil {
		return nil, err
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return resultBytes, nil
}

func (gk *GophKeeper) SelectDataText(key string, privateUserKey uuid.UUID) ([]byte, error) {
	var err error
	data := model.DataText{}
	data.PrivateUserKey = privateUserKey
	data.DataTextKey, err = uuid.Parse(key)
	if err != nil {
		return nil, err
	}

	result, err := gk.str.SelectDataText(data)
	if err != nil {
		return nil, err
	}
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return resultBytes, nil
}

func (gk *GophKeeper) DeleteDataText(key string, privateUserKey uuid.UUID) error {
	var err error
	data := model.DataText{}
	data.PrivateUserKey = privateUserKey
	data.DataTextKey, err = uuid.Parse(key)
	if err != nil {
		return err
	}

	err = gk.str.DeleteDataText(data)
	if err != nil {
		return err
	}

	return nil
}

func (gk *GophKeeper) InsertDataBinary(body []byte, privateUserKey uuid.UUID) ([]byte, error) {
	var data model.DataBinary
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	data.PrivateUserKey = privateUserKey
	result, err := gk.str.InsertDataBinary(data)
	if err != nil {
		return nil, err
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return resultBytes, nil
}

func (gk *GophKeeper) SelectDataBinary(key string, privateUserKey uuid.UUID) ([]byte, error) {
	var err error
	data := model.DataBinary{}
	data.PrivateUserKey = privateUserKey
	data.DataBinaryKey, err = uuid.Parse(key)
	if err != nil {
		return nil, err
	}

	result, err := gk.str.SelectDataBinary(data)
	if err != nil {
		return nil, err
	}
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return resultBytes, nil
}

func (gk *GophKeeper) DeleteDataBinary(key string, privateUserKey uuid.UUID) error {
	var err error
	data := model.DataBinary{}
	data.PrivateUserKey = privateUserKey
	data.DataBinaryKey, err = uuid.Parse(key)
	if err != nil {
		return err
	}

	err = gk.str.DeleteDataBinary(data)
	if err != nil {
		return err
	}

	return nil
}

func (gk *GophKeeper) InsertDataCard(body []byte, privateUserKey uuid.UUID) ([]byte, error) {
	var data model.DataCreditCard
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	data.PrivateUserKey = privateUserKey
	result, err := gk.str.InsertDataCard(data)
	if err != nil {
		return nil, err
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return resultBytes, nil
}

func (gk *GophKeeper) SelectDataCard(key string, privateUserKey uuid.UUID) ([]byte, error) {
	var err error
	data := model.DataCreditCard{}
	data.PrivateUserKey = privateUserKey
	data.DataCreditCardKey, err = uuid.Parse(key)
	if err != nil {
		return nil, err
	}

	result, err := gk.str.SelectDataCard(data)
	if err != nil {
		return nil, err
	}
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return resultBytes, nil
}

func (gk *GophKeeper) DeleteDataCard(key string, privateUserKey uuid.UUID) error {
	var err error
	data := model.DataCreditCard{}
	data.PrivateUserKey = privateUserKey
	data.DataCreditCardKey, err = uuid.Parse(key)
	if err != nil {
		return err
	}

	err = gk.str.DeleteDataCard(data)
	if err != nil {
		return err
	}

	return nil
}
