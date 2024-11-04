export VERSION ?= $(shell cat ./VERSION)


.PHONY: dockerbuild
dockerbuild:
	docker build \
		-t gotask:$(VERSION) \
		--build-arg VERSION=$(VERSION) \
		.

.PHONY: swag
swag:
	swag init -g cmd/gotask/main.go

.PHONY: up
up: swag
	docker compose build --no-cache
	docker compose up -d

.PHONY: down
down:
	docker compose down 

.PHONY: mysql
mysql:
	mysql -h 127.0.0.1 -P 3306 -umysql -pmysql
