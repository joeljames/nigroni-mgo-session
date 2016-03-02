package main

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	nigronimgosession "github.com/joeljames/nigroni-mgo-session"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")
	dbColl := "widget"

	dbAccessor, err := nigronimgosession.NewDatabaseAccessor(dbURL, dbName, dbColl)
	if err != nil {
		panic(err)
	}

	n := negroni.Classic()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		session := sessions.GetSession(req)
		session.Set("hello", "world")
	})

	// Register the middleware here.
	n.Use(middleware.NewDatabase(dbAccessor).Middleware())

	n.UseHandler(mux)
	n.Run(":3000")
}
