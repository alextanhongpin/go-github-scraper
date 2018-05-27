package analyticsvc

import (
	"errors"

	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
	"github.com/alextanhongpin/go-github-scraper/internal/util"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var errTypeAssertion = errors.New("unable to perform type assertion")

// Store represents the interface for the analytic store
type (
	Store interface {
		Init() error
		GetUserCount() (*UserCount, error)
		PostUserCount(count int) error
		GetRepoCount() (*RepoCount, error)
		PostRepoCount(count int) error
		GetReposMostRecent() (*ReposMostRecent, error)
		PostReposMostRecent(data []schema.Repo) error
	}

	store struct {
		db         *database.DB
		collection string
	}
)

// NewStore returns a new analytic
func NewStore(db *database.DB, collection string) Store {
	return &store{
		db:         db,
		collection: collection,
	}
}

func (s *store) Init() error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	return c.EnsureIndex(mgo.Index{
		Key:    []string{"type"},
		Unique: true,
	})
}

func (s *store) GetUserCount() (*UserCount, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res UserCount
	if err := c.
		Find(bson.M{"type": EnumUserCount}).
		One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *store) PostUserCount(count int) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	if _, err := c.Upsert(
		bson.M{"type": EnumUserCount},
		bson.M{
			"$set": bson.M{
				"count":     count,
				"updatedAt": util.NewUTCDate(),
			},
			"$setOnInsert": bson.M{
				"createdAt": util.NewUTCDate(),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func (s *store) GetRepoCount() (*RepoCount, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res RepoCount
	if err := c.
		Find(bson.M{"type": EnumRepoCount}).
		One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *store) PostRepoCount(count int) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	if _, err := c.Upsert(
		bson.M{"type": EnumRepoCount},
		bson.M{
			"$set": bson.M{
				"count":     count,
				"updatedAt": util.NewUTCDate(),
			},
			"$setOnInsert": bson.M{
				"createdAt": util.NewUTCDate(),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func (s *store) GetReposMostRecent() (*ReposMostRecent, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res ReposMostRecent
	if err := c.
		Find(bson.M{"type": EnumReposMostRecent}).
		One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *store) PostReposMostRecent(repos []schema.Repo) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	if _, err := c.Upsert(
		bson.M{"type": EnumReposMostRecent},
		bson.M{
			"$set": bson.M{
				"repos":     repos,
				"updatedAt": util.NewUTCDate(),
			},
			"$setOnInsert": bson.M{
				"createdAt": util.NewUTCDate(),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

// func (s *store) GetEnumRepoCountByUser(data) error  {}
// func (s *store) PostEnumRepoCountByUser(data) error {}

// func (s *store) GetEnumReposMostStars(data) error  {}
// func (s *store) PostEnumReposMostStars(data) error {}

// func (s *store) GetEnumMostPopularLanguage(data) error  {}
// func (s *store) PostEnumMostPopularLanguage(data) error {}

// func (s *store) GetEnumLanguageCountByUser(data) error  {}
// func (s *store) PostEnumLanguageCountByUser(data) error {}

// func (s *store) GetEnumMostRecentReposByLanguage(data) error  {}
// func (s *store) PostEnumMostRecentReposByLanguage(data) error {}

// func (s *store) GetEnumReposByLanguage(data) error  {}
// func (s *store) PostEnumReposByLanguage(data) error {}
