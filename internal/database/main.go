package database

import mgo "gopkg.in/mgo.v2"

type Database struct {
	Session *mgo.Session
	Name    string
}

func (db *Database) Close() {
	db.Session.Close()
}

func (db *Database) Collection(name string) (*mgo.Session, *mgo.Collection) {
	sess := db.Session.Copy()
	coll := sess.DB(db.Name).C(name)
	return sess, coll
}

func New(mongoURI, name string) *Database {
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		panic(err)
	}

	// defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	return &Database{
		Session: session,
		Name:    name,
	}
}
