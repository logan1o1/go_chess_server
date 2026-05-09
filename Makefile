include .env

migrate-up:
	goose -dir schema/migrations up
	pg_dump -U $(DB_USERNAME) -d $(DB_NAME) -s > schema/schema.sql

migrate-down:
	goose -dir schema/migrations down
	pg_dump -U $(DB_USERNAME) -d $(DB_NAME) -s > schema/schema.sql
