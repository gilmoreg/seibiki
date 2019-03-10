setup:
	cp -n ./build/example.env ./build/.env
	@echo Edit ./build/.env and fill in missing values

up:
	docker-compose -f ./build/docker-compose.yml up

up-build:
	docker-compose -f ./build/docker-compose.yml up --build

rebuild:
	docker-compose -f ./build/docker-compose.yml build --no-cache

down:
	docker-compose -f ./build/docker-compose.yml down

client:
	cd web && npm start && cd ..

artifacts:
	GO111MODULE=on CGO_ENABLED=0 GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./build/go-service cmd/*.go

deploy:
	sh ./build/deploy.sh

test:
	go test -v -cover ./...
