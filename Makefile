CONTAINER_NAME=eztool
DB_NAME=eztool
DB_PORT=5434
DB_USER=postgres
DB_PASSWORD=12345
DB_IP=$(shell docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(CONTAINER_NAME))
watch:
	gowatch

# mig-up:
# 	migrate -path db/migrations/ -database "postgres://postgres:12345@localhost:5434/eztool?sslmode=disable" up



## varify result
select:
	docker run -it --rm --network bridge postgres psql -h $(DB_IP)  -U postgres -W $(DB_NAME) -c "select * from users order by nconst limit 5;"

## psql cli
psql:
	docker run -it --rm --network bridge postgres psql -h $(DB_IP) -U $(DB_USER)

## db
db:
	docker run --name $(CONTAINER_NAME) -p $(DB_PORT):5432 -e POSTGRES_PASSWORD=$(DB_PASSWORD) -e POSTGRES_DB=$(DB_NAME) -d postgres:12

rmdb:
	docker stop $(CONTAINER_NAME) 
	docker container rm $(CONTAINER_NAME) 


## migrate
mig-create:
	migrate create -ext sql -dir db/migrations -seq create_users_table

mig-up:
	migrate -path db/migrations/ -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5434/eztool?sslmode=disable" up

mig-down:
	migrate -path db/migrations -database postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5434/eztool?sslmode=disable -verbose down

.PHONY: watch db psql select mig-up rmdb
.SILENT: watch db psql select mig-up rmdb