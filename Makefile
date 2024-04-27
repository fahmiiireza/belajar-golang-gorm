build:
	docker compose up --build
migrateup:
	npx sequelize-cli db:migrate
migratedown:
	npx sequelize-cli db:migrate:undo
seedup:
	npx sequelize-cli db:seed:all
seeddown:
	npx sequelize-cli db:seed:undo:all