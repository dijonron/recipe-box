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