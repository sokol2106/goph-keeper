package service

import (
	"client/internal/model"
	"encoding/json"
	"net/http"
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

func (gk *GophKeeperClient) SetCookie(cookie *http.Cookie) {
	gk.cookie = cookie
}

func (gk *GophKeeperClient) GetCookie() *http.Cookie {
	return gk.cookie
}

func (gk *GophKeeperClient) CreateText(text string) (string, error) {
	return "", nil
}

func (gk *GophKeeperClient) GetText(body []byte) (string, error) {
	var dataJson model.DataTextResponse
	err := json.Unmarshal(body, &dataJson)
	if err != nil {
		return "", err
	}

	return dataJson.Data, nil
}

func (gk *GophKeeperClient) CreateCreditCard(text string) (string, error) {
	return "", nil
}

func (gk *GophKeeperClient) GetCreditCard(body []byte) (string, error) {
	return "", nil
}

func (gk *GophKeeperClient) CreateBinary(text string) (string, error) {
	return "", nil
}

func (gk *GophKeeperClient) GetBinary(body []byte) (string, error) {
	return "", nil
}
