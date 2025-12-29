package main

import (
	"flag"
	"log"
	"strings"

	"github.com/jinrai-js/server/internal/components/pagination"
	"github.com/jinrai-js/server/internal/components/seo_table"
	"github.com/jinrai-js/server/internal/lib/components"
	"github.com/jinrai-js/server/internal/lib/jinrai"
	"github.com/jinrai-js/server/internal/lib/jlog"
)

func main() {
	dist, server, meta, port, assets, caching, verbose := initFlags()
	log.Println("dist", *dist)
	ssr := jinrai.NewX(*dist)
	if *meta != "" {
		ssr.SetMeta(*meta)
	}

	log.Printf("run: http://localhost:%d\n", *port)

	if *assets {
		log.Println("+ serve assets")
		ssr.ServeAssets(true)
	}

	if *verbose {
		jlog.Active = true
		log.Println("+ verbose")
	}

	if *server != "" {
		ssr.SetStringProxy(*server)
	}

	if *caching != "" {
		ssr.SetChashing(strings.Split(*caching, ","))
	}

	// components.Add("table", table.Component)
	components.Add("seo-table", seo_table.Component)
	components.Add("pagination", pagination.Component)

	ssr.ServeX(*port)
}

func initFlags() (dist, proxy, meta *string, port *int, assets *bool, caching *string, verbose *bool) {
	dist = flag.String("dist", "dist", "dist folder")
	proxy = flag.String("proxy", "", `list of proxy servers
	example:
	/api=http://localhost,/profile=http//localhost:3000`)
	port = flag.Int("port", 80, "html server port")
	meta = flag.String("meta", "", "meta date url")

	assets = flag.Bool("a", false, "serve assets")
	verbose = flag.Bool("v", false, "verbose")
	caching = flag.String("caching", "", "caching proxy requests (example: \"/api/v1,/api/v2\")")

	flag.Parse()
	return
}
