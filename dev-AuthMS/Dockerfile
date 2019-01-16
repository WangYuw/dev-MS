FROM golang
LABEL maintainer="yuwei"

WORKDIR /go/src

COPY src /go/src
COPY src/main.go /go/src

RUN go build main.go

ENTRYPOINT [ "./main" ]