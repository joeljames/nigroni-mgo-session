nigroni-mgo-session
===================

`nigroni-mgo-session` is a simple Negroni middleware/handler for easy mgo session management in requests. It handles creating a new mongo session in the begning of a request and handles closing the mongo session at the end of request.


Getting Started
---------------
1. Before getting started make sure we have a copy of MongoDB running and is accessible. Running it locally would be recommended, but if thats not an option you can run on a remote server. Makes sure the mongo server is accessible, and you have a Mongo collection created within the database.
If you haven't installed `nigroni-mgo-session`, you can run the below command to install it.

    ```
    go get github.com/joeljames/nigroni-mgo-session
    ```

2. Now, that you have a running instance of MongoDB you can set your `DATABASE_URL` and `DATABASE_NAME` environment variable.

    ```
    export DATABASE_URL=mongodb://username:password@localhost:27017/db
    export DATABASE_NAME=db
    ```

3. This package will be used with nigroni and mgo. If you haven't installed the dependencies you could install run the below command.

    * To install [Nigroni](https://github.com/codegangsta/negroni) run the command below.

        ```
        go get github.com/codegangsta/negroni
        ```
    * To install [Gorilla Context](https://github.com/gorilla/context) run the command below. This package is used to set variables on request object.

        ```
        go get github.com/gorilla/context
        ```
    * To install [mgo](https://github.com/go-mgo/mgo) run the command below.

        ```
        go get gopkg.in/mgo.v2
        ```

4. Now lets add an example which demonstrates the usage of this middleware. I have added comments to the example to explain the usage.

    ```
    package main

    import (
        "fmt"
        "net/http"
        "os"

        "github.com/codegangsta/negroni"
        "github.com/gorilla/context"
        nigronimgosession "github.com/joeljames/nigroni-mgo-session"
        mgo "gopkg.in/mgo.v2"
    )

    func main() {
        // Use the MongoDB `DATABASE_URL` from the env
        dbURL := os.Getenv("DATABASE_URL")
        // Use the MongoDB `DATABASE_NAME` from the env
        dbName := os.Getenv("DATABASE_NAME")
        // Set the MongoDB collection name
        dbColl := "widget"

        fmt.Println("Connecting to MongoDB: ", dbURL)
        fmt.Println("Database Name: ", dbName)
        fmt.Println("Collection Name: ", dbColl)

        // Creating the database accessor here.
        // Pointer to this database accessor will be passed to the middleware.
        dbAccessor, err := nigronimgosession.NewDatabaseAccessor(dbURL, dbName, dbColl)
        if err != nil {
            panic(err)
        }

        n := negroni.Classic()

        // Registering the middleware here.
        n.Use(nigronimgosession.NewDatabase(*dbAccessor).Middleware())

        mux := http.NewServeMux()
        mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
            // You can access the mgo db object from the request object.
            // The db object is stored in key `db`.
            db := context.Get(request, "db").(*mgo.Database)
            // Now lets perform a count query using mgo db object.
            count, _ := db.C("widget").Find(nil).Count()
            fmt.Fprintf(writer, "Determining the count in the collection using the db object. \n\n")
            fmt.Fprintf(writer, "Total number of object in the mongo database: %d  \n\n", count)

            // You can access the mgo session object from the request object.
            // The session object is stored in key `mgoSession`.
            mgoSession := context.Get(request, "mgoSession").(*mgo.Session)
            count2, _ := mgoSession.DB(dbName).C("widget").Find(nil).Count()
            fmt.Fprintf(writer, "Determining the count in the collection using the session object. \n\n")
            fmt.Fprintf(writer, "Total number of object in the mongo database: %d  \n\n", count2)

        })

        n.UseHandler(mux)
        n.Run(":3000")
    }
    ```

5. If you have the package and dependency downloaded, you could simple copy the above example to a file `example.go` and then run the command `go run example.go` to start up the server.

6. Assuming you are running the server locally, you can hit `http://localhost:3000/` to see the output.


Contribution
------------
1. Fork the package on Github
    ```
    $ git clone https://github.com/joeljames/nigroni-mgo-session.git
    $ cd nigroni-mgo-session
    ```

2. Create a new local branch to submit a pull request.
    ```
    $ git checkout -b name-of-feature
    ```

3. Commit your changes
    ```
    $ git commit -m "Detailed commit message"
    $ git push origin name-of-feature
    ```

4. Submit a pull request.
