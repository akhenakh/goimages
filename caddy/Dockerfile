FROM golang:1.12-alpine as caddy_builder
RUN apk add --no-cache git gcc musl-dev
RUN mkdir -p $GOPATH/src/github.com/mholt && cd $GOPATH/src/github.com/mholt&& \
    git clone https://github.com/mholt/caddy.git && \
    cd caddy && git checkout tags/v0.11.5 && pwd
RUN go get github.com/caddyserver/builds
RUN cd $GOPATH/src/github.com/mholt/caddy/caddy && go get . && go run build.go && pwd

FROM gcr.io/distroless/base
WORKDIR /srv
EXPOSE 80
COPY --from=caddy_builder /go/src/github.com/mholt/caddy/caddy /bin/caddy
COPY Caddyfile /etc/Caddyfile
COPY public /srv

ENTRYPOINT ["/bin/caddy"]
CMD ["--conf", "/etc/Caddyfile", "--log", "stdout", "--agree=false"]
