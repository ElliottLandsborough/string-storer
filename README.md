# string-storer

## Live demo

http://string-storer-lb-tf-663390881.eu-west-2.elb.amazonaws.com/

## Run with docker

Some changes needed in web/default.conf for it to work locally

```bash
docker-compose up -d
```

## Run manually

```bash
cd api
go install
go run main.go
```

```bash
curl -X POST http://0.0.0.0:8080/update -H 'Content-Type: application/json' -d '{"title":"a new title"}'
```
