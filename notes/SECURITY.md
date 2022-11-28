# Seguridad




## Almacenar hash de la contraseña


postgres=# select id, password, email from users;
             id              |                           password                           |      email       
-----------------------------+--------------------------------------------------------------+------------------
 2HkkSg63zLdGMyy4rEMyU1sq8yD | mypassword                                                   | diego@gmail.com
 2HkspPjJUE5qSux0Tg0HOVP6k7k | mypassword                                                   | diego2@gmail.com
 2HnXlRDQJ0Lx6RLEazM4ckPQtjl | $2a$08$eOj1sCjzE7Hr0wbw4k4c4uKqqbptE6.TmJ3HD/.lSMd5E7IbQYmsS | diego3@gmail.com
(3 rows)


> Nota: no puede pegarse alineado requiere de foto.



## Fuente particular


Indagar si es buena practica compartir structs para request y response, puntualmente en el caso de signup y login


Caso upsert .... y derivados ...  investigar 


En este escenario, 

r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)





2022/11/20 01:18:24 crypto/bcrypt: hashedSecret too short to be a bcrypted password

una cosa es el loggger del err y otra http Error


// The first parameter that is sent is the hash, what we are storing in the database, and the second parameter is what thclient is sending
if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
	// Debugging invalid credentials
	log.Println(err)
	http.Error(w, "invalid credentials", http.StatusUnauthorized)
	return
}


