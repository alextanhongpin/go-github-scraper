package repo

import (
	"log"

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
	AggregateLanguages(limit int) ([]Language, error)
	AggregateReposByUser(limit int) ([]User, error)
	AggregateLanguageByUser(login string, limit int) error
	AggregateMostRecentReposByLanguage(language string, limit int) error
	AggregateUsersByLanguage(language string, limit int) error
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

func (s *store) AggregateLanguages(limit int) ([]Language, error) {
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
	var languages []Language
	if err := c.Pipe(pipeline).All(&languages); err != nil {
		return nil, err
	}
	return languages, nil
}

// AggregateReposByUser will return the repos by user
func (s *store) AggregateReposByUser(limit int) ([]User, error) {
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
				"name":  "$_id.login",
			},
		},
	}
	var users []User
	if err := c.Pipe(pipeline).All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// AggregateLanguageByUser returns the user's languages count
func (s *store) AggregateLanguageByUser(login string, limit int) error {
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
	var languages []Language
	if err := c.Pipe(pipeline).All(&languages); err != nil {
		log.Println(err)
		return err
	}
	log.Println(languages)
	return nil
}

// AggregateMostRecentReposByLanguage returns the most recent repos by language
func (s *store) AggregateMostRecentReposByLanguage(language string, limit int) error {
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
				"name":        1,
				"description": 1,
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
	var languages []interface{}
	if err := c.Pipe(pipeline).All(&languages); err != nil {
		log.Println(err)
		return err
	}
	log.Println(languages)
	return nil
}

// AggregateUsersByLanguage returns the users by languages
func (s *store) AggregateUsersByLanguage(language string, limit int) error {
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
				"name":        1,
				"description": 1,
				"login":       1,
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
		// bson.M{
		// 	"$project": bson.M{
		// 		"count": 1,
		// 		"name":  "$_id.language",
		// 		"_id":   0,
		// 	},
		// },
	}
	var languages []interface{}
	if err := c.Pipe(pipeline).All(&languages); err != nil {
		log.Println(err)
		return err
	}
	log.Println(languages)
	return nil
}
