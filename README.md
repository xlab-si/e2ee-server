# End-to-End Encryption (E2EE) server

This is a server for storage of encrypted files. It provides REST API for user accounts, data storage, and sharing information between users. The client is available at https://github.com/xlab-si/e2ee-client.

NOTE: E2EE server is a work in progress. There are issues to be fixed, for example database scheme and authentication (currently, the tokens are used and are issued by the server itself). 
## Install

Set up Go path:
`export GOPATH=$HOME/goworkspace`

Install PostgreSQL and create database e2ee (see db/db.go):
`create database e2ee;`

Redis is used for storing expired tokens, the password needs to be changed and set (see core/redis/redis_cli.go):
`CONFIG SET requirepass "some_password"`

Generate new keys and set up the paths in json files in settings folder.

# Run

Compile and install the project:
`go install`

Run:
`./e2ee`

For running the tests, go into folder with tests and execute:
`go test`

# NOTICE #

This product includes software developed at "XLAB d.o.o, Slovenia". The development started as part of the "SPECS - Secure Provisioning of Cloud Services based on SLA Management" research project (EC FP7-ICT Grant, agreement 610795) and is continued in "WITDOM - empoWering prIvacy and securiTy in non-trusteD envirOnMents" research project (European Union’s Horizon 2020 research and innovation programme, agreement 64437).

* http://www.specs-project.eu/
* http://witdom.eu/
* http://www.xlab.si/


