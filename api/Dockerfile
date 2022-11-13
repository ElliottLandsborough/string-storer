FROM golang:1.19.3

WORKDIR /go/src/app
COPY . .

RUN go install
CMD ["go", "run", "main.go"]
