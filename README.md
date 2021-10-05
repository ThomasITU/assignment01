# assignment01
Check Dockerfil for steps 

remember to copy, avoid parsing go.mod module declares its path "/x/.." but was required "/x"
COPY go.mod /build
COPY go.sum /build 

go build ./... -- to build nested packages

commands
docker build -t "nameOfImage" .
docker run -p "port:port" -tid "nameOfImage"

examples
port 8080:8080
"nameOfImage" = test
