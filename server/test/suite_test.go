package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"os"
	"server/internal/config"
	"server/internal/handlers"
	"server/internal/model"
	"server/internal/server"
	"server/internal/service"
	"server/internal/storage"
	"strings"
	"testing"
)

type ServerTestSuite struct {
	suite.Suite
	server *httptest.Server
	cookie *http.Cookie
}

func (suite *ServerTestSuite) SetupSuite() {
	cfg := config.NewConfig("", "localhost", "5432", "postgres", "12345678", "gophkeeper")
	objStorage := storage.NewPostgresql(*cfg)
	err := objStorage.Connect()
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}
	gophKeeper := service.NewGophKeeper(objStorage)
	handler := handlers.NewHandlers(&gophKeeper)
	suite.server = httptest.NewServer(server.Router(handler))

}

func (suite *ServerTestSuite) TearSuiteDownSuite() {
	suite.server.Close()
}

func (suite *ServerTestSuite) TestAddUser() {
	user := model.User{
		Login:        "UserSuite",
		PasswordHash: "12345678",
	}
	reqBody, err := json.Marshal(user)
	require.NoError(suite.T(), err)
	resp, err := http.Post(suite.server.URL+"/api/register", "application/json", strings.NewReader(string(reqBody)))
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), http.StatusCreated, resp.StatusCode)

	// Чтение тела ответа
	userResponse := model.UserResponse{}
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	require.NoError(suite.T(), err)
	resp.Body.Close()

	require.Len(suite.T(), resp.Cookies(), 1)
	suite.cookie = resp.Cookies()[0]
}

func (suite *ServerTestSuite) TestText() {
	reqBody := `{"data": "text data test suite"}`

	request, err := http.NewRequest("POST", suite.server.URL+"/api/data/text", strings.NewReader(reqBody))
	require.NoError(suite.T(), err)
	request.AddCookie(suite.cookie)

	client := &http.Client{}
	resp, err := client.Do(request)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), http.StatusCreated, resp.StatusCode)

	// Чтение тела ответа
	userResponse := model.DataTextResponse{}
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	require.NoError(suite.T(), err)
	resp.Body.Close()
}

func (suite *ServerTestSuite) TestBinary() {
	reqBody := `{"filename" : "testfilesuite.json", "data": "dGVzdCB0ZXh0IGRhdGEgMg=="}`

	request, err := http.NewRequest("POST", suite.server.URL+"/api/data/binary", strings.NewReader(reqBody))
	require.NoError(suite.T(), err)
	request.AddCookie(suite.cookie)

	client := &http.Client{}
	resp, err := client.Do(request)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), http.StatusCreated, resp.StatusCode)

	// Чтение тела ответа
	userResponse := model.DataBinaryResponse{}
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	require.NoError(suite.T(), err)
	resp.Body.Close()
}

func (suite *ServerTestSuite) TestCard() {
	reqBody := `{"card_number": "4111111111111111",
				"cardholder_name": "John Doe",
				"expiration_date": "12/24",
				"cvv_hash": "f0eae6c8d6784b243ec1393a74e1ab45"
				}`

	request, err := http.NewRequest("POST", suite.server.URL+"/api/data/card", strings.NewReader(reqBody))
	require.NoError(suite.T(), err)
	request.AddCookie(suite.cookie)

	client := &http.Client{}
	resp, err := client.Do(request)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), http.StatusCreated, resp.StatusCode)

	// Чтение тела ответа
	userResponse := model.DataCreditCardResponse{}
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	require.NoError(suite.T(), err)
	resp.Body.Close()
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
