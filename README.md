# string-storer

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

http://0.0.0.0:8080/posts


### Save new string

```bash
curl -X POST http://0.0.0.0:8080/posts -H 'Content-Type: application/json' -d '{"title":"a new title"}'
```

## What is wrong with this solution?

 - It doesn't have a shared state - e.g a database, mongo/dynamo/mysql. So if I scale it up the dynamic string won't be consistent between http requests.
 - I think the golang garbage collection would mean this won't balloon in ram - if I made it so that it saved history of posted items It would fill up ram. Solvable with above point.
 - There is no monitoring at all. I wouldn't actually know if the containers crashed
 - The JS is messy, could be minified and could be using a framework like react.
 - Project structure is not official golang best practice. Work could be done here but there is no point if this microservice is staying as-is
 - Not much actual structure in the golang - could probably be separated out into some sort of DDD so that it can be expanded later
 - No CI or unit tests. Should definitely add an integration test for each endpoint at minimum. Test all golang functions if possible, not that hard to do right now
 - Could even do integration tests on the frontend. Not sure if necesary
 - Nginx configs need to be changed between dev/prod when using docker. Maybe have two separate docker-compose files or some flags at least so that devs don't have to deal with this
 - I am not sure if this is best terraform practice. I don't have loads of experience with it
 - I wanted to create a makefile for all of the provisioning and docker but didn't have time