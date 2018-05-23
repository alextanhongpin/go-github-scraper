## API Calls

```bash
$ curl -H "Authorization: bearer token" -X POST -d " \
 { \
   \"query\": \"query { viewer { login }}\" \
 } \
" https://api.github.com/graphql
```

## Sample Search Query

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
