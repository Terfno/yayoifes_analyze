DB_NAME:=yayoifes_analyze_db
DB_HOST:=127.0.0.1
DB_PORT:=63306
DATASOURCE:=mysql://root:password@tcp($(DB_HOST):$(DB_PORT))/
MIGRATIONPATH:=migrations

init:
	make install
	make d/up
	make migrate/init
	make migrate/up

install:
	brew install golang-migrate
	brew install mysql

d/up:
	docker-compose up -d

d/down:
	docker-compose down

mysql:
	mysql -u root -h $(DB_HOST) -P $(DB_PORT) --protocol tcp -p $(DB_NAME)

migrate/init:
	mysql -u root -h 127.0.0.1 -P $(DB_PORT) --protocol tcp -e "create database \`$(DB_NAME)\`" -p$(DB_PASSWORD)

migrate/new:
	@migrate create -dir migrations -ext sql -seq "$(NAME)"

migrate/up:
	@make migrate_up

migrate/down:
	@make migrate_down

migrate_up:
	@migrate -path $(MIGRATIONPATH) -database "$(DATASOURCE)$(DB_NAME)" up

migrate_down:
	@migrate -path migrations -database "$(DATASOURCE)$(DB_NAME)" down 1

migrate/drop:
	mysql -u root -h $(DB_HOST) --protocol tcp -e "drop database \`$(DB_NAME)\`" -p$(DB_PASSWORD)
