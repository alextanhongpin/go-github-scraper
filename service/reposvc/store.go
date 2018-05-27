package reposvc

import (
	"sort"

	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/schema"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var stopwords = []string{"i", "me", "my", "myself", "we", "our", "ours", "ourselves", "you", "your", "yours", "yourself", "yourselves", "he", "him", "his", "himself", "she", "her", "hers", "herself", "it", "its", "itself", "they", "them", "their", "theirs", "themselves", "what", "which", "who", "whom", "this", "that", "these", "those", "am", "is", "are", "was", "were", "be", "been", "being", "have", "has", "had", "having", "do", "does", "did", "doing", "a", "an", "the", "and", "but", "if", "or", "because", "as", "until", "while", "of", "at", "by", "for", "with", "about", "against", "between", "into", "through", "during", "before", "after", "above", "below", "to", "from", "up", "down", "in", "out", "on", "off", "over", "under", "again", "further", "then", "once", "here", "there", "when", "where", "why", "how", "all", "any", "both", "each", "few", "more", "most", "other", "some", "such", "no", "nor", "not", "only", "own", "same", "so", "than", "too", "very", "s", "t", "can", "will", "just", "don", "should", "now"}

// Store provides the interface for the repo store
type (
	Store interface {
		Init() error
		FindOne(nameWithOwner string) (*schema.Repo, error)
		FindAll(limit int, sort []string) ([]schema.Repo, error)
		FindLastCreatedByUser(login string) (*schema.Repo, error)
		Upsert(github.Repo) error
		BulkUpsert(repos []github.Repo) error
		Count() (int, error)
		AggregateLanguages(limit int) ([]schema.LanguageCount, error)
		AggregateReposByUser(limit int) ([]schema.UserCount, error)
		AggregateLanguageByUser(login string, limit int) ([]schema.LanguageCount, error)
		AggregateMostRecentReposByLanguage(language string, limit int) ([]schema.Repo, error)
		AggregateReposByLanguage(language string, limit int) ([]schema.UserCount, error)
		Drop() error
		WatchersFor(login string) (int64, error)
		StargazersFor(login string) (int64, error)
		ForksFor(login string) (int64, error)
		KeywordsFor(login string, limit int) ([]schema.Keyword, error)
		DistinctLogin() ([]string, error)
	}

	// store is a struct that holds store configuration
	store struct {
		db         *database.DB
		collection string
	}
)

// NewStore returns a new store
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
		Key:    []string{"nameWithOwner"},
		Unique: true,
	})
}

func (s *store) FindOne(nameWithOwner string) (*schema.Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var repo schema.Repo
	if err := c.Find(bson.M{"nameWithOwner": nameWithOwner}).
		One(&repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

func (s *store) FindAll(limit int, sort []string) ([]schema.Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var repos []schema.Repo
	if err := c.Find(bson.M{}).
		Sort(sort...).
		Limit(limit).
		All(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func (s *store) FindLastCreatedByUser(login string) (*schema.Repo, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	var repo schema.Repo
	if err := c.Find(bson.M{
		"login": login,
	}).
		Sort("-createdAt").
		One(&repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

func (s *store) Upsert(repo github.Repo) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	if _, err := c.Upsert(
		bson.M{"nameWithOwner": repo.NameWithOwner},
		bson.M{
			"$set": repo.BSON(),
		},
	); err != nil {
		return err
	}
	return nil
}

func (s *store) BulkUpsert(repos []github.Repo) error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	bulk := c.Bulk()
	for _, repo := range repos {
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

	return nil
}

func (s *store) Count() (int, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return c.Count()
}

func (s *store) AggregateLanguages(limit int) ([]schema.LanguageCount, error) {
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
	if err := c.Pipe(pipeline).All(&languages); err != nil {
		return nil, err
	}
	return languages, nil
}

// AggregateReposByUser will return the repos by user
func (s *store) AggregateReposByUser(limit int) ([]schema.UserCount, error) {
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
	if err := c.Pipe(pipeline).All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// AggregateLanguageByUser returns the user's languages count
func (s *store) AggregateLanguageByUser(login string, limit int) ([]schema.LanguageCount, error) {
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
	if err := c.Pipe(pipeline).All(&languages); err != nil {
		return nil, err
	}
	return languages, nil
}

// AggregateMostRecentReposByLanguage returns the most recent repos by language
func (s *store) AggregateMostRecentReposByLanguage(language string, limit int) ([]schema.Repo, error) {
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
				"forkCount":     1,
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
	if err := c.Pipe(pipeline).All(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}

// AggregateReposByLanguage returns the users by languages
func (s *store) AggregateReposByLanguage(language string, limit int) ([]schema.UserCount, error) {
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
	if err := c.Pipe(pipeline).All(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func (s *store) Drop() error {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	return c.DropCollection()
}

func (s *store) WatchersFor(login string) (int64, error) {
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
			"$group": bson.M{
				"_id":   "$login",
				"count": bson.M{"$sum": "$stargazers"},
			},
		},
		bson.M{
			"$project": bson.M{
				"login": "$_id.login",
				"count": 1,
			},
		},
	}
	var watchers Watchers
	if err := c.Pipe(pipeline).One(&watchers); err != nil {
		if err == mgo.ErrNotFound {
			return 0, nil
		}
		return 0, err
	}
	return watchers.Count, nil
}

func (s *store) StargazersFor(login string) (int64, error) {

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
			"$group": bson.M{
				"_id":   "$login",
				"count": bson.M{"$sum": "$watchers"},
			},
		},
		bson.M{
			"$project": bson.M{
				"login": "$_id.login",
				"count": 1,
			},
		},
	}
	var stargazers Stargazers
	if err := c.Pipe(pipeline).One(&stargazers); err != nil {
		if err == mgo.ErrNotFound {
			return 0, nil
		}
		return 0, err
	}
	return stargazers.Count, nil
}

func (s *store) ForksFor(login string) (int64, error) {

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
			"$group": bson.M{
				"_id":   "$login",
				"count": bson.M{"$sum": "$forks"},
			},
		},
		bson.M{
			"$project": bson.M{
				"login": "$_id.login",
				"count": 1,
			},
		},
	}
	var forks Forks
	if err := c.Pipe(pipeline).One(&forks); err != nil {
		if err == mgo.ErrNotFound {
			return 0, nil
		}
		return 0, err
	}
	return forks.Count, nil
}

func (s *store) KeywordsFor(login string, limit int) ([]schema.Keyword, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()

	job := &mgo.MapReduce{
		Map: `function () {
			var stopwords = ["i", "me", "my", "myself", "we", "our", "ours", "ourselves", "you", "your", "yours", "yourself", "yourselves", "he", "him", "his", "himself", "she", "her", "hers", "herself", "it", "its", "itself", "they", "them", "their", "theirs", "themselves", "what", "which", "who", "whom", "this", "that", "these", "those", "am", "is", "are", "was", "were", "be", "been", "being", "have", "has", "had", "having", "do", "does", "did", "doing", "a", "an", "the", "and", "but", "if", "or", "because", "as", "until", "while", "of", "at", "by", "for", "with", "about", "against", "between", "into", "through", "during", "before", "after", "above", "below", "to", "from", "up", "down", "in", "out", "on", "off", "over", "under", "again", "further", "then", "once", "here", "there", "when", "where", "why", "how", "all", "any", "both", "each", "few", "more", "most", "other", "some", "such", "no", "nor", "not", "only", "own", "same", "so", "than", "too", "very", "s", "t", "can", "will", "just", "don", "should", "now"]
			var desc = this.description
			if (desc) {
				desc = desc.replace(/[^a-zA-Z ]/g, '') 
				desc = desc.toLowerCase().split(' ')
				desc = desc.filter(function (d) {
					return !stopwords.includes(d)
				})
        for (var i = desc.length - 1; i >= 0; i--) {
            // might want to remove punctuation, etc. here
            if (desc[i])  {      // make sure there's something
               emit(desc[i], 1); // store a 1 for each word
            }
        }
			}
		}`,
		Reduce: `function (key, values) {
			var count = 0
			values.forEach(function (v) {
				count += v
			})
			return count
		}`,
	}

	var res []schema.Keyword
	if _, err := c.Find(bson.M{
		"login": login,
	}).MapReduce(job, &res); err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return res, nil
	}

	sort.SliceStable(res, func(i, j int) bool {
		return res[i].Value > res[j].Value
	})
	return res[:min(len(res), 20)], nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// DistinctLogin returns a distinct login, so that we don't need to query the repository if the user does not exists
func (s *store) DistinctLogin() ([]string, error) {
	sess, c := s.db.Collection(s.collection)
	defer sess.Close()
	var res []string
	if err := c.Find(nil).Distinct("login", &res); err != nil {
		return nil, err
	}
	return res, nil
}
