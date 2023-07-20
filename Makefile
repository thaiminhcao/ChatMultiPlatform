run-without-dockerfile:
	sudo docker run -d -p 88:80 --name app -v $(pwd):/var/www/html php:7.2-apache
down:
	sudo docker-compose down
up:
	sudo docker-compose up -d
stop:
	sudo docker stop $(sudo docker ps -a -q)
build:
	sudo docker build -t app .
run:
	sudo docker run -d --name my-app app

psql:
	sudo docker exec -it chatapp psql -U root