#




# docker build -t mysql_dummy .
# docker image build . -t preloaded_db:latest
# docker build -t new_docker_image_name PATH_to_Dockerfile

# docker tag ubuntu-git mrodara538/ubuntu-git:1.0
# docker push mrodara538/ubuntu-git:1.0



# docker run -d --rm --name mysql --user 1000:1000 -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 mysql:latest (Jose)
# docker run -d -p 33060:3306 --name mysql-db -e MYSQL_ROOT_PASSWORD=secret mysql (Fredy)


# Create volume with Docker
# docker rm -f mysql-db; docker volume create mysql-db-data; docker volume ls
# docker run -d -p 33060:3306 --name mysql-db  -e MYSQL_ROOT_PASSWORD=secret --mount src=mysql-db-data,dst=/var/lib/mysql mysql
# docker exec -it mysql-db mysql -p
# ...
# https://platzi.com/tutoriales/1432-docker/3268-como-crear-un-contenedor-con-docker-mysql-y-persistir-la-informacion/


# Create volume with Docker-compose
# https://stackoverflow.com/questions/39175194/docker-compose-persistent-data-mysql

# docker run -d mysql_dummy 


# mysql -proot  (root corresponds to the password) in this mysql -ppassword
# mysql -h 127.0.0.1 -P 3306 -u test_user  -p -e "SHOW DATABASES;"



docker run -it --link some-mysql:mysql --rm mysql sh -c \
'exec mysql -h "$MYSQL_PORT_3306_TCP_ADDR" -
P"$MYSQL_PORT_3306_TCP_PORT" \
-uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'
mysql> create database store;
mysql> use store;
mysql> create table transactions(ccnum varchar(32), date date, amount
float(7,2),
-> cvv char(4), exp date);