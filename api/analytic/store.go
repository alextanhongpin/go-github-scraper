package analytic

import (
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"gopkg.in/mgo.v2/bson"
)

// Store represents the interface for the analytic store
type (
	Store interface {
		GetUserCount() (*UserCount, error)
		PostUserCount(count int) error
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

func (s *store) GetUserCount() (*UserCount, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var out UserCount
	if err := c.Find(bson.M{"type": EnumUserCount.String()}).One(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *store) PostUserCount(count int) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	if _, err := c.Upsert(
		bson.M{"type": EnumUserCount.String()},
		bson.M{
			"$set": bson.M{
				"count":     count,
				"updatedAt": time.Now().UTC().Format(time.RFC3339),
			},
			"$setOnInsert": bson.M{
				"createdAt": time.Now().UTC().Format(time.RFC3339),
			},
		},
	); err != nil {
		return err
	}
	return nil
}
