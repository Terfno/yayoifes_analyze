DB_NAME:=yayoifes_analyze
DB_HOST:=172.0.0.1
DB_PORT:=3306
DATASOURCE:=mysql://root:password@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)
MIGRATIONPATH:=migrations

init:
	brew install golang-migrate
	brew install mysql

mysql:
	mysql -u root -h localhost --protocol tcp -p $(DB_NAME)

migrate/init:
	mysql -u root -h $(DB_HOST) --protocol tcp -e "create database \`$(DB_NAME)\`" -p$(DB_PASSWORD)
	mysql -u root -h $(DB_HOST) --protocol tcp -e "create database \`test\`" -p$(DB_PASSWORD)

migrate/new:
	@migrate create -dir migrations -ext sql -seq "$(NAME)"

migrate/up:
	@make migrate_up
	@make migrate_up DB_NAME=test

migrate/down:
	@make migrate_down
	@make migrate_down DB_NAME=test

migrate_up:
	@migrate -path $(MIGRATIONPATH) -database "$(DATASOURCE)$(DB_NAME)" up

migrate_down:
	@migrate -path migrations -database "$(DATASOURCE)$(DB_NAME)" down 1

migrate/drop:
	mysql -u root -h $(DB_HOST) --protocol tcp -e "drop database \`$(DB_NAME)\`" -p$(DB_PASSWORD)
	mysql -u root -h $(DB_HOST) --protocol tcp -e "drop database \`test\`" -p$(DB_PASSWORD)
