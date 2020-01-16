package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/moguchev/BD-Forum/app/forum/repository"
	"github.com/moguchev/BD-Forum/app/forum/service"
	"github.com/moguchev/BD-Forum/app/server/delivery"
	"github.com/moguchev/BD-Forum/pkg/config"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/gorilla/mux"
)

func NewRouter() (*mux.Router, error) {
	router := mux.NewRouter()

	DBConn, err := OpenSqlxViaPgxConnPool(config.DBPath)
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

	router.HandleFunc("/user/{nickname}/create", h.CreateUser).Methods(http.MethodPost)

	return router, nil
}

func RunServer() {
	router, err := NewRouter()
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Failed to create router")
	}
	log.Fatal(http.ListenAndServe(config.HostAddress, router))
}

func OpenSqlxViaPgxConnPool(psqURI string) (*sqlx.DB, error) {
	conf, err := pgx.ParseURI(psqURI)
	if err != nil {
		return nil, err
	}

	fmt.Println(conf)

	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     conf,
		MaxConnections: config.MaxConn,
	})

	if err != nil {
		log.Println(err.Error())
		log.Fatal("Failed to create connections pool")
	}

	nativeDB := stdlib.OpenDBFromPool(connPool)

	fmt.Println("OpenSqlxViaPgxConnPool: the connection was created")
	return sqlx.NewDb(nativeDB, "pgx"), nil
}
