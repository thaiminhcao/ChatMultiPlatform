ROOT_DIR=$(shell pwd)
API_DIR=$(ROOT_DIR)/api
SCHEMA_DIR=$(ROOT_DIR)/schema
SERVICE=service
SERVICE_DIR=$(ROOT_DIR)/$(SERVICE)


# user service
USER_DIR=$(SERVICE_DIR)/user
USER_API_DIR=$(USER_DIR)/api
USER_MODEL_DIR=$(USER_DIR)/model

# payment service
PAYMENT_DIR=$(SERVICE_DIR)/payment
PAYMENT_API_DIR=$(PAYMENT_DIR)/api
PAYMENT_MODEL_DIR=$(PAYMENT_DIR)/model

# tour service
TOUR_DIR=$(SERVICE_DIR)/tour
TOUR_API_DIR=$(TOUR_DIR)/api
TOUR_MODEL_DIR=$(TOUR_DIR)/model

# message service
MESSAGE_DIR=$(SERVICE_DIR)/message
MESSAGE_API_DIR=$(MESSAGE_DIR)/api
MESSGE_MODEL_DIR=$(MESSAGE_DIR)/model

#gen api code
gen-user-service:
	goctl api go -api $(API_DIR)/user.api -dir $(USER_API_DIR)

gen-payment-service:
	goctl api go -api $(API_DIR)/payment.api -dir $(PAYMENT_API_DIR)

gen-tour-service:
	goctl api go -api $(API_DIR)/tour.api -dir $(TOUR_API_DIR)

gen-message-service:
	goctl api go -api $(API_DIR)/message.api -dir $(MESSAGE_API_DIR)

# gen models

gen-user-model:
	goctl model mysql ddl -src="$(SCHEMA_DIR)/user.sql" -dir="$(USER_MODEL_DIR)"

gen-tour-model:
	goctl model mysql ddl -src="$(SCHEMA_DIR)/tour.sql" -dir="$(TOUR_MODEL_DIR)"

gen-payment-model:
	goctl model mysql ddl -src="$(SCHEMA_DIR)/payment.sql" -dir="$(PAYMENT_MODEL_DIR)"

gen-message-model:
	goctl model mysql ddl -src="$(SCHEMA_DIR)/message.sql" -dir="$(MESSGE_MODEL_DIR)"

runs:
	go run ./main.go -f ROOT_DIR/etc/server.yaml