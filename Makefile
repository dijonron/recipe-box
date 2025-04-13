.PHONY: up
up:
	@echo 'Running Docker Compose for app'
	docker compose -f docker-compose.yml up --build -d
	@echo 'Docker images started!'

.PHONY: down
down:
	@echo 'Stopping docker compose...'
	docker compose -f docker-compose.yml down
	@echo 'Done!'

PHONY: proto
proto:
	@echo 'Generating protobuf...'
	protoc --proto_path=./proto --go_out=./backend/proto/grpc/user --go_opt=paths=source_relative --go-grpc_out=./backend/proto/grpc/user --go-grpc_opt=paths=source_relative ./proto/user.proto
	protoc --proto_path=./proto --go_out=./backend/proto/grpc/auth --go_opt=paths=source_relative --go-grpc_out=./backend/proto/grpc/auth --go-grpc_opt=paths=source_relative ./proto/auth.proto
	protoc --proto_path=./proto --go_out=./backend/proto/grpc/tenant --go_opt=paths=source_relative --go-grpc_out=./backend/proto/grpc/tenant --go-grpc_opt=paths=source_relative ./proto/tenant.proto
	@echo 'Done!'

.PHONY: migrate-create
migrate-create:
	@echo 'Creating migration...'
	migrate create -ext sql -dir backend/migrations -seq ${NAME}
	@echo 'Done!'

.PHONY: migrate-run 
migrate-run:
	@echo "Running migration... "
	migrate -database ${DB_URL} -path backend/migrations up
	@echo 'Done!'

.PHONY: migrate-down
migrate-down:
	@echo 'Reverting migration...'
	migrate -database ${DB_URI} -path ./backend/migrations down
	@echo 'Done!'

.PHONY: fmt
fmt:
	cd ./backend && go fmt ./...