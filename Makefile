NAME :=

build:
	@go build -o ./bin/echo_kick ./cmd/app/main.go

run: generate build
	@./bin/echo_kick

generate:
	@templ generate && npm run build

db_init:
	@go run cmd/migrate/main.go db init

db_migrate:
	@go run cmd/migrate/main.go db migrate

db_rollback:
	@go run cmd/migrate/main.go db rollback

db_create_go:
	$(call __require_NAME)
	@go run cmd/migrate/main.go db create_go ${NAME}

db_create_sql:
	$(call __require_NAME)
	@go run cmd/migrate/main.go db create_sql ${NAME}

db_lock:
	@go run cmd/migrate/main.go db lock

db_unlock:
	@go run cmd/migrate/main.go db unlock

db_status:
	@go run cmd/migrate/main.go db status

db_mark_applied:
	@go run cmd/migrate/main.go db mark_applied

define __require_NAME
    @bash -c "if [ '${NAME}' = '' ]; then echo 'NAME is not defined; you must specify NAME like $$ make NAME=create_xxx_table task'; exit 1; fi"
endef
