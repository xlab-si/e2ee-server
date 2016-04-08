package unit_tests

import (
	"github.com/xlab-si/e2ee-server/services"
	"github.com/mancabizjak/e2ee-server/settings"
	"github.com/xlab-si/e2ee-server/core/authentication"
	"github.com/xlab-si/e2ee-server/core/db"
	"github.com/pborman/uuid"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
	 . "gopkg.in/check.v1"
)

func TestService(t *testing.T) {
        TestingT(t)
}

type AuthenticationServiceTestSuite struct{}

var _ = Suite(&AuthenticationServiceTestSuite{})

func (s *AuthenticationServiceTestSuite) SetUpSuite(c *C) {
	os.Setenv("GO_ENV", "tests")
	settings.Init()
	//db.Init() // initialized in auth_backend_test
}

func (s *AuthenticationServiceTestSuite) TestLogin(c *C) {
        credentials := db.UserCredentials{
                Username: "haku",
                Password: "testing",
        }
	response, token := services.Login(credentials)
	assert.Equal(c, http.StatusOK, response)
	assert.NotEmpty(c, token)
}

func (s *AuthenticationServiceTestSuite) TestLoginIncorrectPassword(c *C) {
	credentials := db.UserCredentials{
                Username: "haku",
                Password: "testing1",
        }
	response, token := services.Login(credentials)
	assert.Equal(c, http.StatusUnauthorized, response)
	assert.Empty(c, token)
}

func (s *AuthenticationServiceTestSuite) TestLoginIncorrectUsername(c *C) {
	credentials := db.UserCredentials{
                Username: "Haku",
                Password: "testing",
        }
	response, token := services.Login(credentials)
	assert.Equal(c, http.StatusUnauthorized, response)
	assert.Empty(c, token)
}

func (s *AuthenticationServiceTestSuite) TestLoginEmptyCredentials(c *C) {
	credentials := db.UserCredentials{
                Username: "",
                Password: "",
        }
	response, token := services.Login(credentials)
	assert.Equal(c, http.StatusUnauthorized, response)
	assert.Empty(c, token)
}

func (s *AuthenticationServiceTestSuite) TestRefreshToken(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	userUUID := uuid.New()
        username := "testUser"
        accountId := uint(10100)
        tokenString, err := authBackend.GenerateToken(userUUID, username, accountId)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	assert.NotEmpty(c, token)
	assert.Nil(c, err)

	newToken := services.RefreshToken(userUUID, username, accountId)
	assert.NotEmpty(c, newToken)
}

func (s *AuthenticationServiceTestSuite) TestLogout(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	userUUID := uuid.New()
        username := "testUser"
        accountId := uint(10100)
        tokenString, err := authBackend.GenerateToken(userUUID, username, accountId)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})

	//err = services.Logout(tokenString, token)
	err = authBackend.Logout(tokenString, token)
	assert.Nil(c, err)
}
