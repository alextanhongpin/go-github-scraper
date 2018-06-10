package reposvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/partitioner"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// While it is not necessary, separating the read and write can be beneficial
// when you need to use two different database access which has constraints
// on the read and write

type (
	// Read defines all read operations by the store
	Read interface {
		Count() (int, error)
		Distinct(field string) ([]string, error)
		FindAll(limit int, sort []string) ([]schema.Repo, error)
		GroupByLanguage(language string, limit int) ([]schema.UserCount, error)
		GroupByLanguageSortByMostRecent(language string, limit int) ([]schema.Repo, error)
		GroupByUser(limit int) ([]schema.UserCount, error)
		Languages(limit int) ([]schema.LanguageCount, error)
		LanguagesBy(login string, limit int) ([]schema.LanguageCount, error)
		LastCreatedBy(login string) (*schema.Repo, error)
		ReposBy(login string) ([]schema.Repo, error)
	}

	// Write defines the write operation for the store
	Write interface {
		BulkUpsert(repos []github.Repo) error
		Init() error
		Drop() error
	}

	// Store represents the repository pattern that implements the Read and Write interface
	Store interface {
		Read
		Write
	}

	// store is a struct that holds store configuration
	store struct {
		db         *database.DB
		collection string
	}
)

// NewStore returns a new store
func NewStore(db *database.DB, collection string) Store {
	return &store{db, collection}
}

func (s *store) Init() error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	return c.EnsureIndex(mgo.Index{
		Key:    []string{"nameWithOwner"},
		Unique: true,
	})
}

func (s *store) Drop() error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	return c.DropCollection()
}

func (s *store) BulkUpsert(repos []github.Repo) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	perBulk := 500
	partitions, bucket := partitioner.New(perBulk, len(repos))

	for i := 0; i < bucket; i++ {
		p := partitions[i]

		bulk := c.Bulk()
		for _, repo := range repos[p.Start:p.End] {
			bulk.Upsert(
				bson.M{"nameWithOwner": repo.NameWithOwner},
				bson.M{
					"$set": repo.BSON(),
				},
			)
		}
		if _, err := bulk.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (s *store) FindAll(limit int, sort []string) ([]schema.Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var repos []schema.Repo
	err := c.Find(nil).
		Sort(sort...).
		Limit(limit).
		All(&repos)

	return repos, err
}

func (s *store) ReposBy(login string) ([]schema.Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var repos []schema.Repo

	query := bson.M{
		"login":  login,
		"isFork": false,
	}

	if err := c.Find(query).All(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func (s *store) LastCreatedBy(login string) (*schema.Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var repo schema.Repo
	query := bson.M{
		"login": login,
	}

	err := c.Find(query).
		Sort("-createdAt").
		One(&repo)

	return &repo, err
}

func (s *store) Count() (int, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return c.Count()
}

// Languages returns the languages that exists in sorted by frequency
func (s *store) Languages(limit int) ([]schema.LanguageCount, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"isFork": false,
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path": "$languages",
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"language": "$languages",
				},
				"count": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{
				"count": -1,
			},
		},
		bson.M{
			"$limit": limit,
		},
		bson.M{
			"$project": bson.M{
				"count": 1,
				"name":  "$_id.language",
				"_id":   0,
			},
		},
	}
	var languages []schema.LanguageCount
	err := c.Pipe(pipeline).All(&languages)

	return languages, err
}

// GroupByUser will return the repo count by user
func (s *store) GroupByUser(limit int) ([]schema.UserCount, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"isFork": false,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"login": "$login",
				},
				"count":     bson.M{"$sum": 1},
				"avatarUrl": bson.M{"$first": "$avatarUrl"},
			},
		},
		bson.M{
			"$sort": bson.M{
				"count": -1,
			},
		},
		bson.M{
			"$limit": limit,
		},
		bson.M{
			"$project": bson.M{
				"count":     1,
				"avatarUrl": 1,
				"name":      "$_id.login",
			},
		},
	}
	var users []schema.UserCount
	err := c.Pipe(pipeline).All(&users)
	return users, err
}

// LanguagesBy returns the user's languages count
func (s *store) LanguagesBy(login string, limit int) ([]schema.LanguageCount, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"isFork": false,
				"login":  login,
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path": "$languages",
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"language": "$languages",
				},
				"count": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{
				"count": -1,
			},
		},
		bson.M{
			"$limit": limit,
		},
		bson.M{
			"$project": bson.M{
				"count": 1,
				"name":  "$_id.language",
				"_id":   0,
			},
		},
	}
	var languages []schema.LanguageCount
	err := c.Pipe(pipeline).All(&languages)
	return languages, err
}

// GroupByLanguageSortByMostRecent returns the most recent repos by language
func (s *store) GroupByLanguageSortByMostRecent(language string, limit int) ([]schema.Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"isFork": false,
			},
		},
		bson.M{
			"$project": bson.M{
				"isLanguage": bson.M{
					"$in": []string{language, "$languages"},
				},
				"name":          1,
				"createdAt":     1,
				"updatedAt":     1,
				"description":   1,
				"languages":     1,
				"homepageUrl":   1,
				"forks":         1,
				"isFork":        1,
				"nameWithOwner": 1,
				"login":         1,
				"avatarUrl":     1,
				"stargazers":    1,
				"watchers":      1,
				"url":           1,
			},
		},
		bson.M{
			"$match": bson.M{
				"isLanguage": true,
			},
		},
		bson.M{
			"$sort": bson.M{
				"updatedAt": -1,
			},
		},
		bson.M{
			"$limit": limit,
		},
	}
	var repos []schema.Repo
	err := c.Pipe(pipeline).All(&repos)
	return repos, err
}

// GroupByLanguage returns the users by languages
func (s *store) GroupByLanguage(language string, limit int) ([]schema.UserCount, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"isFork": false,
			},
		},
		bson.M{
			"$project": bson.M{
				"isLanguage": bson.M{
					"$in": []string{language, "$languages"},
				},
				"login":     1,
				"avatarUrl": 1,
			},
		},
		bson.M{
			"$match": bson.M{
				"isLanguage": true,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"login": "$login",
				},
				"count":     bson.M{"$sum": 1},
				"avatarUrl": bson.M{"$first": "$avatarUrl"},
			},
		},
		bson.M{
			"$sort": bson.M{
				"count": -1,
			},
		},
		bson.M{
			"$limit": limit,
		},
		bson.M{
			"$project": bson.M{
				"count":     1,
				"avatarUrl": 1,
				"name":      "$_id.login",
			},
		},
	}
	var repos []schema.UserCount
	err := c.Pipe(pipeline).All(&repos)
	return repos, err
}

// Distinct returns a distinct login, so that we don't need to query the repository if the user does not exists
func (s *store) Distinct(field string) ([]string, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var res []string
	err := c.Find(nil).Distinct(field, &res)
	return res, err
}
