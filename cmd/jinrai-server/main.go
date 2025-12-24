package main

import (
	"flag"
	"log"
	"strings"

	"github.com/jinrai-js/server/internal/lib/components"
	"github.com/jinrai-js/server/internal/lib/jinrai"
)

type Table struct {
	Url  string `json:"url"`
	Data string `json:"data"`
}

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
		ssr.Debug()
		log.Println("+ verbose")
	}

	if *server != "" {
		ssr.SetStringProxy(*server)
	}

	if *caching != "" {
		ssr.SetChashing(strings.Split(*caching, ","))
	}

	components.Add("tbl", func(props Table) string {
		return "[table: URL" + props.Url + "||" + props.Data + "]"
	})

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
