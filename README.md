# SQLc

1 - Create migration `migrate create -ext=sql -dir=sql/migrations -seq init`

2 - Create UP and DOWN migrations in ./sql/migrations

3 - Up migration `migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up`

4 - Down migration `migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down`

5 - Install SQLc `brew install sqlc`

6 - Edit sqlc.yaml

```version: "2"
sql:
  - schema: "sql/migrations"
    queries: "sql/queries"
    engine: "mysql"
    gen:
      go:
        package: "db"
        out: "internal/db"
```

7 - Generate code ```sqlc generate````

8 - For float64 data include ```          - db_type: "decimal"
            go_type: "float64"``` to sqlc.yaml or else decimal will be received as string

9 - For transactions create a callTx function that receives an anonymous function that calls all transactions, and rollbacks on error or commits on pass.

