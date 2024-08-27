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

# make sure the environment variables are present in the system
# by default I include it from .env file
migrate_up:
	@migrate -path db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5432/${DB_DATABASE}?sslmode=disable" up

migrate_down:	
	@migrate -path db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5432/${DB_DATABASE}?sslmode=disable" down

sqlc:
	sqlc generate