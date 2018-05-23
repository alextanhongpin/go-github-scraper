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