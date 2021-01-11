# STAGE 1: Build
FROM golang:1.12-alpine AS build

# Install Git
RUN apk update && apk upgrade && apk add --no-cache git 

# Get dependencies for Go part of build
RUN go get -u github.com/gorilla/mux
RUN go get -u go.mongodb.org/mongo-driver/bson
RUN go get -u go.mongodb.org/mongo-driver/bson/primitive
RUN go get -u go.mongodb.org/mongo-driver/mongo
RUN go get -u go.mongodb.org/mongo-driver/mongo/options

# Copy all sources in
COPY main.go main.go

# Do the build. Script is part of incoming sources.
RUN go build main.go
RUN cp main /

# STAGE 2: Runtime
FROM alpine

USER nobody:nobody
COPY --from=build /main /main

CMD [ "/main" ]