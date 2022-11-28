

https://es.slideshare.net/montoya118/broker-37606068 patron broker aporta a la seguridad

- Instalación de dependencias

go get github.com/golang-jwt/jwt
go get github.com/gorilla/mux
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt

corroborar beneficios si se puede utilizar (https://github.com/kelseyhightower/envconfig/)

- Defecto en handlefunc (no tiene sentido)  y server que no se usa  GRAU


- Modelo

- Diseño handler de registro signup el response solo retorna el id 

el login retorna el token.

En teoria la autenticacion el registro y autenticacion se hace con email y password.
Se prueba y funciona comentando GetUserByID().

// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

Deberia validarse si se autentica con email y password vacio. Para corroborar esa logica  if user == nil
return                                                        // Verify

if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
	http.Error(w, "Invalid Credentials", http.StatusUnauthorized) // Security (Same message Invalid Credentials Bruteforce tpassword using know user)
	// http.Error(w, err.Error(), http.StatusUnauthorized)
	fmt.Println(err)
	return
}



