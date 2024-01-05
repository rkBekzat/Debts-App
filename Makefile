postgres:
	docker run --name=debts -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

createdb:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

dropdb:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down