
FROM golang

# Install VCS tools to support "go get" commands and install gcc.
RUN apt-get update -qqy && apt-get install -qqy --no-install-recommends git mercurial subversion gcc protobuf-compiler
RUN GIT_TAG="v1.2.0"  && go get -d -u github.com/golang/protobuf/protoc-gen-go && \
    git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG && \
    go install github.com/golang/protobuf/protoc-gen-go
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.15.0
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

RUN GRPC_GATEWAY_VERSION=v1.8.5 && \
    wget -qO/bin/protoc-gen-grpc-gateway https://github.com/grpc-ecosystem/grpc-gateway/releases/download/${GRPC_GATEWAY_VERSION}/protoc-gen-grpc-gateway-${GRPC_GATEWAY_VERSION}-linux-x86_64 && \
    chmod a+x /bin/protoc-gen-grpc-gateway && \
    wget -qO/bin/protoc-gen-swagger https://github.com/grpc-ecosystem/grpc-gateway/releases/download/${GRPC_GATEWAY_VERSION}/protoc-gen-swagger-${GRPC_GATEWAY_VERSION}-linux-x86_64 && \
    chmod a+x /bin/protoc-gen-swagger

# We blank out the GOPATH because the base image sets it, and
# if the user of this build step does *not* set it, we want to
# explore other options for workspace derivation.
ENV GOPATH=

RUN mkdir /builder

COPY go_workspace.go prepare_workspace.inc /builder/

COPY go.bash /builder/bin/
ENV PATH=/builder/bin:$PATH

RUN go build -o /builder/go_workspace /builder/go_workspace.go

ENTRYPOINT ["go.bash"]
