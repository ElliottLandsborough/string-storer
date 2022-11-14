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
