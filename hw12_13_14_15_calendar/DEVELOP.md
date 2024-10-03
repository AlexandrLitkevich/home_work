## Run app

`go run cmd/calendar/main.go --config=configs/config.yaml`   
'go run . --config=PATH'


## Run postgress

docker compose up

## Run postgress
-   docker compose up




## migrations

-  `goose -dir ./migrations postgres "user=alex password=alex host=localhost dbname=calendardb sslmode=disable" up`
-  `goose -dir ./migrations postgres "user=alex password=alex host=localhost dbname=calendardb sslmode=disable" status`