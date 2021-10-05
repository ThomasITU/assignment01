# assignment01
**Check Dockerfile for steps**

remember to copy -> avoid parsing go.mod module declares its path "/x/.." but was required "/x"


COPY go.mod /build

COPY go.sum /build 

go build ./... -- to build nested packages

**commands:**
<ol>
<li>$env:DOCKER_BUILDKIT=1 to build</li>
<li>docker build -t "nameOfImage" .</li>
<li>docker run -p "port:port" -tid "nameOfImage"</li>
</ol>

**examples**


`docker build -t test .`


`docker run -p 8080:8080 -tid test`
