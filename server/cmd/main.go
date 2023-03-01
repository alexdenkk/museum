package main

import (
	"akpl/museum/app"
	"akpl/museum/model"
	"akpl/museum/pkg/db"
	"flag"
	"log"
)

var (
	addr   string
	jwt    string
	dbName string
)

func main() {
	parseFlags()

	// connecting to db
	appDB, err := db.Connect(
		dbName,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("db connected")

	model.Migrate(appDB)

	// create app
	mApp := app.New(appDB, []byte(jwt), addr)
	log.Println("app initialized")

	// run
	log.Fatal(mApp.Run())
}

func parseFlags() {
	flag.StringVar(&addr, "address", ":8080", "address and(or) port for app")
	flag.StringVar(&jwt, "jwt", "fuck", "jwt sign key for generating user tokens")
	flag.StringVar(&dbName, "db-name", "db", "database name")
	flag.Parse()
}
