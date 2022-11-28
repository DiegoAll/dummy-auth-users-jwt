## SQL

Postgree


delete from users;

> Remember the difference between DELETE and TRUNCATE. 

DELETE (DML): Contains a where clausule. Logs the row deletions.

TRUNCATE (DDL): Remove all the rows from a table. Not Contains a where clausule. Is faster than DELETE, We cannot roll back the data after using the TRUNCATE command.  
PROBAR delete * from users; (Error clasico)


## Docker

docker run -d --name mysqlrepo -e MYSQL_ROOT_PASSWORD=passwordrepo mysqlrepo

no se usaron puertos ...


docker run -d --rm --name mysqlrepo --user 1000:1000 -e MYSQL_ROOT_PASSWORD=passwordrepo -p 3306:3306 mysqlrepo no levanta


docker run -d --rm --user 1000:1000 -e MYSQL_ROOT_PASSWORD=passwordrepo -p 3306:3306 mysqlrepo