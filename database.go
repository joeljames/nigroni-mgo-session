package nigronimgosession

import (
	"net/http"

	"github.com/codegangsta/negroni"
)

type Database struct {
	dba DatabaseAccessor
}

func NewDatabase(databaseAccessor DatabaseAccessor) *Database {
	return &Database{databaseAccessor}
}

func (d *Database) Middleware() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
		reqSession := d.dba.Session.Clone()
		defer reqSession.Close()
		d.dba.Set(request, reqSession)
		next(rw, request)
	}
}
