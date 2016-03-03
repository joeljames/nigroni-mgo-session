package nigronimgosession

import (
	"net/http"

	"github.com/gorilla/context"
	mgo "gopkg.in/mgo.v2"
)

type DatabaseAccessor struct {
	*mgo.Session
	url  string
	name string
	coll string
}

func NewDatabaseAccessor(url, name, coll string) (*DatabaseAccessor, error) {
	session, err := mgo.Dial(url)
	if err == nil {
		return &DatabaseAccessor{session, url, name, coll}, nil
	} else {
		return &DatabaseAccessor{}, err
	}
}

func (da *DatabaseAccessor) Set(request *http.Request, session *mgo.Session) {
	db := session.DB(da.name)
	context.Set(request, "db", db)
	context.Set(request, "mgoSession", session)
}
