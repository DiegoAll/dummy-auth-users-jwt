FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=passwordrepo
# ENV MYSQL_ROOT_PASSWORD=root

# A diferencia de la de postgres preconstruida le falta esta variable para arrancar sino el sgte error:
# Can't connect to local MySQL server through socket '/var/run/mysqld/mysqld.sock' (2)

COPY up.sql /docker-entrypoint-initdb.d/

# CMD ["mysql"]


# MYSQL_DATABASE
# MYSQL_ROOT_PASSWORD


# https://www.baeldung.com/ops/docker-mysql-container


#/docker-entrypoint-initdb.d/1.sql    # Preloaded database
# CMD ["postgres"]
# CMD ["/usr/bin/falco"]

# remove password 


