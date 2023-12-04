package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rickCrz7/Sports-API/games"
	"github.com/rickCrz7/Sports-API/leagues"
	"github.com/rickCrz7/Sports-API/sports"
	"github.com/rickCrz7/Sports-API/teams"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Connection parameters
	dbUser := "acme"
	dbPass := "ACM3!"
	dbName := "sports"
	dbHost := "localhost"
	dbPort := "3306"

	// Construct the MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open a connection to the MySQL database
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Ping the database to check if the connection is successful
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	sports.NewHandler(sports.NewDao(conn), r)
	leagues.NewHandler(leagues.NewDao(conn), r)
	teams.NewHandler(teams.NewDao(conn), r)
	games.NewHandler(games.NewDao(conn), r)
	games.NewScoresHandler(games.NewScoresDao(conn), r)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
