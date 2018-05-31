<!-- ## API Calls

```.env
GO_SCRAPER_CRONTAB_USER_TAB="*/20 * * * * *"
GO_SCRAPER_CRONTAB_USER_ENABLE=true
GO_SCRAPER_CRONTAB_REPO_ENABLE=true
GO_SCRAPER_CRONTAB_STAT_ENABLE=true
GO_SCRAPER_CRONTAB_PROFILE_ENABLE=true
GO_SCRAPER_CRONTAB_MATCH_ENABLE=true
GO_SCRAPER_CRONTAB_USER_TRIGGER=false
GO_SCRAPER_CRONTAB_REPO_TRIGGER=false
GO_SCRAPER_CRONTAB_STAT_TRIGGER=false
GO_SCRAPER_CRONTAB_PROFILE_TRIGGER=false
GO_SCRAPER_CRONTAB_MATCH_TRIGGER=false
GO_SCRAPER_DB_USER=${MONGO_USER}
GO_SCRAPER_DB_PASS=${MONGO_PASS}
GO_SCRAPER_DB_NAME="scraper"
GO_SCRAPER_DB_AUTH="admin"
GO_SCRAPER_DB_HOST=${MONGO_HOST}
GO_SCRAPER_GITHUB_TOKEN=${GITHUB_TOKEN}
```

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
  search(query: "location:malaysia created:2018-01-01..2018-02-01", type: USER, last: 10, after: ) {
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
          websiteUrl,
          repositories(last: 0) {
            totalCount
          },
          gists(last: 0) {
            totalCount
          },
          followers(last: 0) {
            totalCount
          },
          following(last: 0) {
            totalCount
          },
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
      "userCount": 77,
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
            "name": "Dr.Saif Yousif",
            "createdAt": "2018-01-29T08:41:08Z",
            "updatedAt": "2018-04-01T13:46:19Z",
            "login": "drsaifyousif",
            "bio": "MDSC, Ph.D, CIW, PMP, CEH",
            "location": "Malaysia",
            "email": "",
            "company": "ARID",
            "avatarUrl": "https://avatars0.githubusercontent.com/u/35916074?v=4",
            "websiteUrl": "arid.my/0001-0001",
            "repositories": {
              "totalCount": 1
            },
            "gists": {
              "totalCount": 0
            },
            "followers": {
              "totalCount": 4
            },
            "following": {
              "totalCount": 0
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjI=",
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
            "websiteUrl": "philip-khor.github.io",
            "repositories": {
              "totalCount": 4
            },
            "gists": {
              "totalCount": 0
            },
            "followers": {
              "totalCount": 2
            },
            "following": {
              "totalCount": 3
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjM=",
          "node": {
            "name": "Shouquat Mahmoud Mazumder",
            "createdAt": "2018-01-27T17:01:41Z",
            "updatedAt": "2018-05-23T04:15:51Z",
            "login": "Shouquat",
            "bio": "",
            "location": "Malaysia",
            "email": "shouquat4281@gmail.com",
            "company": "",
            "avatarUrl": "https://avatars3.githubusercontent.com/u/35873804?v=4",
            "websiteUrl": "CRYPTO LOVER",
            "repositories": {
              "totalCount": 0
            },
            "gists": {
              "totalCount": 0
            },
            "followers": {
              "totalCount": 1
            },
            "following": {
              "totalCount": 5
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjQ=",
          "node": {
            "name": "",
            "createdAt": "2018-01-16T11:25:25Z",
            "updatedAt": "2018-01-16T11:32:15Z",
            "login": "ibc003",
            "bio": "IBC003 is a online casino situated in the Malaysia Which provide the Malaysia sports betting ,online betting ,and many other games related the casino . ",
            "location": "Kuala Lumpur ,Malaysia",
            "email": "",
            "company": "IBC003",
            "avatarUrl": "https://avatars3.githubusercontent.com/u/35489758?v=4",
            "websiteUrl": "https://ibc003.net/",
            "repositories": {
              "totalCount": 0
            },
            "gists": {
              "totalCount": 0
            },
            "followers": {
              "totalCount": 1
            },
            "following": {
              "totalCount": 0
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjU=",
          "node": {
            "name": "Suhail",
            "createdAt": "2018-01-28T04:02:27Z",
            "updatedAt": "2018-04-04T09:50:36Z",
            "login": "ryeshl",
            "bio": "Newbie Data Scientist",
            "location": "Kuala Lumpur, Malaysia",
            "email": "",
            "company": "",
            "avatarUrl": "https://avatars1.githubusercontent.com/u/35884186?v=4",
            "websiteUrl": "",
            "repositories": {
              "totalCount": 2
            },
            "gists": {
              "totalCount": 1
            },
            "followers": {
              "totalCount": 1
            },
            "following": {
              "totalCount": 4
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjY=",
          "node": {
            "name": "Adrian Yeong",
            "createdAt": "2018-01-25T03:18:44Z",
            "updatedAt": "2018-04-27T02:01:11Z",
            "login": "ayeong-SSI",
            "bio": "",
            "location": "Kuala Lumpur, Malaysia",
            "email": "ayeong@servicesource.com",
            "company": "Service Source",
            "avatarUrl": "https://avatars1.githubusercontent.com/u/35789933?v=4",
            "websiteUrl": "",
            "repositories": {
              "totalCount": 0
            },
            "gists": {
              "totalCount": 0
            },
            "followers": {
              "totalCount": 0
            },
            "following": {
              "totalCount": 0
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjc=",
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
            "websiteUrl": "",
            "repositories": {
              "totalCount": 1
            },
            "gists": {
              "totalCount": 0
            },
            "followers": {
              "totalCount": 0
            },
            "following": {
              "totalCount": 1
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjg=",
          "node": {}
        },
        {
          "cursor": "Y3Vyc29yOjk=",
          "node": {}
        },
        {
          "cursor": "Y3Vyc29yOjEw",
          "node": {
            "name": "ZULFAHMYRIZAL",
            "createdAt": "2018-01-18T07:05:38Z",
            "updatedAt": "2018-01-18T07:26:16Z",
            "login": "zulfahmyrizal",
            "bio": "I am a PHP developer and started learning it from 2017 and continue to learn more in PHP. My aim to become finest in PHP and always ready to learn new things .",
            "location": "MALAYSIA",
            "email": "",
            "company": "MYADSFX",
            "avatarUrl": "https://avatars0.githubusercontent.com/u/35556796?v=4",
            "websiteUrl": "http://www.myadsfx.com",
            "repositories": {
              "totalCount": 1
            },
            "gists": {
              "totalCount": 0
            },
            "followers": {
              "totalCount": 0
            },
            "following": {
              "totalCount": 0
            }
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
          description,
          languages (last: 30) {
            totalCount,
            edges {
              node {
                name,
                color
              }
            }
          },
          homepageUrl,
          forkCount,
          isFork,
          nameWithOwner,
          owner {
            login,
            avatarUrl
          },
          stargazers (last: 0) {
						totalCount
          },
          watchers (last: 0) {
            totalCount
          },
          url
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
            "description": "Solving the Stable Marriage/Matching Problem with the Gale–Shapley algorithm",
            "languages": {
              "totalCount": 8,
              "edges": [
                {
                  "node": {
                    "name": "Jupyter Notebook",
                    "color": "#DA5B0B"
                  }
                },
                {
                  "node": {
                    "name": "Haskell",
                    "color": "#5e5086"
                  }
                },
                {
                  "node": {
                    "name": "JavaScript",
                    "color": "#f1e05a"
                  }
                },
                {
                  "node": {
                    "name": "Scala",
                    "color": "#c22d40"
                  }
                },
                {
                  "node": {
                    "name": "Go",
                    "color": "#375eab"
                  }
                },
                {
                  "node": {
                    "name": "Rust",
                    "color": "#dea584"
                  }
                },
                {
                  "node": {
                    "name": "C#",
                    "color": "#178600"
                  }
                },
                {
                  "node": {
                    "name": "Elixir",
                    "color": "#6e4a7e"
                  }
                }
              ]
            },
            "homepageUrl": null,
            "forkCount": 0,
            "isFork": false,
            "nameWithOwner": "alextanhongpin/stable-marriage-problem",
            "owner": {
              "login": "alextanhongpin",
              "avatarUrl": "https://avatars3.githubusercontent.com/u/6033638?v=4"
            },
            "stargazers": {
              "totalCount": 1
            },
            "watchers": {
              "totalCount": 1
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjI=",
          "node": {
            "name": "machine-learning-in-action",
            "createdAt": "2018-01-19T14:01:41Z",
            "updatedAt": "2018-01-19T14:02:13Z",
            "description": "Tutorials from the book",
            "languages": {
              "totalCount": 2,
              "edges": [
                {
                  "node": {
                    "name": "Jupyter Notebook",
                    "color": "#DA5B0B"
                  }
                },
                {
                  "node": {
                    "name": "Python",
                    "color": "#3572A5"
                  }
                }
              ]
            },
            "homepageUrl": null,
            "forkCount": 0,
            "isFork": false,
            "nameWithOwner": "alextanhongpin/machine-learning-in-action",
            "owner": {
              "login": "alextanhongpin",
              "avatarUrl": "https://avatars3.githubusercontent.com/u/6033638?v=4"
            },
            "stargazers": {
              "totalCount": 0
            },
            "watchers": {
              "totalCount": 1
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjM=",
          "node": {
            "name": "trie",
            "createdAt": "2018-01-14T16:30:16Z",
            "updatedAt": "2018-01-14T16:30:32Z",
            "description": "Sample code on Trie data structure",
            "languages": {
              "totalCount": 1,
              "edges": [
                {
                  "node": {
                    "name": "Jupyter Notebook",
                    "color": "#DA5B0B"
                  }
                }
              ]
            },
            "homepageUrl": null,
            "forkCount": 0,
            "isFork": false,
            "nameWithOwner": "alextanhongpin/trie",
            "owner": {
              "login": "alextanhongpin",
              "avatarUrl": "https://avatars3.githubusercontent.com/u/6033638?v=4"
            },
            "stargazers": {
              "totalCount": 0
            },
            "watchers": {
              "totalCount": 1
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjQ=",
          "node": {
            "name": "go-twirp",
            "createdAt": "2018-01-22T10:26:28Z",
            "updatedAt": "2018-01-22T10:40:23Z",
            "description": "Go Twirp",
            "languages": {
              "totalCount": 1,
              "edges": [
                {
                  "node": {
                    "name": "Go",
                    "color": "#375eab"
                  }
                }
              ]
            },
            "homepageUrl": null,
            "forkCount": 0,
            "isFork": false,
            "nameWithOwner": "alextanhongpin/go-twirp",
            "owner": {
              "login": "alextanhongpin",
              "avatarUrl": "https://avatars3.githubusercontent.com/u/6033638?v=4"
            },
            "stargazers": {
              "totalCount": 0
            },
            "watchers": {
              "totalCount": 1
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjU=",
          "node": {
            "name": "elixir-basic",
            "createdAt": "2018-02-02T07:09:13Z",
            "updatedAt": "2018-02-06T05:32:26Z",
            "description": "Learning the basics of Elixir from Elixir School",
            "languages": {
              "totalCount": 1,
              "edges": [
                {
                  "node": {
                    "name": "Elixir",
                    "color": "#6e4a7e"
                  }
                }
              ]
            },
            "homepageUrl": null,
            "forkCount": 0,
            "isFork": false,
            "nameWithOwner": "alextanhongpin/elixir-basic",
            "owner": {
              "login": "alextanhongpin",
              "avatarUrl": "https://avatars3.githubusercontent.com/u/6033638?v=4"
            },
            "stargazers": {
              "totalCount": 0
            },
            "watchers": {
              "totalCount": 1
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjY=",
          "node": {
            "name": "dotnetcore-microservice",
            "createdAt": "2018-02-01T16:07:51Z",
            "updatedAt": "2018-02-18T04:48:51Z",
            "description": "Sample microservice for .NET Core 2.0",
            "languages": {
              "totalCount": 3,
              "edges": [
                {
                  "node": {
                    "name": "C#",
                    "color": "#178600"
                  }
                },
                {
                  "node": {
                    "name": "CSS",
                    "color": "#563d7c"
                  }
                },
                {
                  "node": {
                    "name": "JavaScript",
                    "color": "#f1e05a"
                  }
                }
              ]
            },
            "homepageUrl": null,
            "forkCount": 0,
            "isFork": false,
            "nameWithOwner": "alextanhongpin/dotnetcore-microservice",
            "owner": {
              "login": "alextanhongpin",
              "avatarUrl": "https://avatars3.githubusercontent.com/u/6033638?v=4"
            },
            "stargazers": {
              "totalCount": 0
            },
            "watchers": {
              "totalCount": 1
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjc=",
          "node": {
            "name": "json-schema-doca",
            "createdAt": "2018-01-23T03:49:24Z",
            "updatedAt": "2018-01-23T03:49:24Z",
            "description": "Sample guide on generating JSON-Schema README documentation using Cloudflare's Doca",
            "languages": {
              "totalCount": 0,
              "edges": []
            },
            "homepageUrl": null,
            "forkCount": 0,
            "isFork": false,
            "nameWithOwner": "alextanhongpin/json-schema-doca",
            "owner": {
              "login": "alextanhongpin",
              "avatarUrl": "https://avatars3.githubusercontent.com/u/6033638?v=4"
            },
            "stargazers": {
              "totalCount": 0
            },
            "watchers": {
              "totalCount": 1
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjg=",
          "node": {
            "name": "frecency",
            "createdAt": "2018-01-02T05:51:36Z",
            "updatedAt": "2018-01-02T05:51:36Z",
            "description": "Implementation of Firefox's frecency algorithm - which is the combination of frequency and recency",
            "languages": {
              "totalCount": 0,
              "edges": []
            },
            "homepageUrl": null,
            "forkCount": 0,
            "isFork": false,
            "nameWithOwner": "alextanhongpin/frecency",
            "owner": {
              "login": "alextanhongpin",
              "avatarUrl": "https://avatars3.githubusercontent.com/u/6033638?v=4"
            },
            "stargazers": {
              "totalCount": 0
            },
            "watchers": {
              "totalCount": 1
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjk=",
          "node": {
            "name": "docker-monitoring",
            "createdAt": "2018-01-30T08:52:20Z",
            "updatedAt": "2018-01-30T08:52:20Z",
            "description": "Sample Docker monitoring with Prometheus",
            "languages": {
              "totalCount": 0,
              "edges": []
            },
            "homepageUrl": null,
            "forkCount": 0,
            "isFork": false,
            "nameWithOwner": "alextanhongpin/docker-monitoring",
            "owner": {
              "login": "alextanhongpin",
              "avatarUrl": "https://avatars3.githubusercontent.com/u/6033638?v=4"
            },
            "stargazers": {
              "totalCount": 0
            },
            "watchers": {
              "totalCount": 1
            }
          }
        },
        {
          "cursor": "Y3Vyc29yOjEw",
          "node": {
            "name": "react-redux-boilerplate",
            "createdAt": "2018-01-22T08:42:59Z",
            "updatedAt": "2018-04-24T17:21:05Z",
            "description": "Sample boilerplate for React-Redux",
            "languages": {
              "totalCount": 3,
              "edges": [
                {
                  "node": {
                    "name": "HTML",
                    "color": "#e34c26"
                  }
                },
                {
                  "node": {
                    "name": "CSS",
                    "color": "#563d7c"
                  }
                },
                {
                  "node": {
                    "name": "JavaScript",
                    "color": "#f1e05a"
                  }
                }
              ]
            },
            "homepageUrl": null,
            "forkCount": 0,
            "isFork": false,
            "nameWithOwner": "alextanhongpin/react-redux-boilerplate",
            "owner": {
              "login": "alextanhongpin",
              "avatarUrl": "https://avatars3.githubusercontent.com/u/6033638?v=4"
            },
            "stargazers": {
              "totalCount": 0
            },
            "watchers": {
              "totalCount": 1
            }
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
``` -->

## Golang Enum String 

```golang
var enums = [...]string{
	"user_count",
	"repo_count",
}

const (
	// EnumUserCount represents the type for the user_count
	EnumUserCount Enum = iota
	// EnumRepoCount represents the type for the repo_count
	EnumRepoCount
)

func (e Enum) String() string {
	return enums[e]
}

```

## Profiling

```bash
$ go tool pprof -alloc_space -svg http://localhost:8080/debug/pprof/heap > heap.svg

$ go tool pprof -png http://localhost:8080/debug/pprof/heap > out.png
```

## Logging

- setting up a logger at init is a no-op (also look at singleton pattern when doing lazy initialization, since lazy initialization is not concurrent-safe)
- setting logger in context is a bad design pattern - instead, pass the value you want to have in your logger through context
- logging in concurrent can be tricky - use `requestId` to bind the logs together
- log the start of a service at `model` layer, but don't log the errors there
- do not centralize error logging (e.g. placing them in the model), rather handle each error logs separately for better customization
- instead of passing logger around (preferred), you can replace the global logger (uber's zap) and reuse them - this is slightly preferable for me since I don't have to inject the logger into the dependencies. Imagine having to coordinate multiple service orchestration and passing the `requestId` becomes cumbersome (pass it through context, please), passing a pre-configured logger does not work well.
- logging lifecycle? start of event, success event, error event
- standardize the format of your logger
```
updating users
error updating user
updated user

or

update user, start=true
update user, error=true
update user, success=true
```

## context

- Pass context as the first argument of your `service` methods (not in `store`'s repository pattern). Period.
- The context utility takes a context, injects the context with request context, and returns the injected context.
```
// ContextKey represents the type of the context key
type ContextKey string

// RequestID represents the request id that is passed in the context
const RequestID = ContextKey("RequestID")

// WrapContextWithRequestID wraps the existing context with the requestID field before returning them
func WrapContextWithRequestID(ctx context.Context) context.Context {
	if v := ctx.Value(RequestID); v != nil {
		return ctx
	}

	reqID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return context.WithValue(ctx, RequestID, reqID.String())
}
```
- For the above function, why not just create a function that returns a new context with the request id? This is not ideal when you have to accept a context that has other values prepopulated.

## Simplifying list of functions operations

```go
var wg sync.WaitGroup
wg.Add(8)
go func() {
  msvc.UpdateUserCount()
  wg.Done()
}()
go func() {
  msvc.UpdateRepoCount()
  wg.Done()
}()
go func() {
  msvc.UpdateReposMostRecent(20)
  wg.Done()
}()
go func() {
  msvc.UpdateRepoCountByUser(20)
  wg.Done()
}()
go func() {
  msvc.UpdateReposMostStars(20)
  wg.Done()
}()
go func() {
  msvc.UpdateLanguagesMostPopular(20)
  wg.Done()
}()
go func() {
  msvc.UpdateMostRecentReposByLanguage(20)
  wg.Done()
}()
go func() {
  msvc.UpdateReposByLanguage(20)
  wg.Done()
}()
wg.Wait()
return nil
```
```go
nullFns := []null.Fn{
  null.Fn(func() error { return msvc.UpdateUserCount(ctx) }),
  null.Fn(func() error { return msvc.UpdateRepoCount(ctx) }),
  null.Fn(func() error { return msvc.UpdateReposMostRecent(ctx, 20) }),
  null.Fn(func() error { return msvc.UpdateRepoCountByUser(ctx, 20) }),
  null.Fn(func() error { return msvc.UpdateReposMostStars(ctx, 20) }),
  null.Fn(func() error { return msvc.UpdateLanguagesMostPopular(ctx, 20) }),
  null.Fn(func() error { return msvc.UpdateMostRecentReposByLanguage(ctx, 20) }),
  null.Fn(func() error { return msvc.UpdateReposByLanguage(ctx, 20) }),
}
var wg sync.WaitGroup
wg.Add(len(nullFns))

for _, fn := range nullFns {
  go func(f null.Fn) {
    defer wg.Done()
    f()
  }(fn)
}

wg.Wait()
```

## Http Server Cancellation

In your `server.go`:

```go
	r.GET("/long", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()
		start := time.Now()

		select {
		case <-ctx.Done():
			stdlog.Println("cancelled after", time.Since(start))
			break
		case <-time.After(time.Second * 5):
			break
		}

		fmt.Fprintf(w, "completed after %v", time.Since(start))
	})
```

In your `client.go`:

```go
package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Response *http.Response
	Error    error
}

func main() {
	log.Println("initialized")
	start := time.Now()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*6)
	defer cancel()

	c := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080/long", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.WithContext(ctx)
	ch := make(chan Result, 1)

	go func() {
		res, err := c.Do(req)
		ch <- Result{
			Response: res,
			Error:    err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("cancelled request after %v\n", time.Since(start))
			return
		case res := <-ch:
			if res.Error != nil {
				log.Println(res.Error)
				return
			}
			defer res.Response.Body.Close()
			body, err := ioutil.ReadAll(res.Response.Body)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf(`got body "%s" after %v\n`, string(body), time.Since(start))
			return
		}
	}

	log.Println("completed after", time.Since(start))
}
```

		<!-- r.Get("/throttled", func(w http.ResponseWriter, r *http.Request) {
			select {
			case <-r.Context().Done():
				switch r.Context().Err() {
				case context.DeadlineExceeded:
					w.WriteHeader(504)
					w.Write([]byte("Processing too slow\n"))
				default:
					w.Write([]byte("Canceled\n"))
				}
				return

			case <-time.After(5 * time.Second):
				// The above channel simulates some hard work.
			}

			w.Write([]byte("Processed\n"))
		}) -->

## Useful tools

```bash
$ go get github.com/3rf/codecoroner
$ codecoroner funcs ./...
$ codecoroner idents ./...
```

