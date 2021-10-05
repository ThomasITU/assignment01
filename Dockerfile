FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

RUN go get 


RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["go", "run", "main.go"]
