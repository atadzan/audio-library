# migration
create-migration:
	migrate create -ext sql -dir db/migrations -tz Local $(name)
apply-migration:
	migrate -database 'postgres://helloUser:helloWorld@localhost:6432/helloDB?sslmode=disable' -path db/migrations up
rollback-migration:
	migrate -database 'postgres://helloUser:helloWorld@localhost:6432/helloDB?sslmode=disable' -path db/migrations down $(num)

# Dependencies
dependencies-up:
	docker-compose -f ./deployment/docker-compose.yaml up -d
dependencies-down:
	docker-compose -f ./deployment/docker-compose.yaml down

# application
run:
	go run cmd/main.go

