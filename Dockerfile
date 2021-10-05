FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

RUN go get github.com/ThomasITU/assignment01
RUN cd /build/ && git clone https://github.com/ThomasITU/assignment01.git
RUN cd /build/assignment01/server && go build ./...


EXPOSE 8080

ENTRYPOINT [ "/build/assignment01/server/server" ]

