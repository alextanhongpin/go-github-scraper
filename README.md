## API Calls

```bash
$ curl -H "Authorization: bearer token" -X POST -d " \
 { \
   \"query\": \"query { viewer { login }}\" \
 } \
" https://api.github.com/graphql
```

## Sample Search User Query

```
query {
  search(query: "location:malaysia created:2018-01-01..2018-01-10", type: USER, last: 10, after: "Y3Vyc29yOjIw") {
    userCount,
    pageInfo {
      hasNextPage,
      startCursor,
      endCursor,
      hasPreviousPage,
    },
    edges {
      cursor,
      node {
        ...on User {
          name,
          createdAt,
          updatedAt,
          login,
          bio,
          location,
          email,
          company,
          avatarUrl,
          websiteUrl
        }
      }
    }
  }
}
```

Output:

```json

{
  "data": {
    "search": {
      "userCount": 21,
      "pageInfo": {
        "hasNextPage": true,
        "startCursor": "Y3Vyc29yOjE=",
        "endCursor": "Y3Vyc29yOjEw",
        "hasPreviousPage": false
      },
      "edges": [
        {
          "cursor": "Y3Vyc29yOjE=",
          "node": {
            "name": "Philip Khor",
            "createdAt": "2018-01-02T17:36:06Z",
            "updatedAt": "2018-05-01T11:35:07Z",
            "login": "philip-khor",
            "bio": "Economics and data dabbler",
            "location": "Penang, Malaysia",
            "email": "",
            "company": "",
            "avatarUrl": "https://avatars2.githubusercontent.com/u/35039795?v=4",
            "websiteUrl": "philip-khor.github.io"
          }
        },
        {
          "cursor": "Y3Vyc29yOjI=",
          "node": {
            "name": "rohaizad",
            "createdAt": "2018-01-09T08:01:55Z",
            "updatedAt": "2018-05-10T04:31:17Z",
            "login": "e-zad",
            "bio": "",
            "location": "keningau sabah malaysia",
            "email": "",
            "company": "",
            "avatarUrl": "https://avatars1.githubusercontent.com/u/35253967?v=4",
            "websiteUrl": ""
          }
        },
        {
          "cursor": "Y3Vyc29yOjM=",
          "node": {
            "name": "Hoshito",
            "createdAt": "2018-01-03T16:50:05Z",
            "updatedAt": "2018-02-18T13:07:19Z",
            "login": "Hoshitter1",
            "bio": "Started learning Python on Jan 2018. Currently working on having twitter bot running on heroku (as of Feb'18)",
            "location": "Japan(until Feb'18) Malaysia(until Feb'21)",
            "email": "",
            "company": "",
            "avatarUrl": "https://avatars3.githubusercontent.com/u/35073785?v=4",
            "websiteUrl": "twitter.com/hoshitter1"
          }
        },
        {
          "cursor": "Y3Vyc29yOjQ=",
          "node": {
            "name": "yuzinc",
            "createdAt": "2018-01-01T03:54:16Z",
            "updatedAt": "2018-01-01T03:55:33Z",
            "login": "ciraxes",
            "bio": "",
            "location": "Malaysia",
            "email": "",
            "company": "",
            "avatarUrl": "https://avatars1.githubusercontent.com/u/34994077?v=4",
            "websiteUrl": ""
          }
        },
        {
          "cursor": "Y3Vyc29yOjU=",
          "node": {
            "name": "Izumi Inoue",
            "createdAt": "2018-01-07T09:39:31Z",
            "updatedAt": "2018-05-22T13:15:18Z",
            "login": "izumiinoue",
            "bio": "Founder @ Bereev | Non-Conforming Autodidact",
            "location": "Malaysia",
            "email": "",
            "company": "",
            "avatarUrl": "https://avatars0.githubusercontent.com/u/35189477?v=4",
            "websiteUrl": ""
          }
        },
        {
          "cursor": "Y3Vyc29yOjY=",
          "node": {
            "name": "Saifulke",
            "createdAt": "2018-01-09T14:00:43Z",
            "updatedAt": "2018-05-10T04:31:21Z",
            "login": "saifulke",
            "bio": "Tak ada yang abadi",
            "location": "Malaysia",
            "email": "",
            "company": "@corbemusic",
            "avatarUrl": "https://avatars3.githubusercontent.com/u/35264073?v=4",
            "websiteUrl": "https://www.youtube.com/channel/UCYGztGdclFSX_lRSuFhMEug"
          }
        },
        {
          "cursor": "Y3Vyc29yOjc=",
          "node": {
            "name": "Soo Kin Wah",
            "createdAt": "2018-01-08T12:13:02Z",
            "updatedAt": "2018-03-13T16:24:31Z",
            "login": "kinwah123456",
            "bio": "Always a learner! :D",
            "location": "Malaysia",
            "email": "",
            "company": "",
            "avatarUrl": "https://avatars0.githubusercontent.com/u/35223517?v=4",
            "websiteUrl": ""
          }
        },
        {
          "cursor": "Y3Vyc29yOjg=",
          "node": {
            "name": "Ariel Espina Mendoza",
            "createdAt": "2018-01-10T09:26:03Z",
            "updatedAt": "2018-05-21T16:45:41Z",
            "login": "aemendoza072583",
            "bio": "",
            "location": "Kuala Lumpur, Malaysia",
            "email": "",
            "company": "",
            "avatarUrl": "https://avatars0.githubusercontent.com/u/35292757?v=4",
            "websiteUrl": ""
          }
        },
        {
          "cursor": "Y3Vyc29yOjk=",
          "node": {
            "name": "nando teddy",
            "createdAt": "2018-01-04T03:16:28Z",
            "updatedAt": "2018-04-21T08:33:30Z",
            "login": "ndoteddy",
            "bio": "Proffesional Tech Engineer - Stay Fool - Stay Noob \r\n\r\n#JS #C#",
            "location": "Malaysia",
            "email": "",
            "company": "@teddybrothers",
            "avatarUrl": "https://avatars0.githubusercontent.com/u/35088152?v=4",
            "websiteUrl": "linkedin.com/in/hernandoivanteddy/"
          }
        },
        {
          "cursor": "Y3Vyc29yOjEw",
          "node": {
            "name": "chequesoftware",
            "createdAt": "2018-01-02T08:58:44Z",
            "updatedAt": "2018-01-02T09:01:58Z",
            "login": "chequesoftware",
            "bio": "Looking for LOW PRICE ChequeWrtiePro Software in #MALAYSIA\r\nhttp://denariusoft.com/contactus.html\r\nFill the Form by Clicking the above LINK & Register your Deta",
            "location": "Malaysia",
            "email": "",
            "company": "Chequesoftware",
            "avatarUrl": "https://avatars2.githubusercontent.com/u/35026306?v=4",
            "websiteUrl": "http://www.denariusoft.com/chequewritepro/"
          }
        }
      ]
    }
  }
}
```

## Sample Search Repo Query

```bash
query {
  search(query: "user:alextanhongpin created:2018-01-02..2018-02-05", type: REPOSITORY, last: 10 ) {
    repositoryCount,
    pageInfo {
      hasNextPage,
      startCursor,
      endCursor,
      hasPreviousPage,
    },
    edges {
      cursor,
      node {
        ...on Repository {
          name,
          createdAt,
          updatedAt,
          description
        }
      }
    }
  }
}
```

Output:

```bash
{
  "data": {
    "search": {
      "repositoryCount": 14,
      "pageInfo": {
        "hasNextPage": true,
        "startCursor": "Y3Vyc29yOjE=",
        "endCursor": "Y3Vyc29yOjEw",
        "hasPreviousPage": false
      },
      "edges": [
        {
          "cursor": "Y3Vyc29yOjE=",
          "node": {
            "name": "stable-marriage-problem",
            "createdAt": "2018-01-24T15:51:59Z",
            "updatedAt": "2018-01-29T14:42:36Z",
            "description": "Solving the Stable Marriage/Matching Problem with the Gale–Shapley algorithm"
          }
        },
        {
          "cursor": "Y3Vyc29yOjI=",
          "node": {
            "name": "machine-learning-in-action",
            "createdAt": "2018-01-19T14:01:41Z",
            "updatedAt": "2018-01-19T14:02:13Z",
            "description": "Tutorials from the book"
          }
        },
        {
          "cursor": "Y3Vyc29yOjM=",
          "node": {
            "name": "trie",
            "createdAt": "2018-01-14T16:30:16Z",
            "updatedAt": "2018-01-14T16:30:32Z",
            "description": "Sample code on Trie data structure"
          }
        },
        {
          "cursor": "Y3Vyc29yOjQ=",
          "node": {
            "name": "go-twirp",
            "createdAt": "2018-01-22T10:26:28Z",
            "updatedAt": "2018-01-22T10:40:23Z",
            "description": "Go Twirp"
          }
        },
        {
          "cursor": "Y3Vyc29yOjU=",
          "node": {
            "name": "elixir-basic",
            "createdAt": "2018-02-02T07:09:13Z",
            "updatedAt": "2018-02-06T05:32:26Z",
            "description": "Learning the basics of Elixir from Elixir School"
          }
        },
        {
          "cursor": "Y3Vyc29yOjY=",
          "node": {
            "name": "dotnetcore-microservice",
            "createdAt": "2018-02-01T16:07:51Z",
            "updatedAt": "2018-02-18T04:48:51Z",
            "description": "Sample microservice for .NET Core 2.0"
          }
        },
        {
          "cursor": "Y3Vyc29yOjc=",
          "node": {
            "name": "json-schema-doca",
            "createdAt": "2018-01-23T03:49:24Z",
            "updatedAt": "2018-01-23T03:49:24Z",
            "description": "Sample guide on generating JSON-Schema README documentation using Cloudflare's Doca"
          }
        },
        {
          "cursor": "Y3Vyc29yOjg=",
          "node": {
            "name": "frecency",
            "createdAt": "2018-01-02T05:51:36Z",
            "updatedAt": "2018-01-02T05:51:36Z",
            "description": "Implementation of Firefox's frecency algorithm - which is the combination of frequency and recency"
          }
        },
        {
          "cursor": "Y3Vyc29yOjk=",
          "node": {
            "name": "docker-monitoring",
            "createdAt": "2018-01-30T08:52:20Z",
            "updatedAt": "2018-01-30T08:52:20Z",
            "description": "Sample Docker monitoring with Prometheus"
          }
        },
        {
          "cursor": "Y3Vyc29yOjEw",
          "node": {
            "name": "react-redux-boilerplate",
            "createdAt": "2018-01-22T08:42:59Z",
            "updatedAt": "2018-04-24T17:21:05Z",
            "description": "Sample boilerplate for React-Redux"
          }
        }
      ]
    }
  }
}
```


## Golang mgo

```go
  // Set index
  err := users.EnsureIndex(mgo.Index{
		Key:    []string{"login"},
		Unique: true,
	})
  // 
	if err != nil {
		panic(err)
	}

  // Bulk Upsert - it will insert the document if it doesn't exist yet - 
  // at the same time, it will update the existing documents.
  // The field `login` is unique
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

  // Single Upsert 
	change, err := users.Upsert(
		bson.M{"login": "alextanhongpin"},
		bson.M{"$set": bson.M{
			"timestamp": time.Now(),
			"count":     10,
		}},
	)
	if err != nil {
		panic(err)
	}
	log.Println("change", change)

  // Find  One
	var user User
	if err = users.Find(bson.M{"login": "alextanhongpin"}).
		Select(bson.M{"login": 0}).
		One(&user); err != nil {
		log.Println(err)
	}
	log.Printf("user: %+v\n", user)

  // Find All
	var userColl []User
	if err = users.Find(bson.M{}).All(&userColl); err != nil {
		log.Println(err)
	}
	log.Printf("user: %#v\n count: %d", userColl, len(userColl))

  // Remove All
	change, err := users.RemoveAll(bson.M{})
	if err != nil {
		panic(err)
	}
	log.Println("change", change)
```