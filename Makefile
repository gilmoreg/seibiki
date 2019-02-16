setup:
	cp ./build/example.env ./build/.env
	echo Edit ./build/.env and fill in missing values

up:
	docker-compose -f ./build/docker-compose.yml up

up-build:
	docker-compose -f ./build/docker-compose.yml up --build

down:
	docker-compose -f ./build/docker-compose.yml down

client:
	cd web && npm start

deploy:
	sh ./build/deploy.sh