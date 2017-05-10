nigroni-mgo-session
===================

`nigroni-mgo-session` is a simple Negroni middleware/handler for easy mgo session management in requests. It handles creating a new mongo session in the begning of a request and handles closing the mongo session at the end of request.


Getting Started
---------------
1. Before getting started make sure we have a copy of MongoDB running and is accessible. Running it locally would be recommended, but if thats not an option you can run on a remote server. Makes sure the mongo server is accessible, and you have a Mongo collection created within the database.
If you haven't installed `nigroni-mgo-session`, you can run the below command to install it.

    ```bash
    go get github.com/joeljames/nigroni-mgo-session
    ```

2. This package will be used with nigroni and mgo. If you haven't installed the dependencies you could install run the below command.

    * To install [Nigroni](https://github.com/codegangsta/negroni) run the command below.

        ```bash
        go get github.com/codegangsta/negroni
        ```
    * To install [mgo](https://github.com/go-mgo/mgo) run the command below.

        ```
        go get gopkg.in/mgo.v2
        ```

3. Now lets add an example which demonstrates the usage of this middleware. I have added comments to the example to explain the usage.

    * Now, that you have a running instance of MongoDB you can Shell into mongo by running the command below.
        ```bash
        mongo
        ```

    *  Create your database and collection and insert some data into the collection.

        ```mongo
        use db
        db.createCollection("widget");
        db.widget.insert({'id': '123'});
        db.widget.insert({'id': '456'});
        ```

    * Set your `DATABASE_URL` and `DATABASE_NAME` environment variable.

        ```bash
        export DATABASE_URL=mongodb://username:password@localhost:27017/db
        export DATABASE_NAME=db
        ```

    * If you have the package and dependency downloaded, you could simple copy the above example to a file `example.go` and then run the command `go run example.go` to start up the server.
    ```go
        package main

        import (
            "fmt"
            "net/http"
            "os"

            "github.com/codegangsta/negroni"
            nigronimgosession "github.com/joeljames/nigroni-mgo-session"
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
                // You can access the nms object from the request context.
                // The nms object is stored in key `nigronimgosession.KEY`.
                ctx := request.Context()
                nms := ctx.Value(nigronimgosession.KEY).(*nigronimgosession.NMS)
                // Now lets perform a count query using mgo db object.
                count, _ := nms.DB.C("widget").Find(nil).Count()

                fmt.Fprintf(writer, "Determining the count in the collection using the db object. \n\n")
                fmt.Fprintf(writer, "Total number of object in the mongo database: %d  \n\n", count)

                // You can access the mgo session object from the request object.
                // The session object is stored in key `mgoSession`.
                count2, _ := nms.Session.DB(dbName).C("widget").Find(nil).Count()
                fmt.Fprintf(writer, "Determining the count in the collection using the session object. \n\n")
                fmt.Fprintf(writer, "Total number of object in the mongo database: %d  \n\n", count2)
            })

            n.UseHandler(mux)
            n.Run(":3000")
        }
    ```

4. Assuming you are running the server locally, you can hit `http://localhost:3000/` to see the output.


Contribution
------------
1. Fork the package on Github
    ```bash
    git clone https://github.com/joeljames/nigroni-mgo-session.git
    cd nigroni-mgo-session
    ```

2. Create a new local branch to submit a pull request.
    ```bash
    git checkout -b name-of-feature
    ```

3. Commit your changes
    ```bash
    git commit -m "Detailed commit message"
    git push origin name-of-feature
    ```

4. Submit a pull request.
