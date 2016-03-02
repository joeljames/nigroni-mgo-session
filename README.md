Work in progress. Please donot use the package yet.
---------------------------------------------------

nigroni-mgo-session
===================

`nigroni-mgo-session` is a simple Negroni middleware/handler for easy mgo session management in requests. It handles creating a new mongo session in the begning of a request and handles closing the mongo session at the end of request.


Getting Started
---------------
1. Before getting started make sure we have a copy of MongoDB running and is accessible. Running it locally would be recommended, but if thats not an option you can run on a remote server. Makes sure the mongo server is accessible.
If you haven't installed `nigroni-mgo-session`, you can run the below command to install it.

    ```
    go get github.com/joeljames/nigroni-mgo-session
    ```

2. Now, that you have a running instance of MongoDB you can set your `DATABASE_URL` and `DATABASE_NAME` environment variable.

    ```
    export DATABASE_URL=mongodb://username:password@localhost:27017/db
    export DATABASE_NAME=db
    ```

3. Now, lets install the framework, go mongo adapter and dependencies.

    * To install [Nigroni](https://github.com/codegangsta/negroni) run the command below.

        ```
        go get github.com/codegangsta/negroni
        ```
    * To install [Gorilla Context](https://github.com/gorilla/context) run the command below.

        ```
        go get github.com/gorilla/context
        ```
    * To install [mgo](https://github.com/go-mgo/mgo) run the command below.

        ```
        go get gopkg.in/mgo.v2
        ```

Ussage Example
--------------
```
package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/codegangsta/negroni"
    "github.com/gorilla/context"
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
    mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        db := context.Get(request, 0)
        mgoSession := context.Get(request, 1)
        fmt.Fprintf(writer, "Within the handler")
        fmt.Println("db: ", db)
        fmt.Println("mgoSession: ", mgoSession)
    })

    // Register the middleware here.
    n.Use(nigronimgosession.NewDatabase(*dbAccessor).Middleware())

    n.UseHandler(mux)
    n.Run(":3000")
}
```
