FROM ubuntu:jammy

ARG VERSION="1.21.1" # go version
ARG ARCH="arm64" # go archicture

RUN apt-get update -y && \
    apt install curl -y 

# Get go binary
RUN curl -O -L "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"

# Extract 
RUN tar -C /usr/local -xf "go${VERSION}.linux-${ARCH}.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH