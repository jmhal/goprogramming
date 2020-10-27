# Simple Go Form
FROM ubuntu
MAINTAINER jmarcelo.alencar@gmail.com

RUN apt-get update
RUN apt-get install apt-transport-https ca-certificates -y
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends git golang
RUN git clone https://github.com/jmhal/goprogramming
WORKDIR goprogramming/web
EXPOSE 8080
CMD ["go", "run", "src/forms/forms.go"]