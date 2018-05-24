package repo

import (
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Service provides the interface for the Service struct
type Service interface {
	Init() error
	FindOne(login string) (*model.Repo, error)
	FindAll(limit int) ([]model.Repo, error)
	Upsert(model.Repo) error
	BulkUpsert(repos []model.Repo) error
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
		collection: "repos"
	}
}

func (s *service) Init() error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	return c.EnsureIndex(mgo.Index{
		Key:    []string{"nameWithOwner"},
		Unique: true,
	})
}

func (s *service) FindOne(login string) (*model.Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var repo model.Repo
	if err := c.Find(bson.M{"nameWithOwner": login}).
		One(&repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

func (s *service) FindAll(limit int) ([]model.Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var repos []model.Repo
	if err := c.Find(bson.M{}).
		Limit(limit).
		All(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func (s *service) Upsert(repo model.Repo) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	if _, err := c.Upsert(
		bson.M{"nameWithOwner": repo.NameWithOwner},
		bson.M{"$set": repoToBSON(repo)},
	); err != nil {
		return err
	}
	return nil
}

func (s *service) BulkUpsert(repos []model.Repo) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	bulk := c.Bulk()
	for _, repo := range repos {
		bulk.Upsert(
			bson.M{"nameWithOwner": repo.NameWithOwner},
			bson.M{"$set": repoToBSON(repo)},
		)
	}
	if _, err := bulk.Run(); err != nil {
		return err
	}

	return nil
}

func repoToBSON(repo model.Repo) bson.M {
	var languages []string
	for _, lang := range repo.Languages.Edges {
		languages = append(languages, lang.Node.Name)
	}
	return bson.M{
		"name":       repo.Name,
		"createdAt":  repo.CreatedAt,
		"updatedAt":  repo.UpdatedAt,
		"description": repo.Description,
		"languages": languages,
		"homepageUrl": repo.HomepageURL,
		"forkCount": repo.ForkCount,
		"isFork": repo.IsFork,
		"nameWithOwner": repo.NameWithOwner,
		"stargazers_count": repo.Stargazers.TotalCount,
		"watchers_count": repo.Watchers.TotalCount,
		"login": repo.Owner.Login,
		"avatar_url": repo.Owner.AvatarURL,
	}
}
