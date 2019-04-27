FROM golang:1.12-alpine as caddy_builder
RUN apk add --no-cache git gcc musl-dev
RUN mkdir -p /workdir
COPY build /workdir/build
RUN cd /workdir/build && CGO_ENABLED=0 go build 
WORKDIR /workdir
RUN git clone https://github.com/gohugoio/hugo.git && cd hugo && git checkout tags/v0.55.4 && CGO_ENABLED=0 go build
RUN CGO_ENABLED=0 go get github.com/akhenakh/minigit

FROM gcr.io/distroless/base
WORKDIR /srv
EXPOSE 80 443
COPY --from=caddy_builder /workdir/build/caddy /bin/caddy
COPY --from=caddy_builder /go/bin/minigit /bin/git
COPY --from=caddy_builder /workdir/hugo/hugo /bin/hugo
COPY Caddyfile /etc/Caddyfile
COPY public /public

CMD ["/bin/caddy", "--conf", "/etc/Caddyfile", "--log", "stdout", "--agree=true"]
