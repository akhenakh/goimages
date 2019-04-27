package main

import (
	"github.com/mholt/caddy/caddy/caddymain"

	// plug in plugins here, for example:
	// _ "import/path/here"
	_ "github.com/BTBurke/caddy-jwt"
	_ "github.com/akhenakh/caddy-git"
	_ "github.com/caddyserver/dnsproviders/cloudflare"
	_ "github.com/caddyserver/dnsproviders/gandiv5"
	_ "github.com/caddyserver/dnsproviders/godaddy"
	_ "github.com/caddyserver/dnsproviders/googlecloud"
	_ "github.com/caddyserver/dnsproviders/ovh"
	_ "github.com/captncraig/cors"
	_ "github.com/epicagency/caddy-expires"
	_ "github.com/miekg/caddy-prometheus"
	_ "github.com/nicolasazrak/caddy-cache"
	_ "github.com/pieterlouw/caddy-grpc"
	_ "github.com/pyed/ipfilter"
	_ "github.com/tarent/loginsrv/caddy"
	_ "github.com/wmark/caddy.upload"
	_ "github.com/xuqingfeng/caddy-rate-limit"
	_ "github.com/zikes/gopkg"
)

func main() {
	// optional: disable telemetry
	// caddymain.EnableTelemetry = false
	caddymain.Run()
}
