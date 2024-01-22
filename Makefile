NAME :=

init:
	@go run cmd/migrate/main.go db init

migrate:
	@go run cmd/migrate/main.go db migrate

rollback:
	@go run cmd/migrate/main.go db rollback

create_go:
	$(call __require_NAME)
	@go run cmd/migrate/main.go db create_go ${NAME}

create_sql:
	$(call __require_NAME)
	@go run cmd/migrate/main.go db create_sql ${NAME}

define __require_NAME
    @bash -c "if [ '${NAME}' = '' ]; then echo 'NAME is not defined; you must specify NAME like $$ make NAME=create_xxx_table task'; exit 1; fi"
endef
