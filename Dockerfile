FROM golang
ENV SRC github.com/chenpengcheng/boilerplates
ARG VERSION
COPY . /go/src/${SRC}

RUN go install -ldflags="-X main.version=$VERSION" \
  ${SRC}/cmd/service

ENTRYPOINT [ "/go/bin/service", "--addr", "--db-addr", "--debug" ]
