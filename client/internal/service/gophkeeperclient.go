package service

import (
	"client/internal/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type GophKeeperClient struct {
	token  Token
	cookie *http.Cookie
}

func NewGophKeeperClient() *GophKeeperClient {
	return &GophKeeperClient{}
}

func (gk *GophKeeperClient) SetToken(token Token) {
	gk.token = token
}

func (gk *GophKeeperClient) GetToken() Token {
	return gk.token
}

func (gk *GophKeeperClient) SetCookie(cookie *http.Cookie) error {
	var err error
	gk.cookie = cookie
	gk.token, err = ReadToken(cookie.Value)
	if err != nil {
		return err
	}
	return nil
}

func (gk *GophKeeperClient) GetCookie() *http.Cookie {
	return gk.cookie
}

func (gk *GophKeeperClient) CreateText(text string) (string, error) {
	return encrypt(text, gk.token.EncryptionKey)
}

func (gk *GophKeeperClient) GetText(body []byte) (string, error) {
	var dataJson model.DataTextResponse
	err := json.Unmarshal(body, &dataJson)
	if err != nil {
		return "", err
	}
	return decrypt(dataJson.Data, gk.token.EncryptionKey)
}

func (gk *GophKeeperClient) CreateCreditCard(card model.DataCreditCardResponse) (model.DataCreditCardResponse, error) {
	cardNumber, err := encrypt(card.CardNumber, gk.token.EncryptionKey)
	if err != nil {
		return model.DataCreditCardResponse{}, err
	}
	cvv, err := encrypt(card.CVVHash, gk.token.EncryptionKey)
	if err != nil {
		return model.DataCreditCardResponse{}, err
	}
	card.CardNumber = cardNumber
	card.CVVHash = cvv
	return card, nil
}

func (gk *GophKeeperClient) GetCreditCard(body []byte) (model.DataCreditCardResponse, error) {
	var dataJson model.DataCreditCardResponse
	err := json.Unmarshal(body, &dataJson)
	if err != nil {
		return model.DataCreditCardResponse{}, err
	}

	dataJson.CardNumber, err = decrypt(dataJson.CardNumber, gk.token.EncryptionKey)
	if err != nil {
		return model.DataCreditCardResponse{}, err
	}
	dataJson.CVVHash, err = decrypt(dataJson.CVVHash, gk.token.EncryptionKey)
	if err != nil {
		return model.DataCreditCardResponse{}, err
	}
	return dataJson, nil
}

// CreateBinary принимает имя файла и путь, возвращая DataBinary с Base64-кодированными данными файла
func (gk *GophKeeperClient) CreateBinary(fileName, filePath string) (model.DataBinary, error) {
	newFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return model.DataBinary{}, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer newFile.Close()

	fileContent, err := io.ReadAll(newFile)
	if err != nil {
		return model.DataBinary{}, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	encodedData := base64.StdEncoding.EncodeToString(fileContent)
	return model.DataBinary{
		FileName: fileName,
		Data:     encodedData,
	}, nil
}

func (gk *GophKeeperClient) GetBinary(body []byte) (string, error) {
	var dataJson model.DataBinaryResponse
	err := json.Unmarshal(body, &dataJson)
	if err != nil {
		return "", err
	}

	decodedData, err := base64.StdEncoding.DecodeString(dataJson.Data)
	if err != nil {
		return "", err
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	outputPath := filepath.Join(workingDir, dataJson.FileName)
	newFile, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return "", err
	}
	defer newFile.Close()
	_, err = newFile.Write(decodedData)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
