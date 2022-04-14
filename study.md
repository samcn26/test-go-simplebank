## db migration

- install golang-migrate

```sh
migrate -version
```

- init

```shell
migrate create -ext sql -dir db/migration -seq init_schema
```

- migrate

```shell
# what is ssl
migrate -path db/migration -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

## sqlc

```shell
sqlc generate
```

- specific name

```sql
-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount) -- specific name, if use $2 will be balance
WHERE id = sqlc.arg(id)
RETURNING *;
```

## go

```shell
go mod init github.com/samcn26/test-go-simplebank

# add dependencies
go mod tidy
```

- driver

```shell
#golang lib pq
go get github.com/lib/pq
```

- testify

```shell
go get github.com/stretchr/testify
```

### mock

```shell
mockgen 
```

## postgres

```shell
docker exec -it postgres12 psql -U root -d simple_bank
```