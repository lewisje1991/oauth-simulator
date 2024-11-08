start-server:
	@echo "Starting server..."
	@cd server && go run main.go

start-client:
	@echo "Starting client..."
	@cd client && go run main.go