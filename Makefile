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
	mkdir -p ./.artifacts
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./.artifacts/go-service cmd/*.go

deploy:
	sh ./build/deploy.sh

test:
	go test -v -cover ./...

ci:
	sleep 15
	mongo test --eval 'db.getSiblingDB("jedict");'
	mongo jedict --eval 'db.createUser({user:"reader",pwd:"password",roles:["readWrite"]});'
	mongorestore --host localhost --port 27017 --username reader --password password --authenticationDatabase jedict --drop --gzip --archive=./build/data/jedict.mongodb.archive
	sleep 20

test_deps:
	docker-compose -f ./build/test_deps.yml up

test_deps_down:
	docker-compose -f ./build/test_deps.yml down
