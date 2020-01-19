package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/moguchev/BD-Forum/app/forum/repository"
	"github.com/moguchev/BD-Forum/app/forum/service"
	"github.com/moguchev/BD-Forum/app/server/delivery"
	"github.com/moguchev/BD-Forum/pkg/config"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewRouter() (*mux.Router, error) {
	router := mux.NewRouter()

	DBConn, err := ConnectToDB(config.DBPath)
	if err != nil {
		return nil, err
	}

	uS := service.Service{
		Repository: &repository.Repository{
			DbConn: DBConn,
		},
	}

	//in case launch from docker

	// err = uS.Repository.InitDBSQL()
	// if err != nil {
	// 	return nil, err
	// }

	h := delivery.Handler{
		Service: uS,
	}

	// user
	router.HandleFunc("/user/{nickname}/create", h.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user/{nickname}/profile", h.GetUserByNick).Methods(http.MethodGet)
	router.HandleFunc("/user/{nickname}/profile", h.UpdateUser).Methods(http.MethodPost)
	// service
	router.HandleFunc("/service/clear", h.Clear).Methods(http.MethodPost)
	router.HandleFunc("/service/status", h.Status).Methods(http.MethodGet)
	// forum
	router.HandleFunc("/forum/create", h.CreateForum).Methods(http.MethodPost)
	router.HandleFunc("/forum/{slug}/details", h.GetForum).Methods(http.MethodGet)
	router.HandleFunc("/forum/{slug}/create", h.CreateThread).Methods(http.MethodPost)
	router.HandleFunc("/forum/{slug}/threads", h.GetThreads).Methods(http.MethodGet)

	return router, nil
}

func RunServer() {
	router, err := NewRouter()
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Failed to create router")
	}

	log.Println("Listening ", config.HostAddress)
	log.Fatal(http.ListenAndServe(config.HostAddress, router))
}

func ConnectToDB(psqURI string) (*sqlx.DB, error) {
	conf, err := pgx.ParseURI(psqURI)
	if err != nil {
		return nil, err
	}

	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     conf,
		MaxConnections: config.MaxConn,
	})

	if err != nil {
		log.Println(err.Error())
		log.Fatal("Failed to create connections pool")
	}

	nativeDB := stdlib.OpenDBFromPool(connPool)

	fmt.Println("Connection to DB was created")
	return sqlx.NewDb(nativeDB, "pgx"), nil
}
