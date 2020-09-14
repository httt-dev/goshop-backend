###Cài đặt docker
docker pull postgres
###Thực thi cmd để tạo DB
### Tham khảo : https://towardsdatascience.com/local-development-set-up-of-postgresql-with-docker-c022632f13ea

docker run --restart=always --name dev-timewise -d -p 5432:5432 -e PGDATA=/var/lib/postgresql/data/pgdata -v ${HOME}/postgres:/var/lib/postgresql/data -e POSTGRES_PASSWORD=Abc12345 -e POSTGRES_USER=guru -e POSTGRES_DB=timewise postgres

###Exec sql cmd
docker exec -it dev-timewise psql -U guru -W timewise  

>dev-timewise : container name\
-U guru -> user\
-W dbname

###PgAdmin install
docker run -p 5050:80 -e "PGADMIN_DEFAULT_EMAIL=user@domain.local" -e "PGADMIN_DEFAULT_PASSWORD=Abc12345" --name dev-pgadmin -d dpage/pgadmin4
 
###Accessing the PostgreSQL from the pgAdmin tool
docker inspect dev-timewise -f "{{json .NetworkSettings.Networks }}"

####Access by browser 
localhost:5050
>email : user@domain.local\
>pass : Abc12345 

Connect server :
>Using IP Address get above , ex : 172.17.0.2\
>Host : 172.17.0.2\
>Port : 5432\
>User : guru\
>Pass: Abc12345
>
