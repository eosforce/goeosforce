package main

import (
	"fmt"

	"flag"

	"github.com/eosforce/goeosforce/p2p"
)

var peer = flag.String("peer", "127.0.0.1:27041", "peer")
var listeningAddress = flag.String("listening-address", "0.0.0.0:18888", "address on with the relay will listen")
var showLog = flag.Bool("v", false, "show detail log")

func main() {
	flag.Parse()

	if *showLog {
		p2p.EnableP2PLogging()
	}
	defer p2p.SyncLogger()

	relay := p2p.NewRelay(*listeningAddress, *peer)
	relay.RegisterHandler(p2p.StringLoggerHandler)
	fmt.Println(relay.Start())
}
