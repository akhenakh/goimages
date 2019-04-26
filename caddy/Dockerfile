FROM golang:1.12-alpine as caddy_builder
RUN apk add --no-cache git gcc musl-dev
COPY build /workdir/caddy
RUN cd /workdir/caddy && go build 
RUN go get github.com/akhenakh/minigit

FROM gcr.io/distroless/base
WORKDIR /srv
EXPOSE 80 443
COPY --from=caddy_builder /workdir/caddy/caddy /bin/caddy
COPY --from=caddy_builder /go/bin/minigit /bin/git
COPY Caddyfile /etc/Caddyfile
COPY public /public

ENTRYPOINT ["/bin/caddy"]
CMD ["--conf", "/etc/Caddyfile", "--log", "stdout", "--agree=true"]
