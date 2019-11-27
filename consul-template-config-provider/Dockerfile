FROM golang as buildStage
WORKDIR /go/src/main
COPY . .
RUN CGO_ENABLED=0 go get -v ./...

FROM scratch
COPY --from=buildStage /go/bin/main /
ENTRYPOINT ["/main"]