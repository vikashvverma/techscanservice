FROM alpine:latest

MAINTAINER Vikash <vikash@programminggeek.in>

WORKDIR "/opt"

ADD .docker_build/techscanservice /opt/bin/techscanservice
ADD ./config /opt/config

CMD ["/opt/bin/techscanservice"]

