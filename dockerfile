FROM golang:1.14-alpine

WORKDIR /go/src/app
COPY ./src .

RUN apk add build-base
RUN go get -d -t ./...
RUN go install ./...
RUN go test .

CMD ["app"] 
