package server

import (
	"context"
	"dummy-api-jwt/database"
	"dummy-api-jwt/repository"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

/*
Para que algo sea considerado un servidor, tendra que tener algo llamado config() que retorne una configuracion
como la de arriba
*/
type Server interface {
	Config() *Config
}

/* Comportamiento*/
type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	// Create a broker instance to be able to return it.
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		//hub:    websocket.NewHub(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)

	// repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	repo, err := database.NewMySQLRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)

	log.Println("Starting server on port", b.Config().Port)

	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
