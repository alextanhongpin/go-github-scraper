package user

import (
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

// Service provides the interface for the Service struct
type Service interface {
	FindOne(login string) (*model.User, error)
	BulkInsert(users []model.User) error
}

// service is a struct that holds service configuration
type service struct {
	db *database.DB
}

// New returns a new service
func New(db *database.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) Init() error {
	sess, c := s.db.Collection("users")
	defer sess.Close()

	return c.EnsureIndex(mgo.Index{
		Key:    []string{"login"},
		Unique: true,
	})
}

func (s *service) FindOne(login string) (*model.User, error) {
	sess, c := s.db.Collection("users")
	defer sess.Close()

	var user model.User
	if err := c.Find(bson.M{"login": login}).
		// Select(bson.M{"login": 1}).
		One(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) BulkInsert(users []model.User) error {
	sess, c := s.db.Collection("users")
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

func (s *service) Insert(user model.User) error {
	sess, c := s.db.Collection("users")
	defer sess.Close()

	if _, err := c.Upsert(
		bson.M{"login": user.Login},
		bson.M{"$set": userToBSON(user)},
	); err != nil {
		return err
	}
	return nil
}

// // Sort by timestamp
// err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

// var userColl []User
// if err = users.Find(bson.M{}).All(&userColl); err != nil {
// 	log.Println(err)
// }
// log.Printf("user: %#v\n count: %d", userColl, len(userColl))

// change, err := users.RemoveAll(bson.M{})
// if err != nil {
// 	panic(err)
// }
// log.Println("change", change)
