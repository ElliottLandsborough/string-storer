# String Storer

[![CircleCI](https://circleci.com/gh/ElliottLandsborough/string-storer.svg?style=svg)](https://circleci.com/gh/ElliottLandsborough/string-storer)

## Live demo

http://string-storer-lb-tf-663390881.eu-west-2.elb.amazonaws.com/

## Run with docker

Some changes needed in web/default.conf for it to work locally

```bash
docker-compose up -d
```

## Run manually

### Start go server (port 8080)

```bash
cd api
go install
go run main.go
```

### View current saved string

```bash
curl http://0.0.0.0:8080/posts
```

### Save new string

```bash
curl -X POST http://0.0.0.0:8080/posts -H 'Content-Type: application/json' -d '{"title":"a new title"}'
```

### Some basic tests

```bash
cd api
go test *.go
```

## What is wrong with this solution?

 - It doesn't have a shared state - e.g. a database, mongo/dynamo/MySQL. This means that if I scale it up the dynamic string won't be consistent between http requests (right now the 'desired_count' is set to 1)
 - The above also means that if I redeploy (or if the api container crashes) it will lose its current state and revert back to the default title
 - I think the golang garbage collection would mean this won't balloon in ram - If I made it so that it saved history of posted items It would fill up ram. Solvable with above point about database/shared state
 - Use better input cleansing if storing in MySQL or similar. Right now I store the raw string in memory and then just clean it up when I output it. It might be wise to strip HTML or even limit input entirely to [A-Z0-9] and space. Unless we need multilanguage. It really depends on the use case.
 - I'm not using https. This is not a great idea. I could fix this easily with a CNAME and cloud flare for free if it needed to go live quickly.
 - There is no monitoring at all. I wouldn't know if the containers crashed. Need to add exception monitoring to the golang too.
 - The JS is messy, could be minified and could be using a framework like react.
 - Project structure is not the official golang best practice. Work could be done here but there is no point if this microservice is staying as-is
 - Not much actual structure in the golang - could probably be separated out into some sort of DDD so that it can be expanded/edited easily later
 - Minimal CI/CD and unit tests. I should add an integration test for each endpoint at minimum - and test all golang functions if possible
 - I could also do integration tests on the front end. Not sure if this is necessary
 - The Nginx configs need to be changed between dev/prod when using docker. Maybe have two separate docker-compose files or some flags at least so that developers don't have to deal with this
 - I am not sure if this is best terraform practice. I don't have loads of experience with it
 - I wanted to create a make file for all the provisioning/docker but didn't have time