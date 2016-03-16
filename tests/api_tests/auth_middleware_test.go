package api_tests

import (
	"e2ee/core/authentication"
	"e2ee/routers"
	"e2ee/services"
	"e2ee/settings"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"github.com/pborman/uuid"
	"os"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MiddlewaresTestSuite struct{}

var _ = Suite(&MiddlewaresTestSuite{})
var t *testing.T
var token string
var server *negroni.Negroni

func (s *MiddlewaresTestSuite) SetUpSuite(c *C) {
	os.Setenv("GO_ENV", "tests")
	settings.Init()

	authBackend := authentication.InitJWTAuthenticationBackend()
	assert.NotNil(t, authBackend)
	userUUID := uuid.New()
        username := "testUser"
        accountId := uint(10100)
        token, _ = authBackend.GenerateToken(userUUID, username, accountId)

	router := routers.InitRoutes()
	server = negroni.Classic()
	server.UseHandler(router)
}

func (s *MiddlewaresTestSuite) SetUpTest(c *C) {

}

func (s *MiddlewaresTestSuite) TestRequireTokenAuthentication(c *C) {
	resource := "/test/hello"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	server.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusOK)
}

func (s *MiddlewaresTestSuite) TestRequireTokenAuthenticationInvalidToken(c *C) {
	resource := "/test/hello"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", "token"))
	server.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusUnauthorized)
}

func (s *MiddlewaresTestSuite) TestRequireTokenAuthenticationEmptyToken(c *C) {
	resource := "/test/hello"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", ""))
	server.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusUnauthorized)
}

func (s *MiddlewaresTestSuite) TestRequireTokenAuthenticationWithoutToken(c *C) {
	resource := "/test/hello"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	server.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusUnauthorized)
}

func (suite *MiddlewaresTestSuite) TestRequireTokenAuthenticationAfterLogout(c *C) {
	resource := "/test/hello"

	requestLogout, _ := http.NewRequest("GET", resource, nil)
	requestLogout.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	services.Logout(requestLogout)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	server.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusUnauthorized)
}
