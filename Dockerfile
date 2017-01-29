FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/goseed /opt/bin/goseed

CMD ["/opt/bin/goseed"]