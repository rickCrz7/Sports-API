package main

import (
	"database/sql"
	"net/http"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rickCrz7/Sports-API/games"
	"github.com/rickCrz7/Sports-API/leagues"
	"github.com/rickCrz7/Sports-API/sports"
	"github.com/rickCrz7/Sports-API/teams"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration from config file
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Open a connection to the MySQL database
	db := viper.GetString("db.url")
	conn, err := sql.Open("mysql", db)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Ping the database to check if the connection is successful
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Set up the HTTP server and routes using Gorilla Mux
	r := mux.NewRouter()
	sports.NewHandler(sports.NewDao(conn), r)
	leagues.NewHandler(leagues.NewDao(conn), r)
	teams.NewHandler(teams.NewDao(conn), r)
	games.NewHandler(games.NewDao(conn), r)
	games.NewScoresHandler(games.NewScoresDao(conn), r)

	// Set up the server and start listening for requests
	addr := viper.GetString("app.addr")
	srv := &http.Server{
		Handler: r,
		Addr:    addr,
	}
	log.Infof("Server running on port %s", addr)
	log.Fatal(srv.ListenAndServe())
}
