1 - Create migration ```migrate create -ext=sql -dir=sql/migrations -seq init```

2 - Create UP and DOWN migrations in ./sql/migrations

3 - Up migration ```migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up```

4 - Down migration ```migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down```

5 - Install SQLc ```brew install sqlc```