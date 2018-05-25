package repo

import (
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
	"github.com/alextanhongpin/go-github-scraper/internal/util"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Store provides the interface for the repo store
type Store interface {
	Init() error
	FindOne(login string) (*Repo, error)
	FindAll(limit int, sort []string) ([]Repo, error)
	Upsert(schema.Repo) error
	BulkUpsert(repos []schema.Repo) error
	Count() (int, error)
}

// store is a struct that holds store configuration
type store struct {
	db         *database.DB
	collection string
}

// New returns a new store
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
		Key:    []string{"nameWithOwner"},
		Unique: true,
	})
}

func (s *store) FindOne(login string) (*Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var repo Repo
	if err := c.Find(bson.M{"nameWithOwner": login}).
		One(&repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

func (s *store) FindAll(limit int, sort []string) ([]Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var repos []Repo
	if err := c.Find(bson.M{}).
		Sort(sort...).
		Limit(limit).
		All(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func (s *store) Upsert(repo schema.Repo) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	if _, err := c.Upsert(
		bson.M{"nameWithOwner": repo.NameWithOwner},
		bson.M{
			"$set": repo.BSON(),
			"$setOnInsert": bson.M{
				"createdAt": util.NewUTCDate(),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func (s *store) BulkUpsert(repos []schema.Repo) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	bulk := c.Bulk()
	for _, repo := range repos {
		bulk.Upsert(
			bson.M{"nameWithOwner": repo.NameWithOwner},
			bson.M{
				"$set": repo.BSON(),
				"$setOnInsert": bson.M{
					"createdAt": util.NewUTCDate(),
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
