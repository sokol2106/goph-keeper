package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Login         string `json:"login,omitempty"`
	PasswordHash  string `json:"password_hash"`
	EncryptionKey string `json:"encryption_key"`
}

type UserResponse struct {
	PrivateUserKey uuid.UUID `json:"private_user_key"`
	EncryptionKey  string    `json:"encryption_key"`
}

type DataText struct {
	DataTextKey    uuid.UUID `json:"data_text_key,omitempty"`
	PrivateUserKey uuid.UUID `json:"private_user_key,omitempty"`
	Data           string    `json:"data,omitempty"`
}

type DataTextResponse struct {
	DataTextKey uuid.UUID `json:"data_text_key,omitempty"`
	Data        string    `json:"data,omitempty"`
}

type DataBinary struct {
	DataBinaryKey  uuid.UUID `json:"data_binary_key,omitempty"`
	PrivateUserKey uuid.UUID `json:"private_user_key,omitempty"`
	FileName       string    `json:"filename,omitempty"`
	Data           string    `json:"data,omitempty"`
}

type DataBinaryResponse struct {
	DataBinaryKey uuid.UUID `json:"data_binary_key,omitempty"`
	FileName      string    `json:"filename,omitempty"`
	Data          string    `json:"data,omitempty"`
}

type DataCreditCardResponse struct {
	DataCreditCardKey uuid.UUID `json:"data_credit_card_key,omitempty"`
	CardNumber        string    `json:"card_number,omitempty"`
	CardholderName    string    `json:"cardholder_name,omitempty"`
	ExpirationDate    string    `json:"expiration_date,omitempty"`
	CVVHash           string    `json:"cvv_hash,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
}
