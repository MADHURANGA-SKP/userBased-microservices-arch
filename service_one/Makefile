DB_URL = postgresql://pasan:12345@localhost:5432/sharedDB?sslmode=disable

postgres:
	docker run -d --name sharedDB -p 5432:5432 -e POSTGRES_USER=pasan -e POSTGRES_PASSWORD=12345 postgres:16-alpine

createdb:
	docker exec -it sharedDB createdb --username=pasan --owner=pasan sharedDB

migrate:
	migrate create -ext sql -dir db/migrations -seq init_mg

dropdb:
	docker exec -it sharedDB dropdb --username=pasan sharedDB

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up


migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

redis:
	docker run --name sharedDB -p 6379:6379 -d redis:alpine3.19
