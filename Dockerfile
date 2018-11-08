FROM golang:1.11
LABEL maintainer="@punguin-in-the-sky"

RUN apt-get update -qq && apt-get install -y vim

ENV APP_NAME /myapp
RUN mkdir $GOPATH/src/${APP_NAME}
WORKDIR $GOPATH/src/${APP_NAME}
COPY .${APP_NAME} .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go get -u gopkg.in/go-playground/validator.v9