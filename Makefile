run: build-frontend build-api start
	
start:
	docker-compose up

build-frontend:
	cd frontend; yarn build-prod

build-api:
	cd api; go build