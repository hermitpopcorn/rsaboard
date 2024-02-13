package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hermitpopcorn/rsaboard/internal/database"
)

var db database.Database

func init() {
	sqliteDatabase, err := database.OpenSQLiteDatabase("sqlite.db")
	if err != nil {
		panic(err)
	}

	db = sqliteDatabase
}

type configuration struct {
	port   int
	origin string
}

var config configuration

func init() {
	port := flag.Int("port", 8080, "Port number")
	origin := flag.String("origin", "*", "Origin host")
	flag.Parse()

	config = configuration{
		port:   *port,
		origin: *origin,
	}
}

func main() {
	startScheduler()

	router := gin.Default()
	injectRoutes(router, db)

	addr := fmt.Sprintf(":%d", config.port)
	router.Run(addr)
}
