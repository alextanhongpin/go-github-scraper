package user

import (
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Service provides the interface for the Service struct
type Service interface {
	Init() error
	FindOne(login string) (*model.User, error)
	FindAll(limit int) ([]model.User, error)
	Upsert(model.User) error
	BulkUpsert(users []model.User) error
}

// service is a struct that holds service configuration
type service struct {
	db *database.DB
	collection string
}

// New returns a new service
func New(db *database.DB) Service {
	return &service{
		db: db,
		collection: "users"
	}
}

func (s *service) Init() error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	return c.EnsureIndex(mgo.Index{
		Key:    []string{"login"},
		Unique: true,
	})
}

func (s *service) FindOne(login string) (*model.User, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var user model.User
	if err := c.Find(bson.M{"login": login}).
		// Select(bson.M{"login": 1}).
		One(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) FindAll(limit int) ([]model.User, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var users []model.User
	if err := c.Find(bson.M{}).
		Limit(limit).
		All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Upsert(user model.User) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	if _, err := c.Upsert(
		bson.M{"login": user.Login},
		bson.M{"$set": userToBSON(user)},
	); err != nil {
		return err
	}
	return nil
}

func (s *service) BulkUpsert(users []model.User) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	bulk := c.Bulk()
	for _, user := range users {
		bulk.Upsert(
			bson.M{"login": user.Login},
			bson.M{"$set": userToBSON(user)},
		)
	}
	if _, err := bulk.Run(); err != nil {
		return err
	}

	return nil
}

func userToBSON(user model.User) bson.M {
	return bson.M{
		"name":       user.Name,
		"createdAt":  user.CreatedAt,
		"updatedAt":  user.UpdatedAt,
		"login":      user.Login,
		"bio":        user.Bio,
		"location":   user.Location,
		"email":      user.Email,
		"company":    user.Company,
		"avatarUrl":  user.AvatarURL,
		"websiteUrl": user.WebsiteURL,
	}
}
