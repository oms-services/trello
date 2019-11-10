FROM golang

RUN go get github.com/gorilla/mux
RUN go get github.com/adlio/trello
RUN go get github.com/cloudevents/sdk-go

WORKDIR /go/src/github.com/oms-services/trello

ADD . /go/src/github.com/oms-services/trello

RUN go install github.com/oms-services/trello

ENTRYPOINT trello

EXPOSE 3000