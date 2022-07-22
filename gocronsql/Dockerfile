FROM golang:1.16-alpine

WORKDIR /code

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy
RUN go get github.com/go-co-op/gocron
RUN go get github.com/go-sql-driver/mysql

RUN go mod download

COPY *.go ./

RUN go build -o /main


CMD [ "/main" ]