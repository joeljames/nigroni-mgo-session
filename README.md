nigroni-mgo-session
===================

`nigroni-mgo-session` is a simple Negroni middleware/handler for easy mgo session management in requests.

Ussage
------
```
package main

import (
    "net/http"

    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
)

func main() {
    dbURL := os.Getenv("DATABASE_URL")
    dbName := os.Getenv("DATABASE_NAME")
    dbColl := "widget"

    dbAccessor, err := utils.NewDatabaseAccessor(dbURL, dbName, dbColl)
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
    n.Use(middleware.NewDatabase(databaseAccessor).Middleware())

    n.UseHandler(mux)
    n.Run(":3000")
}

```
