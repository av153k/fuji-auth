
APP_NAME = fuji-auth
BUILD_DIR = $(PWD)./build


build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) cmd/main.go
run: build
	docker compose up
run-silent: build
	docker compose up -d
# test:
# 	make run-silent && sleep 1 && \
# 	go test -v tests/api_integration_test.go
clean:
	docker compose down && \
	docker rmi fuji-auth:latest

docker.run: docker.network swag docker.postgres docker.fuji_auth 

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fuji_auth.build:
	docker build -t fuji_auth .

docker.fuji_auth: docker.fuji_auth.build
	docker run --rm -d \
		--name fuji-auth-ms \
		--network dev-network \
		-p 5000:5000 \
		fuji_auth

docker.postgres:
	docker run --rm -d \
		--name fuji-app-postgres \
		--network dev-network \
		-e POSTGRES_USER=mtfuji \
		-e POSTGRES_PASSWORD=generated1321 \
		-e POSTGRES_DB=fuji_auth \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

swag:
	swag init  --output docs/ --dir cmd,pkg/handlers,pkg/models,pkg/routes