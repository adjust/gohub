FROM golang:1.4.2-wheezy

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        curl \
    && rm -rf /var/lib/apt/lists/*

RUN curl -s https://get.docker.io/ubuntu/ | sh && \
    echo 'DOCKER_OPTS="-H :2375 -H unix:///var/run/docker.sock"' >> /etc/default/docker

ADD . /go/src/github.com/adjust/gohub

WORKDIR /go/src/github.com/adjust/gohub

RUN go install github.com/adjust/gohub

EXPOSE 6578

ENTRYPOINT ["/go/bin/gohub"]

CMD ["--port", "6578", "--log", "/var/log/webhook.log"]
