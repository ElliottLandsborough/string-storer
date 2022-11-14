# login to ecr
aws ecr get-login-password --region eu-west-2 | docker login --username AWS --password-stdin 633322954385.dkr.ecr.eu-west-2.amazonaws.com
# build and push api
docker build -t string-storer-api ./api
docker tag string-storer-api:latest 633322954385.dkr.ecr.eu-west-2.amazonaws.com/string-storer:latest
docker push 633322954385.dkr.ecr.eu-west-2.amazonaws.com/string-storer:latest
# build and push web
docker build -t string-storer-web ./web
docker tag string-storer-web:latest 633322954385.dkr.ecr.eu-west-2.amazonaws.com/string-storer:latest
docker push 633322954385.dkr.ecr.eu-west-2.amazonaws.com/string-storer:latest