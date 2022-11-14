FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go install
CMD ["go", "run", "main.go"]
