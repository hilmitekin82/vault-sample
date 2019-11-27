# # First stage: start with a Golang base image
FROM golang

WORKDIR /go/src/hello

COPY . .

RUN CGO_ENABLED=0 go get -v ./...

FROM docker.io/hilmit82/envconsul-0.9.0:scratch

# Copy the binary from the first stage
COPY --from=0 /go/bin/hello /

# Tell Docker what executable to run by default when starting this container
ENTRYPOINT ["/envconsul","-config=/etc/envconsul/envconsul-config.hcl","/hello"]