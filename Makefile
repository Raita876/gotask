.PHONY: up
up:
	docker compose up -d
	sleep 10

.PHONY: down
down:
	docker compose down 

.PHONY: mysql
mysql:
	mysql -h 127.0.0.1 -P 3306 -umysql -pmysql
