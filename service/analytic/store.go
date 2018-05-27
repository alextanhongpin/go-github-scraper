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
		GetUserCount() (*schema.UserCount, error)
		PostUserCount(count int) error
		GetRepoCount() (*RepoCount, error)
		PostRepoCount(count int) error
		GetReposMostRecent() (*ReposMostRecent, error)
		PostReposMostRecent(data []schema.Repo) error
		GetRepoCountByUser() (*RepoCountByUser, error)
		PostRepoCountByUser(repos []schema.UserCount) error
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

func (s *store) GetUserCount() (*schema.UserCount, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res schema.UserCount
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

func (s *store) GetRepoCountByUser() (*RepoCountByUser, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res RepoCountByUser
	if err := c.
		Find(bson.M{"type": EnumRepoCountByUser}).
		One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *store) PostRepoCountByUser(users []schema.UserCount) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	if _, err := c.Upsert(
		bson.M{"type": EnumRepoCountByUser},
		bson.M{
			"$set": bson.M{
				"users":     users,
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

// func (s *store) GetReposMostStars(data) error  {}
// func (s *store) PostReposMostStars(data) error {}

// func (s *store) GetMostPopularLanguage(data) error  {}
// func (s *store) PostMostPopularLanguage(data) error {}

// func (s *store) GetLanguageCountByUser(data) error  {}
// func (s *store) PostLanguageCountByUser(data) error {}

// func (s *store) GetMostRecentReposByLanguage(data) error  {}
// func (s *store) PostMostRecentReposByLanguage(data) error {}

// func (s *store) GetReposByLanguage(data) error  {}
// func (s *store) PostReposByLanguage(data) error {}
