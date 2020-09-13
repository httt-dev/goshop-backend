db-up:
	sql-migrate up -config=dbconfig.yml -env="development"
db-down:
	sql-migrate down -config=dbconfig.yml -env="development"
start-db:
	docker run --restart=always --name timewise -d -p 5432:5432 -e PGDATA=/