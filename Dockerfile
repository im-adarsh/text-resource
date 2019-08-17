From golang:1.12

RUN apt-get update || exit 0
RUN apt-get upgrade -y

RUN mkdir -p /go/src/github.com/im-adarsh/text-resource
COPY . /go/src/github.com/im-adarsh/text-resource
RUN go install github.com/im-adarsh/text-resource/text-resource/cmd/server/
