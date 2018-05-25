package user

import (
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
	"github.com/alextanhongpin/go-github-scraper/internal/util"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Store provides the interface for the Service struct
type Store interface {
	Init() error
	FindOne(login string) (*User, error)
	FindAll(limit int, sort []string) ([]User, error)
	Upsert(schema.User) error
	BulkUpsert(users []schema.User) error
	Count() (int, error)
}

// store is a struct that holds service configuration
type store struct {
	db         *database.DB
	collection string
}

// New returns a new service
func New(db *database.DB, collection string) Store {
	return &store{
		db:         db,
		collection: collection,
	}
}

func (s *store) Init() error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	return c.EnsureIndex(mgo.Index{
		Key:    []string{"login"},
		Unique: true,
	})
}

func (s *store) FindOne(login string) (*User, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var user User
	if err := c.Find(bson.M{"login": login}).
		One(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *store) FindAll(limit int, sort []string) ([]User, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var users []User
	if err := c.Find(bson.M{}).
		Sort(sort...).
		Limit(limit).
		All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *store) Upsert(user schema.User) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	if _, err := c.Upsert(
		bson.M{"login": user.Login},
		bson.M{
			"$set": user.BSON(),
			"$setOnInsert": bson.M{
				"createdAt": util.NewUTCDate(),
				"fetchedAt": util.NewUTCDate(),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func (s *store) BulkUpsert(users []schema.User) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	bulk := c.Bulk()
	for _, user := range users {
		bulk.Upsert(
			bson.M{"login": user.Login},
			bson.M{
				"$set": user.BSON(),
				"$setOnInsert": bson.M{
					"createdAt": util.NewUTCDate(),
					"fetchedAt": util.NewUTCDate(),
				},
			},
		)
	}
	if _, err := bulk.Run(); err != nil {
		return err
	}

	return nil
}

func (s *store) Count() (int, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return c.Count()
}
