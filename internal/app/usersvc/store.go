package usersvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/moment"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/partitioner"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Store provides the interface for the Service struct
type (
	Store interface {
		BulkUpsert(users []github.User) error
		BulkUpdate(users []User) error
		Count() (int, error)
		Drop() error
		FindOne(login string) (*User, error)
		FindAll(limit int, sort []string) ([]User, error)
		FindLastCreated() (*User, error)
		Init() error
		PickLogin() ([]string, error)
		UpdateOne(login string) error
		Upsert(github.User) error
		WithRepos(count int) ([]User, error)
	}

	// store is a struct that holds service configuration
	store struct {
		db         *database.DB
		collection string
	}
)

// NewStore returns a new service
func NewStore(db *database.DB, collection string) Store {
	return &store{
		db:         db,
		collection: collection,
	}
}

// Init performs initialization of the database including setting indices or creating tables
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

func (s *store) FindLastCreated() (*User, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var user User
	if err := c.Find(bson.M{}).
		Sort("-$createdAt").
		One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// PickLogin takes the login field
func (s *store) PickLogin() ([]string, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	pipeline := []bson.M{
		bson.M{
			"$group": bson.M{
				"_id": nil,
				"logins": bson.M{
					"$push": "$login",
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"items": "$logins",
			},
		},
	}

	var i Logins
	if err := c.Pipe(pipeline).One(&i); err != nil {
		return i.Items, err
	}
	return i.Items, nil
}

func (s *store) Upsert(user github.User) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	if _, err := c.Upsert(
		bson.M{"login": user.Login},
		bson.M{
			"$set": user.BSON(),
		},
	); err != nil {
		return err
	}
	return nil
}

func (s *store) BulkUpsert(users []github.User) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	perBulk := 500

	partitions, bucket := partitioner.New(perBulk, len(users))

	for i := 0; i < bucket; i++ {
		p := partitions[i]

		bulk := c.Bulk()
		for _, user := range users[p.Start:p.End] {
			bulk.Upsert(
				bson.M{"login": user.Login},
				bson.M{
					"$set": user.BSON(),
				},
			)
		}

		if _, err := bulk.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (s *store) BulkUpdate(users []User) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	// Mongo can only process a max of 1000 items
	perBulk := 500
	partitions, bucket := partitioner.New(perBulk, len(users))

	for i := 0; i < bucket; i++ {
		p := partitions[i]

		bulk := c.Bulk()
		for _, user := range users[p.Start:p.End] {
			bulk.Upsert(
				bson.M{"login": user.Login},
				bson.M{
					"$set": user.Profile.BSON(),
				},
			)
		}
		if _, err := bulk.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (s *store) Count() (int, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return c.Count()
}

func (s *store) Drop() error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return c.DropCollection()
}

func (s *store) UpdateOne(login string) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	_, err := c.Upsert(bson.M{"login": login},
		bson.M{
			"$set": bson.M{
				"fetchedAt": moment.NewUTCDate(),
			},
		})
	if err != nil {
		return err
	}
	return nil
}

func (s *store) WithRepos(count int) ([]User, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var users []User
	if err := c.Find(bson.M{
		"repositories": bson.M{
			"$gt": count,
		},
	}).All(&users); err != nil {
		return nil, err
	}

	return users, nil
}
