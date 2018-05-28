package analyticsvc

import (
	"errors"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
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

		GetReposMostStars() (*ReposMostStars, error)
		PostReposMostStars(repos []schema.Repo) error

		GetMostPopularLanguage() (*MostPopularLanguage, error)
		PostMostPopularLanguage(languages []schema.LanguageCount) error

		GetLanguageCountByUser() (*LanguageCountByUser, error)
		PostLanguageCountByUser(languages []schema.LanguageCount) error

		GetMostRecentReposByLanguage() (*MostRecentReposByLanguage, error)
		PostMostRecentReposByLanguage(repos []schema.RepoLanguage) error

		GetReposByLanguage() (*ReposByLanguage, error)
		PostReposByLanguage(users []schema.UserCountByLanguage) error
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
	return upsert(c, EnumUserCount, bson.M{
		"count":     count,
		"updatedAt": util.NewUTCDate(),
	})
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
	return upsert(c, EnumRepoCount, bson.M{
		"count":     count,
		"updatedAt": util.NewUTCDate(),
	})
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
	return upsert(c, EnumReposMostRecent, bson.M{
		"repos":     repos,
		"updatedAt": util.NewUTCDate(),
	})
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
	return upsert(c, EnumRepoCountByUser, bson.M{
		"users":     users,
		"updatedAt": util.NewUTCDate(),
	})
}

func (s *store) GetReposMostStars() (*ReposMostStars, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res ReposMostStars
	if err := c.
		Find(bson.M{"type": EnumReposMostStars}).
		One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *store) PostReposMostStars(repos []schema.Repo) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return upsert(c, EnumReposMostStars, bson.M{
		"repos":     repos,
		"updatedAt": util.NewUTCDate(),
	})
}

func (s *store) GetMostPopularLanguage() (*MostPopularLanguage, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res MostPopularLanguage
	if err := c.
		Find(bson.M{"type": EnumMostPopularLanguage}).
		One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *store) PostMostPopularLanguage(languages []schema.LanguageCount) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return upsert(c, EnumMostPopularLanguage, bson.M{
		"languages": languages,
		"updatedAt": util.NewUTCDate(),
	})
}

func (s *store) GetLanguageCountByUser() (*LanguageCountByUser, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res LanguageCountByUser
	if err := c.
		Find(bson.M{"type": EnumLanguageCountByUser}).
		One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *store) PostLanguageCountByUser(languages []schema.LanguageCount) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return upsert(c, EnumLanguageCountByUser, bson.M{
		"languages": languages,
		"updatedAt": util.NewUTCDate(),
	})
}

func (s *store) GetMostRecentReposByLanguage() (*MostRecentReposByLanguage, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res MostRecentReposByLanguage
	if err := c.
		Find(bson.M{"type": EnumMostRecentReposByLanguage}).
		One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *store) PostMostRecentReposByLanguage(repos []schema.RepoLanguage) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return upsert(c, EnumMostRecentReposByLanguage, bson.M{
		"repos":     repos,
		"updatedAt": util.NewUTCDate(),
	})
}

func (s *store) GetReposByLanguage() (*ReposByLanguage, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res ReposByLanguage
	if err := c.
		Find(bson.M{"type": EnumReposByLanguage}).
		One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *store) PostReposByLanguage(users []schema.UserCountByLanguage) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return upsert(c, EnumReposByLanguage, bson.M{
		"users":     users,
		"updatedAt": util.NewUTCDate(),
	})
}

func upsert(c *mgo.Collection, enum string, data bson.M) error {
	if _, err := c.Upsert(
		bson.M{"type": enum},
		bson.M{
			"$set": data,
			"$setOnInsert": bson.M{
				"createdAt": util.NewUTCDate(),
			},
		},
	); err != nil {
		return err
	}
	return nil
}
