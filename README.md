# End-to-End Encryption (E2EE) server
[![Build Status](https://travis-ci.org/xlab-si/e2ee-server.svg?branch=master)](https://travis-ci.org/xlab-si/e2ee-server)

This is a server for storage of encrypted files. It provides REST API for user accounts, data storage, and sharing information between users. The client is available at https://github.com/xlab-si/e2ee-client.

*NOTE*: E2EE server is a work in progress. There are issues to be fixed, for example database scheme.

## Installation
You can install E2EE server manually, or via Chef.

### Chef installation
See instructions in the _chef_ directory.

### Manual installation
#### 1. Prerequisites
In order to get E2EE server up and running, you will need the following:
* golang
* postgresql

##### Golang
Install Golang and setup the GOPATH environment variable, for instance:
```sh
$ mkdir ~/goworkspace
$ export GOPATH=$HOME/goworkspace
```

You can then fetch this repository and all its dependencies by running
```sh
$ go get github.com/xlab-si/e2ee-server
```
##### Postgresql
Install PostgreSQL and create a new database, for instance :
`CREATE DATABASE e2ee;`. Depending on what OS distribution you are running, you might also have to run `ALTER USER postgres PASSWORD 'postgres';`. For details, see db/db.go

#### 2. Configure E2EE server

All configuration is placed in config.json.

***IMPORTANT*** E2EE server uses HTTPS. For convenience, we included server's initial private key and self-signed certificate in the _config/ssl/_ folder. However, **you should never use these in your setup! Always be sure to change config in order to point to your own certificate and key!**

If you use a self signed certificate, you will need to resolve the privacy error before using E2EE client - go to the E2EE server address in browser and resolve the notification about non trusted certificate).

#### 3. Compile, test and run
Navigate to your $GOPATH/src/github.com/xlab-si/e2ee-server directory and run
```sh
$ go install
```
This will put the _e2ee-server_ executable in your $GOPATH/bin. For convenience, you can run
```sh
$ export PATH=$PATH:$GOPATH/bin 
```
and you will be able to run E2EE server from arbitrary location with `./e2ee-server`.

##### Running tests
Before starting E2EE server, it is advisable to run tests to check whether everything is working properly. Navigate to the tests/ directory and from the api_tests/ directory run
```sh
$ go test
```

## API

E2EE server exposes a REST API which is used by E2EE client. The API has to be accessed over HTTPS. All data is sent and received as JSON. The API comprises the following functions:

### Account

**GET /accountexists**

Response:

```
{ 
	exists: true
}
```

**GET /account**

Response:

```
{ 
	account: 
	{
		containerNameHmacKeyCiphertext:"{"ciphertext":{"iv":"XmKfWFQDau37fd9PHtTd2A==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","kemtag":"7d1KIFQnsf8aFZnrnEerOd4gDqhzWd86ZOa188l2UNM2D3L0BHIQN8jJtvJov7PZNcUYtWmpi9WlfFeLtlZfZnP8sPqJxukKbdJyckZ9YWAd93fPD62KkpePYafV1uOm","ct":"/lcAhqCbNTbvhcx0BFrLsi4sd22Chj8wS7wZz6HRNnpUFZQibhxPiDoXFOCnXUSZXFvY+IOo0BbM1CrmgybGw2E3R38H4yj0qlk8m6uesi7j9MswKWqCtDqPdjarv3hAceZt"},"signature":[-1199505871,-991057351,1364294406,2016053111,1645785580,-2003212914,-806225612,-402617770,1248545696,143717142,-508260121,-602242414,-1663260530,-824357734,898219474,-1702908234,-1432459445,-543935734,-667956687,-170956949,631701317,-365402822,1803303614,-1317342230],"error":null}",
		hmacKeyCiphertext:"{"ciphertext"{"iv":"BhGb42dDYHnXUBrQB1Up5g==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","kemtag":"d4/7JBQIXWOpXSZU1ABw5+0EouiwGvejnvtFlDrFfll0e19RADg4Xi9JZaPWFwT3Cl8Zuz9XUFLFSuVxGzeKUJC0Cx/UdyLG6aItaGclrbv0T0MR+LbE6pE4CmoSSKuS","ct":"3e96QTxJ4xVOzVd2D1q8M0r0vbocPO6a7wdw+USLw3RpFuVpk/rA1YoE7P+u8yH+3cze+caK0ar0CqOgRhn9qLxMnaChKc1Hz1bbLgja4IwW3EokwYv4RHSCNL/T6JbvAefzy4="},"signature":[1599258020,-585701646,-394967423,761738869,-904872919,-313674992,1714515183,-1894379390,-138877781,104554641,1138401032,1825120319,-312336471,1120441251,483998996,-1143116341,1581155343,1317987871,1238654783,-1132548681,-1768190919,-105824315,-238576744,-419672131],"error":null}",
keypairCiphertext:"{"iv":"RUcGtDIsy1FHWX81gMFLA==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","ct":"ma8O7ichq6+TmtQBd3HHxw/IvUL9VuXIOE7lYdsnV6FpBaqBqhKi6pVQ5+ebDEVia7mWwm7efhcW5HOxu4Hfd9uapGad5mYdqfvsyJxbn5w6MoIxuOxEB1oG5egjsHxA7LqPXu/D+GxLv1CpZCu29gB5i2YHwZFhAvKH9SZnveocNxkSIbg/GMgdYIWJPuikhHK6xTR3DVpCjwVz+2mD454c5qTan7Ky"}",
		keypairMac:"f06ea0723e7d547af272525fab5dc9d5196bfc7cfc9d66890b59051f6d5502e1",
		keypairMacSalt:"[439658234,276866814,-741508973,343630873,180564535,-998893911,-420863898,-1393801538]",
		keypairSalt:"[932085101,590755397,-2116677445,1532349300,673959209,-2031130816,1666731880,-331511823]",
		pubKey:"{"type":"elGamal","secretKey":false,"point":"JbfTmXTHHKVUbxP06W9D9CzEyYsuHi2hxwCbh2IeM27Ddl0VQSAowxmuoROmoeycdzYai5truBy+WNdWevHaG0f0iI4Ju3cz08sd/8whCRiakBgp2ZqZV0t53pZUrVPK","curve":"c384"}","signKeyPrivateCiphertext":"{"iv":"kLrgCG38DW0XZISyMj80mg==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","ct":"9m3E8h+7N5SuYiUI5jI5tgwXMviuuVO0ZL51e5HrXTo25AxygVg+SF8M96qsXTEgc7qVRnn5H4z4ERAN1b/YAbXNkJO6yPjX6mYV04l/NjYsfr4GLk9YVri39TQdySgEl/19fTiFqCRhmoZllWSElqRSwwqR/vMcj/f744GIaEECyi1qHijQoxLI96BQAlIfFPCBJAT7jDcl6//KUt/mRtmEVgr8+Q=="}",
		signKeyPrivateMac:"e2d2be26d3f1704d793640b2f013d540cff3a622aaab3009a9308266ac414e4b",
		signKeyPrivateMacSalt:"[-214028465,-522674977,994236422,-340920408,295090173,1281192400,-1033081743,948008516]",
		signKeyPub:"{"type":"ecdsa","secretKey":false,"point":"p+bOXeuGrQznyLGB/K2TJsKND82dP62yxCtJpaujv/L9cZkqrbtdyDQV3j8cr5NCR6JALUhiVL98bixGtWiR/Ky2o8tbcLN8qv4xJ0RvF+3iOL4QaJWp5t7YQDeqbDqR","curve":"c384"}",
		username:"miha",
		accountId:2
	},
	success: true
}

```

**POST /account**

Response:

```
{ 
	success: true,
	error: ""
}
```


### Containers

**PUT /container/{containerNameHmac}**

Request (it creates an empty container which is later to be filled with container records; toAccountId value is the accout ID of the creator, sessionKeyCiphertext contains the key with which the container records are encrypted):

```
{ 
	toAccountId: 2,
	sessionKeyCiphertext: 
	{"ciphertext":{"iv":"w9LQ/R+Av5wzKx29HZ+M3w==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","kemtag":"cJaqO8nSqNvtT0q/2nuZC858dGq9USoTfgc4n8eC4nXIdLeLkEaq7oIPW+tWG9uTyYqhX3kyYHw+skqb1RXKS0iJUnTmNR6qbXBLWt/lD2ifxPVGwJKFN+BBhTGFz4tC","ct":"bVqCrb4aaJ89HYiqbOI94HRPDppilpsaSBPXhegmeFIWw1nE+s/X/XAf2OE1a0z8XLJQKi79YLpM/rxavv8oBJ6KXj4wt0/nzeGif+zl2DzP1ffKktTv32wZ085KSDc/
"},"signature":[-506773525,1540725851,1613865196,517852257,1327901141,-52110893,-930585250,22323590,-1545085722,2094984420,1654824431,-1313354640,788905832,-2107172083,-1910426198,768768161,-1868473650,-305781219,-165321269,-1585626635,1483808224,-2068299750,-1104578426,-426273143]}
}
```

Response:

```
{ 
	success: true,
	error: ""
}
```

**POST /container/record**

Request:

```
{ 
	containerNameHmac: "1f806eda7c2b249b315853fe3c117a919d1bab4a9a776cce6b02beaa7987671a",
	payloadCiphertext: "{"ciphertext":"{"iv":"3JWSMwo4PfmR7oPJLUfKnQ==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","ciphertextBuffer":"tgUh3FFvmIoVvgd8mOf63rriBjhw6jSZPg16S3zo2zFqLhhupUce1cU8Zzn2TY2lRlqzw3nowxRAQlhg/eSpwV5oJDToDv8CqHBZDKw3KYDpTY1fNOcDl12C8VNb5NwD5VlM7yUz8pKAc6aNCKu1qYCRnEFc70setYrgsUXBxxqhPdHiykgwHiUv4LAv7oa4BXDjE6YAq6XPnh/Tr08fYsWX6TtvzWrckW3LcsxYmGwZKUDSAtILCLgE/SJX3h/NXjm50u160xyBlTr2gV6aKChGDEwgvfm2nEdTDK1LfGJhbF4TTGr4qolYqjLl4WG6czTtyFztJga0K7WKJnflbtBcEu++5gaLnRUUc/hnm2AoDXnLf0lis8DVHRcS5pNwy0Ur6wPaPiJaPPEr8gT9pzKLmkFf6MJgDeEbb990bZbOhgPb5RixQcpdl+O1mECtdS+T5hIAB/9u42vvI5uvbtcyJha+i7QZ4LRvI3U6eg/oDsn5S7joMFuyy20tgXNC7Z5qvIZjhBk3DV7R+P0BkxawYvRIfhgypLAJVZDeg0o4I6Fk/G4l0Uob9zfTAX1ktjWNCduxt5xt6I2TvJ4=","ciphertextTag":"QIlcm5SNozU="}","signature":[915405229,344965878,-1549702246,1687141383,-1432993905,-325126114,-1665284064,-2016789398,-280544420,-27187270,-365294673,-718813603,650634367,914218426,528827448,-384805708,86829866,493938309,1550644324,1845554639,-1367210999,-1487352858,-961216500,283126276]}"
}
```

Response:

```
{ 
	success: true,
	error: ""
}
```

**POST /container/{containerNameHmac}**

Response:

```
{ 
	success: true,
	error: "",
	records:
	[{"ID":10,"CreatedAt":"2016-03-29T07:42:41.663631Z","UpdatedAt":"2016-03-29T07:42:41.663631Z","DeletedAt":null,"containerId":7,"accountId":2,"payloadCiphertext":"{"ciphertext":"{"iv":"3JWSMwo4PfmR7oPJLUfKnQ==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","ciphertextBuffer":"tgUh3FFvmIoVvgd8mOf63rriBjhw6jSZPg16S3zo2zFqLhhupUce1cU8Zzn2TY2lRlqzw3nowxRAQlhg/eSpwV5oJDToDv8CqHBZDKw3KYDpTY1fNOcDl12C8VNb5NwD5VlM7yUz8pKAc6aNCKu1qYCRnEFc70setYrgsUXBxxqhPdHiykgwHiUv4LAv7oa4BXDjE6YAq6XPnh/Tr08fYsWX6TtvzWrckW3LcsxYmGwZKUDSAtILCLgE/SJX3h/NXjm50u160xyBlTr2gV6aKChGDEwgvfm2nEdTDK1LfGJhbF4TTGr4qolYqjLl4WG6czTtyFztJga0K7WKJnflbtBcEu++5gaLnRUUc/hnm2AoDXnLf0lis8DVHRcS5pNwy0Ur6wPaPiJaPPEr8gT9pzKLmkFf6MJgDeEbb990bZbOhgPb5RixQcpdl+O1mECtdS+T5hIAB/9u42vvI5uvbtcyJha+i7QZ4LRvI3U6eg/oDsn5S7joMFuyy20tgXNC7Z5qvIZjhBk3DV7R+P0BkxawYvRIfhgypLAJVZDeg0o4I6Fk/G4l0Uob9zfTAX1ktjWNCduxt5xt6I2TvJ4=","ciphertextTag":"QIlcm5SNozU="}","signature":[915405229,344965878,-1549702246,1687141383,-1432993905,-325126114,-1665284064,-2016789398,-280544420,-27187270,-365294673,-718813603,650634367,914218426,528827448,-384805708,86829866,493938309,1550644324,1845554639,-1367210999,-1487352858,-961216500,283126276]}","sessionKeyCiphertext":"{"ciphertext":{"iv":"rTd3d13/8oqf9wZ0IHva4A==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","kemtag":"DAlMVughvGbrfiXKw/p9Cldje0gfHsxD5/k083b+J9w27sslRUuC+tcAqpRaTw4QX8StpP7XB4Zz+ZXsQKc4WCG/Efzv0NFfmVvQ3bUrojkH7y0BHnEmOlMU28HprjLN","ct":"dMf98DZsHhKRfmIcZ6h+kcGPBM4OQ6k47rTbo9FfhzBtBIui4ggMRWAlxfNJyzG9p8US1X3+BSjE9FbVSfsgvThb7wFRdqIXdxucmkiVd+wNvLiqxuBcBb+EES2qQWI="},"signature":[150564118,1446265489,2084221956,-1667603380,-998974136,-1609862114,1241937389,-1148565839,-460838425,592597932,1845002474,1263007687,1643457844,2066116816,-1901716334,15110227,-186550058,722248554,-727154892,-904125964,1654994764,1488491239,-84698901,-770011829]}"}]
}
```

**POST /container/share**

Request (toAccountId value is the accout ID of the user to whom the container is shared, sessionKeyCiphertext contains the key with which the container records are encrypted – the records are not encrypted for each user, only the key with which the records are encrypted is encrypted with a public key of the user, meaning that sharing operation adds only a (encrypted) key to the database):

```
{ 
	toAccountId: 3,
	containerNameHmac: "1f806eda7c2b249b315853fe3c117a919d1bab4a9a776cce6b02beaa7987671a",
sessionKeyCiphertext: "{"ciphertext":{"iv":"lL7Vt3hsmLDg+nu6pkv7Kg==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","kemtag":"6Mwg9dDq4+EFifvvf2BgC3BQGMlYqCkwsSH9N89DZkZI3O3k/eIbaEjSPADYoMqR6uKaoGWaPxMT56igFaZqgaZKT9AoPAZm95O3sJ57zAZtmfF5iMtOLJPgqRPAkPGB","ct":"LIP2NRfV/azC6Z4zUhlWoOUryyhB77edvVMyMivNFguIXPhsxR6FdoRMBUHiLFbpBGPv+P4k2r2+5Wdjr0+bNll6IH6sCamvQgqwL5VZSRwirOPmjPG6DlHn5cru594="},"signature":[-1214549106,254013776,-1610257010,1615925727,491321452,-709938059,-777100020,2142751685,1631452866,-269601299,126078124,-2028350235,-1511946586,872561864,1061797659,1639542073,-1264617108,1859989011,-1347439519,470056394,1836540034,-2007599379,-1758720317,-1920919531]}"
 }
```

Response:

```
{ 
	success: true,
	error: ""
}
```

**POST /container/unshare**

Request:

```
{ 
	containerNameHmac: "1f806eda7c2b249b315853fe3c117a919d1bab4a9a776cce6b02beaa7987671a",
	toAccountId: 3 
}
```

Response:

```
{ 
	success: true,
	error: ""
}
```

**DELETE /container/{containerNameHmac}**

Response:

```
{ 
	success: true,
	error: ""
}
```

### PEERs

**GET /peer/{username}**

Response:

```
{
	success:true,
	peer:
	{
		accountId:1,
		username:"haku",
		pubKey:"{"type":"elGamal","secretKey":false,"point":"iHtRG/uqh1QvRxoosviBLRL/4ohRb5dM9QwoidGjhDi0XTgLIQfI+h4sKKsxaMDncV6oFNcDiY7kcKSQDW1tI1INe3bDWSgH8kw7ROFUebBIhwvZY0Eu90+t1HMm2naH","curve":"c384"}",
		signKeyPub:"{"type":"ecdsa","secretKey":false,"point":"AaVm72T8oxj0JY2ifK6A6+n3BfmBpQqELJYU3oqrCfJ7fu05TILRjRBAUHP7RSx3cd9Jh7zqbeOFxXVrphp9V2K6Ih/lnAKN6w0uCcWq28TzzcKzXXJkAdFKT9zqiNiT","curve":"c384"}"
	}
}
```

**POST /peer**

Request (send an encrypted message to peer – for example a notification that the file has been shared):

```
{ 
	fromUsername: "miha",
	headersCiphertext: "{"ciphertext":{"iv":"dHtLw5HXNjmtQmEs8g28Pg==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","kemtag":"wQlWaOOdNqMPm4ViKlDI6w+nKuYuPQ6bYLGV7TmNQINwLUkjXezQxnfm53hrmGNQY7w7ujmmhXhsBzT+RKlOYkMcD6vpU4raTsBGZpOtYUlIJJ40mD+Dlai9rc07HXLL","ct":"NzZFLmXg6bvLww=="},"signature":[-1581752453,-1123926707,-523085310,-1388833493,-15832945,-1839818309,2054263925,-1194390647,-138020499,874127401,251032393,1115022351,-246136575,-1648873893,2141973775,-1122199345,782805691,-849692190,-240335402,-1904324519,1279549937,-920291541,-1014084421,-293706101],"error":null}",
	payloadCiphertext: "{"ciphertext":{"iv":"Cids9RzIZqmccB4+WWDGbA==","v":1,"iter":1000,"ks":128,"ts":64,"mode":"ccm","adata":"","cipher":"aes","kemtag":"Km8o3Nz4LdW+RAon1Dg9AhVCQmI6pdO8w9KqyLGebE8wHXLuZtijT4LQn/QlT8sb520jokWrtBfNPgd2IKYXqmYrZq7ZQkB1goQv7VvpkRWaYXz7BvYMfOj8M5mjy7Gh","ct":"K6C5TfjD++oIJ3j71kDPQ3ybePAF5UMaUHoK+CBuFuoZPK/nk2c/O8gRJXne384W36HfMqAlZGRxUkYgS/9JBQeOWhi37AYC6xDax0LddVcTAvVFo5AEzrrqrokYvZavQF1FLmg/bYtAvAg1iSw9KHyvK/F+FpeUcgntpR2uMWz79lqAO9gucjM2RT0/fd87XvyYuG0XLJtYRS8mXJCc/2YJkjy0lwOMnxpVQ5VR"},"signature":[1802005427,244184609,-1578169920,-191539282,1763156810,1744368745,-2089881600,-403680331,-1260472379,498301772,1094732038,-1486869832,1408127850,1305469289,527521014,284064914,97992341,1568715445,159529717,427965579,1462325612,-286693103,-1286074589,1587341698],"error":null}",
	toAccountId: 1
}
```

### Messages

**GET /messages**

Returns all messages sent from other peers.

**DELETE /messages**

Deletes all messages sent from other peers.

## NOTICE #

This product includes software developed at "XLAB d.o.o, Slovenia". The development started as part of the "SPECS - Secure Provisioning of Cloud Services based on SLA Management" research project (EC FP7-ICT Grant, agreement 610795) and is continued in "WITDOM - empoWering prIvacy and securiTy in non-trusteD envirOnMents" research project (European Union’s Horizon 2020 research and innovation programme, agreement 64437).

* http://www.specs-project.eu/
* http://witdom.eu/
* http://www.xlab.si/
