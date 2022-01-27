
docker-build:
		docker build -t basic-webserver .

run-docker-image:
		docker run -d --name basic-webserver -p 8080:8080  basic-webserver