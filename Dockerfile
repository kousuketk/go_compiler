FROM ubuntu:latest
WORKDIR /home/projects
# ENV GOVERSION 1.16.3
# ENV ARCH amd64
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
  vim \
  gcc \
  g++ \
  golang
  # && curl -s -o /tmp/go.tar.gz https://storage.googleapis.com/golang/go$GOVERSION.linux-$ARCH.tar.gz \
  # && tar -C /usr/local -xzf /tmp/go.tar.gz \
  # && rm /tmp/go.tar.gz