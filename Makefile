postgres:
	docker run --name jumiap14 -p 5432:5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_OWNER) -d postgres:14-alpine

create_db:
	docker exec -it jumiap14 createdb --username=$(DB_USER) --owner=$(DB_OWNER) $(DB_NAME)

drop_db:
	docker exec -it jumiap14 dropdb $(DB_NAME)

migrate_up:
	migrate -path db/migration/ -database $(DB_SOURCE) -verbose up

migrate_down:
	migrate -path db/migration/ -database $(DB_SOURCE) -verbose down

sqlc:
	sqlc generate

check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

generate_server:
	swagger generate server -f ./swagger.yaml -t ./api -A  stocksTracker

test:
	go test -v -race -short -cover ./...

mockgen:
	mockgen -package mockdb -destination db/mock/store.go github.com/TamerB/ecommerce-stocks-tracker/db/sqlc Store

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o "bin/ecommerce-stocks-tracker-service" "github.com/TamerB/ecommerce-stocks-tracker/api/cmd/stocks-tracker-server"

run:
	./bin/ecommerce-stocks-tracker

.PHONY: postgres createdb dropdb migrateup migratedown sqlc check_install generate_server test mockgen build run