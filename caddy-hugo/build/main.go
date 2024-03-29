package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"

	// plug in plugins here, for example:
	// _ "import/path/here"
	_ "github.com/akhenakh/caddy-puregit"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
