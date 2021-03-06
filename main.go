package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kaakaa/matterpoll-emoji/poll"
)

var config = flag.String("c", "config.json", "optional path to the config file")

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c, err := poll.LoadConf(*config)
	if err != nil {
		log.Fatal(err)
	}
	ps := poll.Server{c}
	http.HandleFunc("/poll", ps.Cmd)
	if err := http.ListenAndServe(c.Listen, nil); err != nil {
		log.Fatal(err)
	}
}
