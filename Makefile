postgres:
		# docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine
		docker run --name postgres14 --network trackit -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

createdb:
		docker exec -it postgres14 createdb --username=postgres --owner=postgres trackit

dropdb:
		docker exec -it postgres14 dropdb --username=postgres trackit

psql: # log in to trackit db in psql terminal
		docker exec -it postgres14 psql -U postgres -d trackit

dcup:
		docker-compose up

dcdown:
		docker-compose down

dcdownforce:
		docker-compose down --rmi all -v

createmigration:
		migrate -help
		migrate create -ext sql -dir db/migration -seq [ADDNAME]

migrateup:
		# migrate -path db/migration -database "postgresql://postgres:postgres@postgres:5432/trackit?sslmode=disable" -verbose up
		migrate -path db/migration -database "postgres://zypwgdad:qfaaZy7k6Xd_Y7xtMwSPG7IIyTuqRWl2@raja.db.elephantsql.com/zypwgdad?sslmode=disable&timezone=Africa/lagos" -verbose up

migrateup1:
		# migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/trackit?sslmode=disable" -verbose up 1
		migrate -path db/migration -database "postgres://zypwgdad:qfaaZy7k6Xd_Y7xtMwSPG7IIyTuqRWl2@raja.db.elephantsql.com/zypwgdad?sslmode=disable&timezone=Africa/lagos" -verbose up 1

migratedown:
		# migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/trackit?sslmode=disable" -verbose down
		migrate -path db/migration -database "postgres://zypwgdad:qfaaZy7k6Xd_Y7xtMwSPG7IIyTuqRWl2@raja.db.elephantsql.com/zypwgdad?sslmode=disable&timezone=Africa/lagos" -verbose down

migratedown1:
		# migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/trackit?sslmode=disable" -verbose down 1
		migrate -path db/migration -database "postgres://zypwgdad:qfaaZy7k6Xd_Y7xtMwSPG7IIyTuqRWl2@raja.db.elephantsql.com/zypwgdad?sslmode=disable&timezone=Africa/lagos" -verbose down 1

sqlc:
		sqlc generate

test:
		go test -v -cover ./...

server:
		go run main.go
	
mock:
		mockgen -package mockdb -destination db/mock/store.go trackit/db/sqlc Store

.PHONY: postgres createdb dropdb psql createmigration migrateup migratedown migrateup1 migratedown1 sqlc test server mock dc