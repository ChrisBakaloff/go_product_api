package go_product_api

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Init(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s sslmodel=disable", user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {}
