package api_tests

import (
	"github.com/xlab-si/e2ee-server/routers"
	"github.com/xlab-si/e2ee-server/core/db"
	"github.com/xlab-si/e2ee-server/controllers"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"os"
	"bytes"
	"testing"
)

func Test1(t1 *testing.T) {
	// TestingT should be called only once and it is called
	// already in auth_middleware_test
	TestingT(t1) 
}

type StorageTestSuite struct{}

var _ = Suite(&StorageTestSuite{})
var t1 *testing.T
var token1 string
var token2 string
var user1 string
var user2 string
var server1 *negroni.Negroni
var containerNameHmac string
var ciphertext string
var payloadciphertext string
var accountId1 string
var accountId2 string

func (s *StorageTestSuite) SetUpSuite(c *C) {
	os.Setenv("GO_ENV", "tests")
	db.Init()

	router := routers.InitRoutes(false)
	server1 = negroni.Classic()
	server1.UseHandler(router)

	containerNameHmac = "hmacfoo"
	ciphertext = "ciphertext"
	payloadciphertext = "payloadciphertext"
	user1 = "foo"
	accountId1 = "fooid"
	token1 = "thisistesttoken " + user1 + " " + accountId1

	user2 = "goo"
	accountId2 = "gooid"
	token2 = "thisistesttoken " + user2 + " " + accountId2

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
        	//Username: "foo", // retrieved from token
        	//AccountId: "fooid", // retrieved from token
	}
	resource := "/account"
	response := httptest.NewRecorder()

	jsonStr, _ := json.Marshal(account)

	request, _ := http.NewRequest("POST", resource, bytes.NewBuffer(jsonStr) )
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)

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
        	//Username: "foo", // retrieved from token
        	//AccountId: "fooid", // retrieved from token
	}
	resource = "/account"
	response = httptest.NewRecorder()

	//jsonStr, _ = json.Marshal(account)

	// TODO: tests are to be fixed, now token1 can share files of token2, but
	// this is due to the fact that this are fake tokens and some hecks were
	// needed to enable tests

	//request, _ = http.NewRequest("POST", resource, bytes.NewBuffer(jsonStr) )
	//request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token3))
	//server1.ServeHTTP(response, request)
}

func (s *StorageTestSuite) SetUpTest(c *C) {

}

func (s *StorageTestSuite) TestAccountNotExists(c *C) {
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
		ToAccountId: accountId1,
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

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token1))
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

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token1))
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

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token1))
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

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
	server1.ServeHTTP(response, request)
	var res controllers.NotificationsPacket
	_ = json.Unmarshal(response.Body.Bytes(), &res)
	assert.Equal(t1, res.Success, true)
	assert.Equal(t1, res.Notifications[0].PayloadCiphertext, payloadciphertext)

	resource = "/messages"
	response = httptest.NewRecorder()
	request, _ = http.NewRequest("DELETE", resource, nil)

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token2))
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



