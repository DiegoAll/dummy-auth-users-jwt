# dummy-auth-users-jwt

Dummy Rest API with Go that implements JWT based authentication.

        go build main.go
        ./main

### Home

        GET http://localhost:5050/home

### Signup

        POST http://localhost:8000/signin
        {
        "email": "oelo@gmail.com",
        "password": "oelo"
        }

### Login

        POST http://localhost:8000/login
        {
        "email": "oelo@gmail.com",
        "password": "oelo"
        }