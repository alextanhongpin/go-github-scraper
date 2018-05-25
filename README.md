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
```