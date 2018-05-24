package main

import (
	"flag"
	"log"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/database"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Login     string    `json:"login,omitempty" bson:"login,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Count     int64     `json:"count,omitempty"`
}

func main() {
	githubToken := flag.String("github_token", "", "The Github's access token used to make calls to the GraphQL endpoint")
	dbName := flag.String("db_name", "scraper", "The name of the database")
	dbHost := flag.String("db_host", "mongodb://myuser:mypass@localhost:27017", "The hostname of the database")

	_ = githubToken
	flag.Parse()

	db := database.New(*dbHost, *dbName)
	defer db.Close()

	// Create a new collection with the session
	sess, users := db.Collection("users")
	defer sess.Close()

	err := users.EnsureIndex(mgo.Index{
		Key:    []string{"login"},
		Unique: true,
	})

	if err != nil {
		panic(err)
	}

	bulk := users.Bulk()
	bulk.Upsert(bson.M{"login": "johndoe"}, bson.M{"$set": bson.M{"count": 1}})
	bulk.Upsert(bson.M{"login": "alextanhongpin"}, bson.M{"$set": bson.M{"count": 10}})
	bulk.Upsert(bson.M{"login": "hello"}, bson.M{"$set": bson.M{"count": 10}})
	change, err := bulk.Run()
	if err != nil {
		log.Println(err)
	}
	log.Println(change)

	// Sort by timestamp
	// err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	// change, err := users.Upsert(
	// 	bson.M{"login": "alextanhongpin"},
	// 	bson.M{"$set": bson.M{
	// 		"timestamp": time.Now(),
	// 		"count":     10,
	// 	}},
	// )
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("change", change)

	// var user User
	// if err = users.Find(bson.M{"login": "alextanhongpin"}).
	// 	Select(bson.M{"login": 0}).
	// 	One(&user); err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("user: %+v\n", user)

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

	// gapi := github.New(*githubToken, "https://api.github.com/graphql", "Malaysia")
	// cursor := ""
	// hasNextPage := true
	// var users []model.User
	// for hasNextPage {
	// 	res, err := gapi.FetchUsers("2018-01-01", "2018-02-01", cursor)
	// 	if err != nil {
	// 		break
	// 	}
	// 	log.Println("got user count:", res.Data.Search.UserCount)
	// 	hasNextPage = res.Data.Search.PageInfo.HasNextPage
	// 	cursor = res.Data.Search.PageInfo.EndCursor
	// 	for _, edge := range res.Data.Search.Edges {
	// 		users = append(users, edge.Node)
	// 	}
	// }
	// log.Println(len(users))
}
