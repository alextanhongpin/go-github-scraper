# go-github-scraper

A Github Scraper and recommendation engine written in `Go`, to replace the previous version written in NodeJS (TypeScript).


## Development

Create a `.env` that contains minimum the following environment variables:

```.env
GITHUB_TOKEN=<your_github_token>
DB_NAME=<mgo_db_name>
DB_HOST=<mgo_db_host>
DB_USER=<mgo_db_user>
DB_AUTH=<mgo_db_auth>
DB_PASS=<mgo_db_pass>
```

## Start

```bash
$ make start
```

## Build Docker Image

```bash
$ make docker
```

## CPU Profiling

Or to look at a 30-second CPU profile:

```bash
$ go tool pprof  http://localhost:6060:/debug/pprof/profile

$ go tool pprof /Users/alextanhongpin/pprof/pprof.samples.cpu.003.pb.gz

$ (pprof) top10
$ (pprof) top5 -cum
```


## Memory Profiling

Then use the pprof tool to look at the heap profile:

```bash
$ go tool pprof http://localhost:6060/debug/pprof/heap

(pprof) top5
(pprof) list FnName
```




One option is ‘–alloc_space’ which tells you how many megabytes have been allocated.

```bash
$ go tool pprof --alloc_space http://localhost:6060/debug/pprof/heap
```

The other – ‘–inuse_space’ tells you know how many are still in use.

```
<!-- $ go tool pprof --inuse_objects http://localhost:6060/debug/pprof/heap -->
$ go tool pprof --inuse_space http://localhost:6060/debug/pprof/heap
```