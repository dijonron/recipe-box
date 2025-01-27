.PHONY: run
run:
	@echo 'Running Docker Compose for app'
	docker compose -f docker-compose.yml up --build -d
	@echo 'Docker images started!'


.PHONY: down
down:
	@echo 'Stopping docker compose...'
	docker compose -f docker-compose.yml down
	@echo 'Done!'


.PHONY: protog
protog:
	@echo 'Generate protobuf...'
	protoc --go_out=. --go-grpc_out=. \
	server/proto/ingredient_service/v1/ingredient_service_v1.proto \
	server/proto/recipe_service/v1/recipe_service_v1.proto
	@echo 'Done!'