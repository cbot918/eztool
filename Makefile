CONTAINER_NAME=eztool
DB_NAME=eztool
DB_PORT=5434
DB_USER=postgres
DB_PASSWORD=12345

watch:
	gowatch

mig-up:
	migrate -path db/migrations/ -database "postgres://postgres:12345@localhost:5434/eztool?sslmode=disable" up

## varify result
select:
	docker run -it --rm --network bridge postgres psql -h $(shell docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(CONTAINER_NAME)) -U postgres -W $(DB_NAME) -c "select * from names order by nconst limit 5;"

## psql cli
psql:
	docker run -it --rm --network bridge postgres psql -h $(shell docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(CONTAINER_NAME)) -U $(DB_USER)

db:
	docker run --name $(CONTAINER_NAME) -p $(DB_PORT):5432 -e POSTGRES_PASSWORD=$(DB_PASSWORD) -e POSTGRES_DB=$(DB_NAME) -d postgres:12

rmdb:
	docker stop $(CONTAINER_NAME) 
	docker container rm $(CONTAINER_NAME) 

.PHONY: watch db psql select mig-up rmdb
.SILENT: watch db psql select mig-up rmdb