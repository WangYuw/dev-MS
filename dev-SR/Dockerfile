FROM golang
LABEL maintainer="yuwei"

WORKDIR /go/src

ADD src /go/src
ADD src/main.go /go/src

RUN go build main.go

ENTRYPOINT [ "./main" ]

EXPOSE 8080