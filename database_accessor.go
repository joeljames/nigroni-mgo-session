package nigronimgosession

import (
	"context"
	"net/http"
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

func (da *DatabaseAccessor) Set(ctx context.Context, request *http.Request, session *mgo.Session) context.Context{
	db := session.DB(da.name)
	nms := &NMS{db, session}
	return context.WithValue(ctx, KEY, nms)
}
