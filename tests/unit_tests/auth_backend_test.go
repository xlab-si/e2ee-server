package unit_tests

import (
	"github.com/xlab-si/e2ee-server/core/authentication"
	"github.com/xlab-si/e2ee-server/core/redis"
	"github.com/mancabizjak/e2ee-server/settings"
	"github.com/xlab-si/e2ee-server/core/db"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

func TestBackend(t *testing.T) {
	TestingT(t)
}

type AuthenticationBackendTestSuite struct{}

var _ = Suite(&AuthenticationBackendTestSuite{})
var t *testing.T

func (s *AuthenticationBackendTestSuite) SetUpSuite(c *C) {
	os.Setenv("GO_ENV", "tests")
	settings.Init()
	db.Init()
}

func (suite *AuthenticationBackendTestSuite) TestInitJWTAuthenticationBackend(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	c.Assert(authBackend, NotNil)
	c.Assert(authBackend.PublicKey, NotNil)
}

func (suite *AuthenticationBackendTestSuite) TestGenerateToken(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	userUUID := uuid.New()
	username := "testUser"
	accountId := uint(10100)
	tokenString, err := authBackend.GenerateToken(userUUID, username, accountId)

	assert.Nil(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})

	assert.Nil(t, err)
	assert.True(t, token.Valid)
}

func (suite *AuthenticationBackendTestSuite) TestAuthenticate(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	credentials := db.UserCredentials{
		Username: "haku",
		Password: "testing",
     	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)
	user := db.User{
                Username: "haku",
                HashedPassword: string(hashedPassword),
                Uuid:     uuid.New(),
                Token:     "",
        }
	c.Assert(authBackend.Authenticate(user, credentials), Equals, true)
}

func (suite *AuthenticationBackendTestSuite) TestAuthenticateIncorrectPass(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	credentials := db.UserCredentials{
		Username: "haku",
		Password: "testing1",
     	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)
	user := db.User{
                Username: "haku",
                HashedPassword: string(hashedPassword),
                Uuid:     uuid.New(),
                Token:     "",
        }
	c.Assert(authBackend.Authenticate(user, credentials), Equals, false)
}

func (suite *AuthenticationBackendTestSuite) TestAuthenticateIncorrectUsername(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	credentials := db.UserCredentials{
		Username: "Haku",
		Password: "testing",
     	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)
	user := db.User{
                Username: "haku",
                HashedPassword: string(hashedPassword),
                Uuid:     uuid.New(),
                Token:     "",
        }

	c.Assert(authBackend.Authenticate(user, credentials), Equals, false)
}

func (suite *AuthenticationBackendTestSuite) TestLogout(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	userUUID := uuid.New()
	username := "testUser"
	accountId := uint(10100)
	tokenString, err := authBackend.GenerateToken(userUUID, username, accountId)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	err = authBackend.Logout(tokenString, token)
	assert.Nil(t, err)

	redisConn := redis.Connect()
	redisValue, err := redisConn.GetValue(tokenString)
	assert.Nil(t, err)
	assert.NotEmpty(t, redisValue)
}

func (suite *AuthenticationBackendTestSuite) TestIsInBlacklist(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	userUUID := uuid.New()
	username := "testUser"
	accountId := uint(10100)
	tokenString, err := authBackend.GenerateToken(userUUID, username, accountId)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	err = authBackend.Logout(tokenString, token)
	assert.Nil(t, err)

	assert.True(t, authBackend.IsInBlacklist(tokenString))
}

func (suite *AuthenticationBackendTestSuite) TestIsNotInBlacklist(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	assert.False(t, authBackend.IsInBlacklist("1234"))
}
