ENVFILE = .env
ENVFILE_EXAMPLE = .env.sample

# Create the .env file if not exit.
ifeq ("$(wildcard $(ENVFILE))","")
		$(shell cp $(ENVFILE_EXAMPLE) $(ENVFILE))
endif
# Load the environment variables from .env file
include $(ENVFILE)

create_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

compose_up:
	docker compose up -d

compose_down:
	docker compose down

compose_test_up:
	@docker compose -f compose-test.yaml up -d

compose_test_down:
	@docker compose -f compose-test.yaml down

# make sure the environment variables are present in the system
# by default I include it from .env file
migrate_up:
	@migrate -path db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5432/${DB_DATABASE}?sslmode=disable" up

migrate_down:	
	@migrate -path db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5432/${DB_DATABASE}?sslmode=disable" down

migrate_test_up:
	@migrate -path db/migrations -database "postgres://${TEST_DB_USERNAME}:${TEST_DB_PASSWORD}@localhost:${TEST_DB_PORT}/${TEST_DB_DATABASE}?sslmode=disable" up

migrate_test_down:	
	@migrate -path db/migrations -database "postgres://${TEST_DB_USERNAME}:${TEST_DB_PASSWORD}@localhost:${TEST_DB_PORT}/${TEST_DB_DATABASE}?sslmode=disable" down
	

sqlc:
	sqlc generate

test: 
	@echo "Creating db containers"
	@make -s compose_test_up 
	@echo "Waiting for db be ready"
	@sleep 2
	echo "Running migrations"
	@make -s migrate_test_up
	
	@echo "Starting tests"
	@go test ./... -v -cover
	@echo "Removing db containers"
	@make -s compose_test_down

dev:
	air