package profilesvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/moment"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// Store represents the methods that the store expose
	Store interface {
		Init() error
		GetProfiles() ([]schema.Profile, error)
		GetProfile(login string) (*schema.Profile, error)
		UpdateProfile(login string, profile schema.Profile) error
		BulkUpsert(profiles []schema.Profile) error
	}

	store struct {
		db         *database.DB
		collection string
	}
)

// NewStore returns store that fulfils the Store interface
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
		Key:    []string{"login"},
		Unique: true,
	})
}

func (s *store) GetProfiles() ([]schema.Profile, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var profiles []schema.Profile
	if err := c.Find(nil).All(&profiles); err != nil {
		return nil, err
	}
	return profiles, nil
}

func (s *store) GetProfile(login string) (*schema.Profile, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var profile schema.Profile
	if err := c.Find(bson.M{
		"login": login,
	}).One(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *store) UpdateProfile(login string, profile schema.Profile) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	if _, err := c.Upsert(bson.M{
		"login": login,
	}, bson.M{
		"$set": profile.BSON(),
		"$setOnInsert": bson.M{
			"createdAt": moment.NewUTCDate(),
		},
	}); err != nil {
		return err
	}
	return nil
}

func (s *store) BulkUpsert(profiles []schema.Profile) error {
	// sess, c := s.db.Collection(s.collection)
	// defer sess.Close()

	// // Mongo can only process a max of 1000 items
	// perBulk := 500
	// partitions, bucket := partitioner.New(perBulk, len(profiles))

	// for i := 0; i < bucket; i++ {
	// 	p := partitions[i]

	// 	bulk := c.Bulk()
	// 	for _, profile := range profiles[p.Start:p.End] {
	// 		bulk.Upsert(
	// 			bson.M{"login": profile.Login},
	// 			bson.M{
	// 				"$set": profile.BSON(),
	// 				"$setOnInsert": bson.M{
	// 					"createdAt": moment.NewUTCDate(),
	// 				},
	// 			},
	// 		)
	// 	}
	// 	if _, err := bulk.Run(); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
