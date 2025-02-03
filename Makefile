export $(shell cat .env | sed 's/=.*//' | xargs)


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

.PHONY: bufgen
bufgen:
	@echo 'Generate protobuf...'
	cd ./server && buf generate && cd ../
	@echo 'Done!'

.PHONY: migrate-create
migrate-create:
	@echo 'Creating migration...'
	migrate create -ext sql -dir server/internal/migrations -seq $(name) 
	@echo 'Done!'

.PHONY: migrate-run 
migrate-run:
	@echo "Running migration... "
	migrate -database 'postgres://dijon:mustard@localhost:5432/recipe_box?sslmode=disable' -path server/internal/migrations up
	@echo 'Done!'