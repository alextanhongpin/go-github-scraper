package database

import (
	mgo "gopkg.in/mgo.v2"
)

// DB holds the session to the mongo db
type DB struct {
	Session *mgo.Session
	Name    string
}

// Close close the global database session
func (db *DB) Close() {
	db.Session.Close()
}

// Collection creates a copy of the session and returns the collection and the session
func (db *DB) Collection(name string) (*mgo.Session, *mgo.Collection) {
	sess := db.Session.Copy()
	coll := sess.DB(db.Name).C(name)
	return sess, coll
}

// New returns a new pointer to the DB struct
func New(mongoURI, name string) *DB {
	s, err := mgo.Dial(mongoURI)
	if err != nil {
		panic(err)
	}

	s.SetMode(mgo.Monotonic, true)

	return &DB{
		Session: s,
		Name:    name,
	}
}
