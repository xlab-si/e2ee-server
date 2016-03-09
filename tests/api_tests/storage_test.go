package api_tests

import (
	"e2ee/core/authentication"
	"e2ee/routers"
	//"e2ee/services"
	"e2ee/settings"
	"e2ee/core/db"
	"e2ee/controllers"
	"e2ee/api/parameters"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"code.google.com/p/go-uuid/uuid"
	"os"
	"bytes"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"testing"
	"log"
)

func Test1(t1 *testing.T) {
	// TestingT should be called only once and it is called
	// already in auth_middleware_test
	//TestingT(t1) 
}

type StorageTestSuite struct{}

var _ = Suite(&StorageTestSuite{})
var t1 *testing.T
var token1 string
var token2 string
var token3 string
var user1 string
var user2 string
var server1 *negroni.Negroni
var containerNameHmac string
var ciphertext string
var payloadciphertext string

func (s *StorageTestSuite) SetUpSuite(c *C) {
	os.Setenv("GO_ENV", "tests")
	settings.Init()
	db.Init()

	authBackend := authentication.InitJWTAuthenticationBackend()
	assert.NotNil(t, authBackend)
	router := routers.InitRoutes()
	server1 = negroni.Classic()
	server1.UseHandler(router)

	containerNameHmac = "hmacfoo"
	ciphertext = "ciphertext"
	payloadciphertext = "payloadciphertext"

	userUUID := uuid.New()
        accountId := uint(1234123)
        token1, _ = authBackend.GenerateToken(userUUID, "testUser", accountId)

	user1 = "foo"
	pass := "foobar"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
	user := db.User{
                Username: user1,
                HashedPassword: string(hashedPassword),
                Uuid:     uuid.New(),
                Token:     "",
        }

	dbP, err := gorm.Open("postgres", "host=172.17.0.2 dbname=e2ee user=postgres sslmode=disable")
        if err != nil {
                log.Fatal(err)
        }
	u := db.GetUser(user1)
	if u.Username == "" {
		dbP.Save(&user)
	}
	credentials := db.UserCredentials{
                Username: user1,
                Password: pass,
        }
	jsonStr, _ := json.Marshal(credentials)

	resource := "/token-auth"
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", resource, bytes.NewBuffer(jsonStr))

	server1.ServeHTTP(response, request)
	var token parameters.TokenAuthentication
	_ = json.Unmarshal(response.Body.Bytes(), &token)
	token2 = token.Token

	var account = db.Account{
		ContainerNameHmacKeyCiphertext: "",
        	HmacKeyCiphertext: "",
        	KeypairCiphertext: "",
        	KeypairMac: "",
        	KeypairMacSalt: "",
        	KeypairSalt: "",
        	PubKey: "",
        	SignKeyPrivateCiphertext: "",
        	SignKeyPrivateMac: "",
        	SignKeyPrivateMacSalt: "",
        	SignKeyPub: "",
        	//Username: "blafoo", // retrieved from token
	}
	resource = "/account"
	response = httptest.NewRecorder()

	jsonStr, _ = json.Marshal(account)

	request, _ = http.NewRequest("POST", resource, bytes.NewBuffer(jsonStr) )
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)

	// another user (for sharing):
	user2 = "bar"
	pass = "foobar"
	hashedPassword, _ = bcrypt.GenerateFromPassword([]byte(pass), 10)
	user = db.User{
                Username: user2,
                HashedPassword: string(hashedPassword),
                Uuid:     uuid.New(),
                Token:     "",
        }

	u = db.GetUser(user2)
	if u.Username == "" {
		dbP.Save(&user)
	}
	credentials = db.UserCredentials{
                Username: user2,
                Password: pass,
        }
	jsonStr, _ = json.Marshal(credentials)

	resource = "/token-auth"
	response = httptest.NewRecorder()
	request, _ = http.NewRequest("POST", resource, bytes.NewBuffer(jsonStr))

	server1.ServeHTTP(response, request)
	_ = json.Unmarshal(response.Body.Bytes(), &token)
	token3 = token.Token

	account = db.Account{
		ContainerNameHmacKeyCiphertext: "",
        	HmacKeyCiphertext: "",
        	KeypairCiphertext: "",
        	KeypairMac: "",
        	KeypairMacSalt: "",
        	KeypairSalt: "",
        	PubKey: "",
        	SignKeyPrivateCiphertext: "",
        	SignKeyPrivateMac: "",
        	SignKeyPrivateMacSalt: "",
        	SignKeyPub: "",
        	//Username: "blafoo", // retrieved from token
	}
	resource = "/account"
	response = httptest.NewRecorder()

	jsonStr, _ = json.Marshal(account)

	request, _ = http.NewRequest("POST", resource, bytes.NewBuffer(jsonStr) )
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token3))
	server1.ServeHTTP(response, request)

}

func (s *StorageTestSuite) SetUpTest(c *C) {

}

func (s *StorageTestSuite) TestAccountNotExists(c *C) {
	// token was generated for testUser (see SetUpTest, token1), but
	// no account was generated and stored
	resource := "/accountexists"
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token1))
	server1.ServeHTTP(response, request)
	var res controllers.AccountExistsResponse
	_ = json.Unmarshal(response.Body.Bytes(), &res)

	assert.Equal(t1, res.Exists, false)
}

func (s *StorageTestSuite) TestAccountExists(c *C) {
	// token was generated for foofoo (see SetUpTest, token2) and
	// account was generated and stored (TestSaveAccount)
	resource := "/accountexists"
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var res controllers.AccountExistsResponse
	_ = json.Unmarshal(response.Body.Bytes(), &res)

	assert.Equal(t1, res.Exists, true)
}

func (s *StorageTestSuite) TestGetAccount(c *C) {
	resource := "/account"
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var res controllers.AccountResponseMessage
	_ = json.Unmarshal(response.Body.Bytes(), &res)

	assert.Equal(t1, res.Success, true)
}

func (s *StorageTestSuite) TestContainerCreate(c *C) {
	resource := "/container/" + containerNameHmac

	var chunk = controllers.ContainerCreateChunk{
		ToAccountId: 0,
		SessionKeyCiphertext: "",
	}

	response := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(chunk)
	request, _ := http.NewRequest("PUT", resource, bytes.NewBuffer(jsonStr) )

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var res controllers.ContainerResponseMessage
	_ = json.Unmarshal(response.Body.Bytes(), &res)

	assert.Equal(t1, res.Success, true)
}

func (s *StorageTestSuite) TestContainerRecordCreate(c *C) {
	resource := "/container/record"

	var chunk = controllers.RecordCreateChunk{
		ContainerNameHmac: containerNameHmac,
		PayloadCiphertext: ciphertext,
	}

	response := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(chunk)
	request, _ := http.NewRequest("POST", resource, bytes.NewBuffer(jsonStr) )

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var res controllers.ContainerResponseMessage
	_ = json.Unmarshal(response.Body.Bytes(), &res)

	assert.Equal(t1, res.Success, true)
}

func (s *StorageTestSuite) TestContainerSGet(c *C) {
	// it seems tests are executed in alphabetical order - so I renamed
	// this test to be after TestContainerRecordCreate
	resource := "/container/" + containerNameHmac

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var res controllers.ContainerResponseMessage
	_ = json.Unmarshal(response.Body.Bytes(), &res)

	assert.Equal(t1, res.Success, true)
	assert.Equal(t1, res.Records[0].PayloadCiphertext, ciphertext)
}

func (s *StorageTestSuite) TestContainerShare(c *C) {
	// it seems tests are executed in alphabetical order
	resource := "/peer/" + user2

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var res controllers.PeerMessage
	_ = json.Unmarshal(response.Body.Bytes(), &res)
	assert.Equal(t1, res.Success, true)
	var accountId = res.Peer.AccountId

	resource = "/container/share"
	var chunk = controllers.ContainerShareChunk{
		ContainerNameHmac: containerNameHmac,
		ToAccountId: accountId,
		SessionKeyCiphertext: ciphertext,
	}

	response = httptest.NewRecorder()
	jsonStr, _ := json.Marshal(chunk)
	request, _ = http.NewRequest("POST", resource, bytes.NewBuffer(jsonStr) )

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var shareRes controllers.ContainerResponseMessage
	_ = json.Unmarshal(response.Body.Bytes(), &shareRes)
	assert.Equal(t1, shareRes.Success, true)

	resource = "/peer"
	var chunk1 = controllers.Notification{
		FromUsername: user1,
		ToAccountId: accountId,
		HeadersCiphertext: "headersciphertext",
		PayloadCiphertext: payloadciphertext,
	}

	response = httptest.NewRecorder()
	jsonStr, _ = json.Marshal(chunk1)
	request, _ = http.NewRequest("POST", resource, bytes.NewBuffer(jsonStr) )

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var notificationRes controllers.NotificationResponse
	_ = json.Unmarshal(response.Body.Bytes(), &notificationRes)
	assert.Equal(t1, notificationRes.Success, true)

}

func (s *StorageTestSuite) TestContainerSzGetMessages(c *C) {
	// it seems tests are executed in alphabetical order
	resource := "/messages"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token3))
	server1.ServeHTTP(response, request)
	var res controllers.NotificationsPacket
	_ = json.Unmarshal(response.Body.Bytes(), &res)
	assert.Equal(t1, res.Success, true)
	assert.Equal(t1, res.Notifications[0].PayloadCiphertext, payloadciphertext)

	resource = "/messages"
	response = httptest.NewRecorder()
	request, _ = http.NewRequest("DELETE", resource, nil)

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token3))
	server1.ServeHTTP(response, request)
	var deleteRes controllers.NotificationsDeleteResponse
	_ = json.Unmarshal(response.Body.Bytes(), &deleteRes)
	assert.Equal(t1, deleteRes.Success, true)
}

func (s *StorageTestSuite) TestContainerUnshare(c *C) {
	// it seems tests are executed in alphabetical order
	resource := "/peer/" + user2

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var res controllers.PeerMessage
	_ = json.Unmarshal(response.Body.Bytes(), &res)
	assert.Equal(t1, res.Success, true)
	var accountId = res.Peer.AccountId

	resource = "/container/unshare"
	var chunk = controllers.ContainerUnshareChunk{
		ContainerNameHmac: containerNameHmac,
		ToAccountId: accountId,
	}

	response = httptest.NewRecorder()
	jsonStr, _ := json.Marshal(chunk)
	request, _ = http.NewRequest("POST", resource, bytes.NewBuffer(jsonStr) )

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var shareRes controllers.ContainerResponseMessage
	_ = json.Unmarshal(response.Body.Bytes(), &shareRes)

	assert.Equal(t1, shareRes.Success, true)
}

func (s *StorageTestSuite) TestContainerZDelete(c *C) {
	// it seems tests are executed in alphabetical order
	resource := "/container/" + containerNameHmac

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", resource, nil)

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var res controllers.ContainerResponseMessage
	_ = json.Unmarshal(response.Body.Bytes(), &res)

	assert.Equal(t1, res.Success, true)
}



