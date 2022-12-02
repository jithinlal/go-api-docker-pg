start:
	docker-compose build --pull
	docker-compose up -d

stop:
	docker-compose down --remove-orphans

restart:
	make stop
	make start

destroy:
	docker-compose down --rmi all -v --remove-orphans

rebuild:
	make destroy
	make start
