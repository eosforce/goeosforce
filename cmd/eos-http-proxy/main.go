package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	// "net/http"
	// "strings"
	// "bytes"
	//"io/ioutil"

	"github.com/eosforce/goeosforce/p2p"
)

//var peer1 = flag.String("peer1", "localhost:9876", "peer 1")
var peer2 = flag.String("peer2", "localhost:27041", "peer 2")
var showLog = flag.Bool("v", false, "show detail log")
var chainID = flag.String("chain-id", "bd61ae3a031e8ef2f97ee3b0e62776d6d30d4833c8f7c1645c657b149151004b", "peer 1")

func main() {
	flag.Parse()

	fmt.Println("P2P Proxy")

	if *showLog {
		p2p.EnableP2PLogging()
	}
	//defer p2p.SyncLogger()

	cID, err := hex.DecodeString(*chainID)
	if err != nil {
		log.Fatal(err)
	}

	httpRelay := p2p.NewHttpRelay(*peer2,cID)
	httpRelay.RegisterHandler(p2p.StringLoggerHandler)
	httpRelay.Start()



}
