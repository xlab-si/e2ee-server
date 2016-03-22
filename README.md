# End-to-End Encryption (E2EE) server

This is a server for storage of encrypted files. It provides REST API for user accounts, data storage, and sharing information between users. The client is available at https://github.com/xlab-si/e2ee-client.

*NOTE*: E2EE server is a work in progress. There are issues to be fixed, for example database scheme and authentication (currently, the tokens are used and are issued by the server itself). 

### 1.) Prerequisites
In order to get E2EE server up and running, you will need the following:
* golang
* redis
* postgresql

#### Golang
Install Golang and setup the GOPATH environment variable, for instance:
```sh
$ mkdir ~/goworkspace
$ export GOPATH=$HOME/goworkspace
```

You can then fetch this repository and all its dependencies by running
```sh
$ go get github.com/xlab-si/e2ee-server
```
#### Redis
Redis is used for storing expired tokens, the password needs to be changed and set (see core/redis/redis_cli.go):
`CONFIG SET requirepass "some_password"`

#### Postgresql
Install PostgreSQL and create database e2ee:
`CREATE DATABASE e2ee;`. Depending on what OS distribution you are running, you might also have to run `ALTER USER postgres PASSWORD 'postgres';`. For details, see db/db.go

### 2.) Configure keys and paths
Every environment (testing, preproduction, production) can be configured via appropriate .json configuration files residing in the settings/ folder. By default, server's public and private key are stored in the keys/ subfolder. You should generate a new RSA keypair for your server and update .json files to point to the locations where certificate and key are stored.
Additionally, you have to update paths pointing to these configuration files in _settings/settings.go_ (see environments hash).

### 3.) Compile, test and run
Navigate to your $GOPATH/src/github.com/xlab-si/e2ee-server directory and run
```sh
$ go install
```
This will put the _e2ee-server_ executable in your $GOPATH/bin. For convenience, you can run
```sh
$ export PATH=$PATH:$GOPATH/bin 
```
and you will be able to run E2EE server from arbitrary location with `./e2ee-server`.

#### Running tests
Before starting E2EE server, it is advisable to run tests to check whether everything is working properly. Navigate to the tests/ directory and from both subdirectories (api_tests/ and unit_tests/) run
```sh
$ go test
```

[e2ee-client]: <https://github.com/xlab-si/e2ee-client>

# NOTICE #

This product includes software developed at "XLAB d.o.o, Slovenia". The development started as part of the "SPECS - Secure Provisioning of Cloud Services based on SLA Management" research project (EC FP7-ICT Grant, agreement 610795) and is continued in "WITDOM - empoWering prIvacy and securiTy in non-trusteD envirOnMents" research project (European Union’s Horizon 2020 research and innovation programme, agreement 64437).

* http://www.specs-project.eu/
* http://witdom.eu/
* http://www.xlab.si/



