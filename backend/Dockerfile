# FROM golang:1.18.3-buster

# # COPY server /usr/local/go/src
# WORKDIR /usr/local/go/src

# # ARG GO111MODULE=on

# RUN apt-get update && apt-get install git
# run go get -u github.com/gin-contrib/cors
# run go get github.com/gin-gonic/gin
# run go get github.com/go-sql-driver/mysql
#   # && go mod init
#   # && GO111MODULE=off go get -u github.com/gin-gonic/gin \

# # CMD gin -i run

FROM golang:1.15.7-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

CMD ["go", "run", "main.go"]