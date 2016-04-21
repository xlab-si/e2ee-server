package db

import (
        "github.com/jinzhu/gorm"
        _"github.com/jinzhu/gorm/dialects/postgres"
        _"github.com/lib/pq"
	"github.com/spf13/viper"
        "fmt"
	"log"
)

var db *gorm.DB

func Init() {
	viper.SetConfigName("config") 
	viper.AddConfigPath("$GOPATH/src/github.com/xlab-si/e2ee-server/")
 
	conf_err := viper.ReadInConfig()
	if conf_err != nil {
		fmt.Println(conf_err)
	}

	var conf = viper.GetStringMap("database")
	var db_type = conf["type"].(string)
	var conn_str = fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
								conf["ip"], conf["name"], conf["user"], conf["password"])
	//log.Println(conn_str)
		
	var err error
	db, err = gorm.Open(db_type, conn_str)
	if err != nil {
			panic(err)
	}

	// for testing:
	db.DropTable(&Account{})
	db.DropTable(&Container{})
	db.DropTable(&ContainerRecord{})
	db.DropTable(&ContainerSessionKeyShare{})
	db.DropTable(&Message{})

	db.CreateTable(&Account{})
	db.CreateTable(&Container{})
	db.CreateTable(&ContainerRecord{})
	db.CreateTable(&ContainerSessionKeyShare{})
	db.CreateTable(&Message{})
}

func FindAccount(accountId string) Account {
        var account Account
        db.Where("account_id = ?", accountId).Find(&account)
        return account
}

func FindAccountByName(username string) Account {
        var account Account
	log.Println(username)
	log.Println(len(username))
        db.Where("username = ?", username).Find(&account)
	log.Println(account)
        return account
}

func StoreAccount(a Account) {
        db.Save(a)
}

func FindContainer(containerNameHmac string) Container {
        var container Container
        db.Where("container_name_hmac = ?", containerNameHmac).Find(&container)
        return container
}

func CreateContainer(accountId string, containerNameHmac string) uint {
        c := Container {
            AccountId: accountId,
            ContainerNameHmac: containerNameHmac,
            LatestRecordId: 0, // todo
        }
        db.Save(&c)
        return c.ID
}

func GetContainerRecords(containerId uint, accountId string) []ContainerRecord {
        var containerRecords []ContainerRecord
        db.Where("container_id = ?", containerId).Find(&containerRecords)
        share := GetContainerSessionKeyShare(containerId, accountId)
        for i := 0; i < len(containerRecords); i++ {
                containerRecords[i].SessionKeyCiphertext = share.SessionKeyCiphertext
        }
        return containerRecords
}

func CreateContainerRecord(containerId uint, accountId string, payloadCiphertext string) {
        r := ContainerRecord {
            ContainerId: containerId,
            AccountId: accountId,
            PayloadCiphertext: payloadCiphertext,
        }
        db.Save(&r)
}

func CreateContainerSessionKeyShare(containerNameHmac string, sessionKeyCiphertext string, accountId string, toAccountId string) {
        containerId := FindContainer(containerNameHmac).ID
        s := ContainerSessionKeyShare{
            ContainerId: containerId,
            AccountId: accountId,
            ToAccountId: toAccountId,
            SessionKeyCiphertext: sessionKeyCiphertext,
        }
        db.Save(&s)
}

func DeleteContainerSessionKeyShare(containerNameHmac string, accountId string, toAccountId string) {
        containerId := FindContainer(containerNameHmac).ID
        db.Where("container_id = ? and account_id = ? and to_account_id = ?", containerId, accountId, toAccountId).Delete(&ContainerSessionKeyShare{})
}

func GetContainerSessionKeyShare(containerId uint, accountId string) ContainerSessionKeyShare {
        var share ContainerSessionKeyShare
        db.Where("container_id = ? and to_account_id = ?", containerId, accountId).Find(&share)
        return share
}

func DeleteContainer(containerNameHmac string) {
        var container Container
        // these are soft deletes, use Unscoped to delete records permanentely
        db.Where("container_name_hmac = ?", containerNameHmac).Find(&container)
        db.Where("container_id = ?", container.ID).Delete(&ContainerRecord{})
        db.Where("container_id = ?", container.ID).Delete(&ContainerSessionKeyShare{})
        db.Delete(&container)
}

func CreateNotification(fromAccountId string, toAccountId string, headersCiphertext string, payloadCiphertext string) uint {
        m := Message {
            FromAccountId: fromAccountId,
            ToAccountId: toAccountId,
            HeadersCiphertext: headersCiphertext,
            PayloadCiphertext: payloadCiphertext,
        }
        db.Save(&m)
        return m.ID
}

func GetNotifications(accountId string) []Message {
        var messages []Message
        db.Where("to_account_id = ?", accountId).Find(&messages)
        return messages
}

func DBDeleteNotifications(accountId string) {
        db.Where("to_account_id = ?", accountId).Delete(Message{})
}
