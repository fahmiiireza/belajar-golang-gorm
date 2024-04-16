build:
	docker compose up --build
migrateup:
	migrate -path migrate -database "postgresql://user:password@db:5432/library?sslmode=disable" -verbose up 
migratedown:
	migrate -path migrate -database "postgresql://user:password@db:5432/library?sslmode=disable" -verbose down 