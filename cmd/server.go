package main

import (
	"flag"
	"log"

	"github.com/jinrai-js/go/pkg/jinrai"
)

func main() {
	dist, server, meta, port, assets, verbose := initFlags()
	ssr := jinrai.NewX(*dist, *server, meta)

	log.Printf("jinrai: http://localhost:%d\n", *port)

	if *assets {
		log.Println("+ serve assets")
		ssr.ServeAssets(true)
	}

	if *verbose {
		ssr.Debug()
		log.Println("+ verbose")
	}

	ssr.ServeX(*port)
}

func initFlags() (dist, server, meta *string, port *int, assets *bool, verbose *bool) {
	dist = flag.String("dist", "dist", "dist folder")
	server = flag.String("proxy", "/Api=http://localhost", `list of proxy servers
	example:
	/api=http://localhost,/profile=http//localhost:3000`)
	port = flag.Int("port", 80, "html server port")
	meta = flag.String("meta", "/Api/Meta/GetMetaDate", "meta date url")

	assets = flag.Bool("a", false, "serve assets")
	verbose = flag.Bool("v", true, "verbose")

	flag.Parse()
	return
}
