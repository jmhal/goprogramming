# Simple Go Notes
FROM ubuntu
MAINTAINER jmarcelo.alencar@gmail.com

RUN apt-get update
RUN apt-get install apt-transport-https ca-certificates -y
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends git golang
RUN go get go.mongodb.org/mongo-driver/bson
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get go.mongodb.org/mongo-driver/mongo/options
RUN git clone https://github.com/jmhal/goprogramming
WORKDIR goprogramming/notes
EXPOSE 8080
CMD ["go", "run", "notes.go"]