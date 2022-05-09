setup:
	go1.17.2 get
	docker pull postgres14:alpine

postgresup:
	docker run -p 5432:5432 -d --name marksman-postgres -e POSTGRES_USER=$U -e POSTGRES_PASSWORD=$P postgres:14-alpine

postgresdown:
	docker stop marksman-postgres
	docker rm marksman-postgres

createdb:
	docker exec -it marksman-postgres createdb --username=$U --owner=$U marksman

dropdb:
	docker exec -it marksman-postgres dropdb --username=$U marksman

migrateup:
	migrate -path db/migrations/ -database "postgresql://$U:$P@localhost:5432/marksman?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations/ -database "postgresql://$U:$P@localhost:5432/marksman?sslmode=disable" -verbose down

migratefix:
	migrate -path db/migrations/ -database "postgresql://$U:$P@localhost:5432/marksman?sslmode=disable" -verbose force $V

sqlc:
	sqlc compile
	sqlc generate

.PHONY:
	postgresup
	postgresdown
	createdb
	dropdb
	migrateup
	migratedown
	sqlc
