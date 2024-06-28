CURRENT_DIR=$(shell pwd)

proto-gen:
	./script/gen-proto.sh ${CURRENT_DIR}

transport-create:
	migrate create -ext sql -dir storage/migrateons/transport -seq insert_transport_table

transport-up:
	migrate -database 'postgres://postgres:hamidjon4424@localhost:5432/newbase?sslmode=disable' -path storage/migrateons/transport up

transport-down:
	migrate -database 'postgres://postgres:hamidjon4424@localhost:5432/newbase?sslmode=disable' -path storage/migrateons/transport down

weather-create:
	migrate create -ext sql -dir storage/migrateons/weather -seq insert_weather_table

weather-up:
	migrate -database 'postgres://postgres:hamidjon4424@localhost:5432/newbase?sslmode=disable' -path storage/migrateons/weather up

weather-down:
	migrate -database 'postgres://postgres:hamidjon4424@localhost:5432/newbase?sslmode=disable' -path storage/migrateons/weather down

