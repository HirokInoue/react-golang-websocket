run: build-frontend build-api start
	
start:
	env DB_HOST=db:28015 docker-compose up

build-frontend:
	cd frontend; yarn build-prod

build-api:
	cd api; go build