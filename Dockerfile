FROM golang:1.15.3-alpine
WORKDIR /tick
ADD . /tick
RUN cd /tick && go build
ENTRYPOINT ./tick