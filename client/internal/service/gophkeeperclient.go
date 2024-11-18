package service

type GophKeeperClient struct {
	token Token
}

func NewGophKeeperClient() *GophKeeperClient {
	return &GophKeeperClient{}
}

func (gk *GophKeeperClient) SetToken(token Token) {
	gk.token = token
}
