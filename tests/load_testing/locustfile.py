from locust import HttpLocust, TaskSet, task
import json
import requests
import pprint

class UserBehavior(TaskSet):
    def on_start(self):
        """ on_start is called when a Locust start before any task is scheduled """
        self.login()

    def login(self):
	self.user1 = "miha"
	self.user2 = "haku"
	j = {"username":self.user1, "password":"testing"}
        response = self.client.post("/token-auth", json=j)
	j1 = json.loads(response.content)
	self.token1 = j1["token"]
	self.containerNameHmac = "1f806eda7c2b249b3153"
	j = {"username":self.user2, "password":"testing"}
        response = self.client.post("/token-auth", json=j)
	j2 = json.loads(response.content)
	self.token2 = j2["token"]

    @task(1)
    def account_exists(self):
	headers = {"Authorization":"Bearer %s" % self.token1}
        response = self.client.get("/accountexists", headers=headers)
	#print "Response content: %s" % response.content

    @task(2)
    def account_create1(self):
	headers = {"Authorization":"Bearer %s" % self.token1}
	account = {
		"ContainerNameHmacKeyCiphertext": "",
                "HmacKeyCiphertext": "",
                "KeypairCiphertext": "",
                "KeypairMac": "",
                "KeypairMacSalt": "",
                "KeypairSalt": "",
                "PubKey": "",
                "SignKeyPrivateCiphertext": "",
                "SignKeyPrivateMac": "",
                "SignKeyPrivateMacSalt": "",
                "SignKeyPub": ""}

        response = self.client.post("/account", headers=headers, json=account)
	#print "Response content: %s" % response.content

    @task(3)
    def account_create1(self):
	headers = {"Authorization":"Bearer %s" % self.token2}
	account = {
		"ContainerNameHmacKeyCiphertext": "",
                "HmacKeyCiphertext": "",
                "KeypairCiphertext": "",
                "KeypairMac": "",
                "KeypairMacSalt": "",
                "KeypairSalt": "",
                "PubKey": "",
                "SignKeyPrivateCiphertext": "",
                "SignKeyPrivateMac": "",
                "SignKeyPrivateMacSalt": "",
                "SignKeyPub": ""}

        response = self.client.post("/account", headers=headers, json=account)
	#print "Response content: %s" % response.content

    @task(4)
    def account_get(self):
	headers = {"Authorization":"Bearer %s" % self.token1}
        response = self.client.get("/account", headers=headers)
	#print "Response content: %s" % response.content

    @task(5)
    def container_create(self):
	headers = {"Authorization":"Bearer %s" % self.token1}
	container_chunk = {
	    "toAccountId": 0,
	    "sessionKeyCiphertext": ""
	}
        response = self.client.put("/container/" + self.containerNameHmac, headers=headers, json=container_chunk)
	#print "Response content: %s" % response.content

    @task(6)
    def container_record_create(self):
	headers = {"Authorization":"Bearer %s" % self.token1}
	container_record_chunk = {
	    "containerNameHmac": self.containerNameHmac,
	    "payloadCiphertext": ""
	}
        response = self.client.post("/container/record", headers=headers, json=container_record_chunk)
	#print "Response content: %s" % response.content

    @task(7)
    def container_get(self):
	headers = {"Authorization":"Bearer %s" % self.token1}
        response = self.client.get("/container/" + self.containerNameHmac, headers=headers)
	#print "Response content: %s" % response.content

    @task(8)
    def container_share(self):
	headers = {"Authorization":"Bearer %s" % self.token1}

        response = self.client.get("/peer/" + self.user2, headers=headers)
	j1 = json.loads(response.content)
	accountId = j1["peer"]["accountId"]

	container_share_chunk = {
	    "containerNameHmac": self.containerNameHmac,
	    "toAccountId": accountId,
	    "sessionKeyCiphertext": ""
	}
        response = self.client.post("/container/share", headers=headers, json=container_share_chunk)

	container_notification_chunk = {
	    "fromUsername": self.user1,
	    "toAccountId": accountId,
	    "headersCiphertext": "",
	    "payloadCiphertext": ""
	}
        #response = self.client.post("/peer", headers=headers, json=container_notification_chunk)
	#print "Response content: %s" % response.content

    #@task(9)
    def container_get_notifications(self):
	headers = {"Authorization":"Bearer %s" % self.token2}

        response = self.client.get("/messages", headers=headers)
        #print "Response content: %s" % response.content

        response = self.client.delete("/messages", headers=headers)
        #print "Response content: %s" % response.content



class WebsiteUser(HttpLocust):
    task_set = UserBehavior
    min_wait=5000
    max_wait=9000
