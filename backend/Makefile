ROOT_DIR=$(shell pwd)
API_DIR=$(ROOT_DIR)/api
SCHEMA_DIR=$(ROOT_DIR)/db/schema
MIGRATIONS_DIR=$(ROOT_DIR)/db/migrations
SRC_STRING="mysql://sql6633512:MPdcLJKEmX@tcp(sql6.freesqldatabase.com:3306)/sql6633512"

SERVICE=service
SERVICE_DIR=$(ROOT_DIR)/$(SERVICE)


# user service
USER_DIR=$(SERVICE_DIR)/user
USER_API_DIR=$(USER_DIR)/api
USER_MODEL_DIR=$(USER_DIR)/model

# message service
MESSAGE_DIR=$(SERVICE_DIR)/message
MESSAGE_API_DIR=$(MESSAGE_DIR)/api
MESSGE_MODEL_DIR=$(MESSAGE_DIR)/model

#gen api code
gen-user-service:
	goctl api go -api $(API_DIR)/user.api -dir $(USER_API_DIR)

gen-message-service:
	goctl api go -api $(API_DIR)/message.api -dir $(MESSAGE_API_DIR)

# gen models

gen-user-model:
	goctl model mysql ddl -src="$(SCHEMA_DIR)/users.sql" -dir="$(USER_MODEL_DIR)"
# goctl model pg datasource --url=$(SRC_STRING) --dir="$(USER_MODEL_DIR)"  --table=users

gen-message-model:
	goctl model mysql ddl -src="$(SCHEMA_DIR)/message_information.sql" -dir="$(MESSGE_MODEL_DIR)"
# goctl model pg datasource -url=$(SRC_STRING) --dir="$(MESSGE_MODEL_DIR)" --table=message_information

runs:
	go run ./main.go -f ROOT_DIR/etc/server.yaml

logs-container:
	sudo docker logs chatapp

migrate-common_schemadb:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq common_schema

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(SRC_STRING) -verbose up 1

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database $(SRC_STRING) -verbose down 1

drop-table:
	migrate -path $(MIGRATIONS_DIR) -database $(SRC_STRING) -verbose force 1