From golang:1.12

RUN apt-get update || exit 0
RUN apt-get upgrade -y
RUN apt-get install vim sudo dbus curl bc supervisor -y

## setup gdb
RUN apt-get install gdb -y
RUN go get github.com/derekparker/delve/cmd/dlv

#setup infra repo
#RUN sh -c "$(wget http://repo.carousell.io/setup.sh -O -)"
RUN echo "2018-08-01"
RUN apt-get update || exit 0
#RUN apt-get install config-service log-service -y

#run pre code script which are cached accross runs
COPY docker/gdbinit /root/.gdbinit
COPY docker/scripts/pre_code.sh /tmp/
RUN bash /tmp/pre_code.sh

RUN mkdir -p /go/src/github.com/im-adarsh/text-resource
RUN mkdir -p /opt/config/
COPY . /go/src/github.com/im-adarsh/text-resource
RUN go install github.com/im-adarsh/text-resource/text-resource/cmd/server

#COPY docker/scripts/post_code.sh /tmp/
#RUN bash  /tmp/post_code.sh

EXPOSE 9281 9282 9283 9284