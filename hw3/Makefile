pg_start:
	docker-compose start
	
pg_stop:
	docker-compose stop

createdb:
	docker-compose exec db createdb --username=postgres --owner=postgres bankstoredb

dropdb:
	docker-compose exec db dropdb -U postgres bankstoredb

migrateup:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable" -verbose down

connect:
	docker-compose exec db psql -U postgres -d postgres

migrateup1:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

.PHONY: pg_start pg_stop createdb dropdb migrateup migratedown migrateup1 migratedown1 connect sqlc 