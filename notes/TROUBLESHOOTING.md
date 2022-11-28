# Repositorio con MySQL

generar DSN

# DATABASE_URL=root:passwordrepo@tcp(localhost:3306)/dummyusers
# DATABASE_URL=root:passwordrepo@tcp(localhost:33060)/dummyusers?param=value
# DATABASE_URL=postgres://postgres:postgres@localhost:3306/postgres?sslmode=disable
# DSN
# username:password@protocol(address)/dbname?param=value
# root:passwordrepo@tcp(localhost:3306)/dummyusers?param=value

docker run -d --rm --name mysqlrepo --user 1000:1000 -e MYSQL_ROOT_PASSWORD=passwordrepo -p 3306:3306 mysqlrepo no levanta

docker run -d --rm --user 1000:1000 -e MYSQL_ROOT_PASSWORD=passwordrepo -p 3306:3306 mysqlrepo  levanta 

al parecer conecta
Error 1054: Unknown column '$1' in 'field list'

corroborar mysql schema y volver a tirar con el coman do que no levanta es atipico .. solo se le quito el --name

al menos al parecer conecta

homologar .. aunque son iguales las tablas entonces ...

revisar repositorios especificos  y poner logs.

o subirla sin comando de jose que se vea el innodb donde se pueda visualizar lo que se ejecuta ...
o poner logs de las querys


RETAKE


$ docker build -t diegoall1990/ubuntu-lab .

$ docker run mysqlrepo mysql

$ docker run -d --rm --user 1000:1000 -e MYSQL_ROOT_PASSWORD=passwordrepo -p 3306:3306 mysqlrepo

$ docker exec -it mysqlrepo mysql -u root -p     MYSQL_ROOT_PASSWORD=passwordrepo


## connection refused

$ docker start mysqlrepo

```
CO0C02GD0T7MD6M:dummy-auth-users-jwt dposada$ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED       STATUS         PORTS                 NAMES
b55d9c7f077a   b8e3ce14f010   "docker-entrypoint.s…"   2 weeks ago   Up 3 seconds   3306/tcp, 33060/tcp   mysqlrepo
```

dial tcp 127.0.0.1:33060: connect: connection refused

$ docker run -d --rm --user 1000:1000 -e MYSQL_ROOT_PASSWORD=passwordrepo -p 3306:3306 mysqlrepo

```
CO0C02GD0T7MD6M:dummy-auth-users-jwt dposada$ docker ps
CONTAINER ID   IMAGE       COMMAND                  CREATED              STATUS              PORTS                               NAMES
f765319a2798   mysqlrepo   "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:3306->3306/tcp, 33060/tcp   goofy_euclid
```

Error 1054: Unknown column '$1' in 'field list'

AHORA SI HAY CONEXION A LA BASE DE DATOS CON ESTE STRING DE CONEXION

```
DATABASE_URL="root:passwordrepo@tcp(127.0.0.1:3306)/dummyusers"
```


## Unknown column '$1' in 'field list'

Hint: https://sebhastian.com/mysql-error-1054-fix/#:~:text=In%20short%2C%20ERROR%201054%20means,column%20name%20in%20your%20statement

- debido a un nombre de columna incorrecto NO! 

	Se valida con describe|explain \d

	Describe table: \d table_name


### Podria ser un tema de definicion del tipo de dato Postgree es de una forma y MySQL de otra ... Validar y homologar ..

postgres=# \d users
                           Table "public.users"
   Column   |            Type             | Collation | Nullable | Default 
------------+-----------------------------+-----------+----------+---------
 id         | character varying(32)       |           | not null | 
 password   | character varying(255)      |           | not null | 
 email      | character varying(255)      |           | not null | 
 created_at | timestamp without time zone |           | not null | now()
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)


mysql> explain users;
+------------+--------------+------+-----+-------------------+-------------------+
| Field      | Type         | Null | Key | Default           | Extra             |
+------------+--------------+------+-----+-------------------+-------------------+
| id         | varchar(32)  | NO   | PRI | NULL              |                   |
| password   | varchar(255) | NO   |     | NULL              |                   |
| email      | varchar(255) | NO   |     | NULL              |                   |
| created_at | timestamp    | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
+------------+--------------+------+-----+-------------------+-------------------+


```
Conclusion:

VARCHAR is an alias for CHARACTER VARYING, so no difference, see documentation :)

The notations varchar(n) and char(n) are aliases for character varying(n) and character(n), respectively. character without length specifier is equivalent to character(1). If character varying is used without length specifier, the type accepts strings of any size. The latter is a PostgreSQL extension.

The only difference is that CHARACTER VARYING is more human friendly than VARCHAR

https://stackoverflow.com/questions/1199468/what-is-the-difference-between-character-varying-and-varchar-in-postgresql
```
creo el usuario
{2IkpaoG9FqJ0nlkiloDXPG2yqNP diego@diego.com $2a$08$iWa/WBSIPqu9v1x2u1RUwuoRkhBuPB7YbD/6NVwjXf4i8AtrTd69C}
context.Background.WithValue(type *http.contextKey, val <not Stringer>).WithValue(type *http.contextKey, val [::1]:5050).WithCancel.WithCancel.WithValue(type mux.contextKey, val <not Stringer>).WithValue(type mux.contextKey, val <not Stringer>)
no pudo insertar el usuario

lo mas probable es que sea un error de como se envia a la base de datos intentar usando comillas sencillas o dobles en algun dato.

sin embargo buscando bajo este error.  Podria haber un indicio que sea tema de mux o de GO como tal,

(type *http.contextKey, val <not Stringer>).WithValue(type *http.contextKey,


Prueba

Postgres
_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)

2IkuxLMl91eKsSyUKUYlhhlC4vl diego@diego.com $2a$08$e5vvwJ0yQUymBD8E2j4qhOMdK/rRzIrd/8KdxXndW0EdZ5YbBYHvS

INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password


-- Este insert funciona en MySQL pero no en postgres. (Con comillas dobles)
INSERT INTO users (id, email, password) VALUES ("2IkuxLMl91eKsSyUKUYlhhlC4vl", "diego@diego.com", "$2a$08$e5vvwJ0yQUymBD8E2j4qhOMdK/rRzIrd/8KdxXndW0EdZ5YbBYHvS");

-- Este insert funciona en MySQL y en Postgres. (Con comillas sencillas)
INSERT INTO users (id, email, password) VALUES ('3IkuxLMl91eKsSyUKUYlhhlC4vl', 'diego2@diego2.com', '$3a$08$e5vvwJ0yQUymBD8E2j4qhOMdK/rRzIrd/8KdxXndW0EdZ5YbBYHvS');

-- Cambiando a comilla simple funciona en Postgres (Con comillas simples)
INSERT INTO users (id, email, password) VALUES ('2IkuxLMl91eKsSyUKUYlhhlC4vl', 'diego@diego.com', '$2a$08$e5vvwJ0yQUymBD8E2j4qhOMdK/rRzIrd/8KdxXndW0EdZ5YbBYHvS');

-- No lo admite MySQL, si es admitido por Postgres con comillas dobles en los campos.
CREATE TABLE PASSENGERS2("Id" INT PRIMARY KEY NOT NULL,"Name" VARCHAR (100) NOT NULL,"Email" VARCHAR (255) UNIQUE NOT NULL,"Age" INTEGER NOT NULL,"Travel_to" VARCHAR (255) NOT NULL,"Payment" INTEGER,"Travel_date" DATE);

-- Es admitido en MySQL y Postgres sin comillas en los campos
CREATE TABLE PASSENGERS(Id INT PRIMARY KEY NOT NULL,Name VARCHAR (100) NOT NULL,Email VARCHAR (255) UNIQUE NOT NULL,Age INTEGER NOT NULL,Travel_to VARCHAR (255) NOT NULL,Payment INTEGER,Travel_date DATE);

INSERT INTO "passengers" ("Id", "Name", "Email", "Age", "Travel_to", "Payment", "Travel_date") VALUES (1, 'Jack', 'jack12@gmail.com', 20, 'Paris', 79000, '2018-1-1');

```
Conclusión

En la parte de VALUES se mandan valores (un objeto var user *models.User)... afectaría en algo el tipo de comillas? 
Se corroboran en diversos ejemplos y para MySQL es indiferente usar comillas dobles y simples, en cambio hasta donde se vió Postgree requiere comillas simples para los strings en los **VALUES.**

Postgree presenta particularidades para crear tablas e insertar datos. 
INSERT INTO "passengers" ("Id", "Name"...) VALUES VALUES (1, 'Jack'...)

Cable aclarar que el ejemplo inicial es con Posgres y el repositorio inserta bien los datos. Por ende hay que buscar la solución de cara a MySQL.
```


### '$1' en algunos casos aparece en el error el value mas no la referencia al parametro.

ERROR 1054 (42S22): Unknown column 'displayname' in 'field list'  Topico (https://sebhastian.com/mysql-error-1054-fix/#:~:text=In%20short%2C%20ERROR%201054%20means,column%20name%20in%20your%20statement)

Error 1054: Unknown column '$1' in 'field list'




Tendria que ver con la serializacion?

		var user = models.User{
			Email:    request.Email,
			Password: string(hashedPassword), // Security
			//Id:       id.String(),
			Id: string(id.String()),
		}

Se revisa y se verifica que todos son strings ...


pendiente finiquitar.

### Encontrada en youtube se cambia el tipo de dato TIMESTAMP NOW por CURRENT_TIMESTAMP

https://www.youtube.com/watch?v=y2PvbM03TFY (Indio)

POSTGREES

-- Prebuild database

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(32) PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

MYSQL

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(32) PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

**MYSQL CHANGE** 

CREATE DATABASE dummyusers;

USE dummyusers;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(32) PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

```
Conclusion:

No resuelve el problema cambiando el tipo de dato created_at

```

### Orden en que se envian los parametros

Se esta verificando el orden en que se mandan los parametros, se decide retornar al master para probar con db de postgres y funciona con normalidad.

Asi es el orden en la base de datos Postgres.

id        password   email      created_at

Asi esta formado el statement sql, aunque no lleva el mismo orden funciona e inserta el registro normalmente en Postgres.

INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)

user models.User  id    email     password

{2JAJoKGXC7X3WEKezHIX7dzQjoy pablo@pablo.com $2a$08$nMnRZ1.tSqRmzZbraxznoOjny5qmP8cr55yLZQGUfTuvrJt.WH9F.}

```
Conclusión:
Aunque el orden de envio de los parametros esta en desorden, esta correcto ya que hay correspondencia entre el nombre del campo y el values. Para el repositorio especifico de Postgres funciona con normalidad.
```



### Cambiar el indice $1 por $0

repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($0, $2, $3)", user.Id, user.Email, user.Password)

Error 1054: Unknown column '$0' in 'field list'

```
En efecto no conoce el nombre de la columna.
Ademas se prueba con el endpoint de login y aparece un mensaje de error similar

Error 1054: Unknown column '$1' in 'where clause'

Pareciera que con MySQL hay un inconveniente para utilizar $1

```




### Cambiar los $1, $2, ...  por ?

Se suprimen los simbolos $
```
_, err := repo.db.ExecContext(ctx, `INSERT INTO users (id, email, password) VALUES (1, 2, 3)`, user.Id, user.Email, user.Password)
return err
```
sql: expected 0 arguments, got 3

```
func addAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
```

Error 1054: Unknown column '$1' in 'where clause'



Se revisa obsec analyzer si es un tema del statement SQL con las comillas.

//_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
_, err := repo.db.ExecContext(ctx, `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`, user.Id, user.Email, user.Password)


Y el error persiste, pero si recuerdo una informacion hay particularidades entre SELECT e INSERT (ambas son DML)

Error 1054: Unknown column '$1' in 'field list'



## CONCLUSION FINAL!!!

**Para utilizar el repositorio con MySQL**

_, err := repo.db.ExecContext(ctx, `INSERT INTO users (id, email, password) VALUES (?, ?, ?)`, user.Id, user.Email, user.Password)


**Para utilizar el repositorio con Postgres**

_, err := repo.db.ExecContext(ctx, `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`, user.Id, user.Email, user.Password)



## INTERESANTE

Ver tipo de dato de una variable

fmt.Printf("Data = %v, Type = %T", user.Id, user.Id)
fmt.Println("\n")
fmt.Println(reflect.TypeOf(user.Id))

fmt.Println("imprimiendo\nusuario\nmail\ncontraseña")

fmt.Println(r.Context())



## Nuevo Hallazgo 

### revisar statements SQL vs (obsecanalyzer-boo-roz)

Obsec-Analyzer con respecto a INSERT
```
func (r *repository) SaveIPInfo(ctx context.Context, ip models.IP_Data) error {
	query := `INSERT INTO ip_data (ip, abuse_confidence_score, is_whitelisted, totalReports) VALUES(?, ?, ?, ?)`
	stmt, err := r.db.PrepareContext(ctx, query)
... 
_, err = stmt.Exec(ip.IP, ip.Abuse_confidence_score, ip.Is_whitelisted, ip.TotalReports)
```



Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running? 

Brujas o cambio a colima desde AD?

colima stop
colima status
colima start
colima status
colima --help
colima start --help
