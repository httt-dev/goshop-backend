db-up:
	sql-migrate up -config=dbconfig.yml -env="development"
db-down:
	sql-migrate down -config=dbconfig.yml -env="development"
start-db:
	docker run --restart=always --name dev-timewise -d -p 5432:5432 -e PGDATA=/var/lib/postgresql/data/pgdata -v ${HOME}/postgres:/var/lib/postgresql/data -e POSTGRES_PASSWORD=Abc12345 -e POSTGRES_USER=guru -e POSTGRES_DB=timewise postgres