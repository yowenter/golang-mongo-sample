FROM google/golang
MAINTAINER Golfen Guo "golfen.guo@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get github.com/DaoCloud/golang-mongo-sample
RUN install github.com/DaoCloud/golang-mongo-sample

# nginx will bind to port 80, exposing API
EXPOSE 80
CMD ["/gopath/app/bin/golang-mongo-sample"]