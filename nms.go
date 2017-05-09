package nigronimgosession

import "gopkg.in/mgo.v2"

type NMS struct {
	DB      *mgo.Database
	Session *mgo.Session
}
