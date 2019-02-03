FROM golang:1.10-alpine

WORKDIR /go/src/app
COPY /src .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 9000

CMD ["app"]