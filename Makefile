start-server:
	@echo "Starting server..."
	@cd server && go run main.go

start-client:
	@echo "Starting client..."
	@cd client && go run main.go

generate-private-key:
	@openssl ecparam -genkey -name prime256v1 -noout -out private.ec.key
