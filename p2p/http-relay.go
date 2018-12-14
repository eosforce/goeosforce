package p2p

import (
	"fmt"
//	"net"
	"net/http"
	"go.uber.org/zap"
	"strings"
	"bytes"
	"github.com/eosforce/goeosforce"
	"time"
)



type HttpRelay struct {
	destinationPeerAddress string
	cID						eos.Checksum256					
	handlers               []Handler
	inAddr				  []string

}

func NewHttpRelay( destinationPeerAddress string,chainId eos.Checksum256) *HttpRelay {
	return &HttpRelay{
		destinationPeerAddress: destinationPeerAddress,
		cID:chainId,
	}
}

func (r *HttpRelay) RegisterHandler(handler Handler) {

	r.handlers = append(r.handlers, handler)
}

func (r *HttpRelay) FindAddr(addr string) bool{
	for _, value := range r.inAddr {
		if addr == value {
			return true
		}
	}
	return false 
}

func (r *HttpRelay) deleteAddr(addr string){
	for index, value := range r.inAddr {
		if addr == value {
			r.inAddr = append(r.inAddr[:index], r.inAddr[index+1:]...)
			break;
		}
	}
}

func (r *HttpRelay) startProxy(remoteAddress string) {


	p2pLog.Info("Initiating proxy",
		zap.String("peer1", remoteAddress),
		zap.String("peer2", r.destinationPeerAddress))

	destinationPeer := NewOutgoingPeer(r.destinationPeerAddress, "eos-relay", &HandshakeInfo{
		ChainID: r.cID,
	})
	remotePeer := newPeer(remoteAddress, fmt.Sprintf("agent-%s", remoteAddress), false, nil)
	proxy := NewProxy(remotePeer, destinationPeer)

	proxy.RegisterHandlers(r.handlers)

	proxy.ConnectAndStart()	
	r.deleteAddr(remoteAddress)
}

func (r *HttpRelay)AddPeer(addr string) {
	if r.FindAddr(addr) == false {
		r.inAddr = append(r.inAddr, addr)
		go r.startProxy(addr)
	}
}

func (r *HttpRelay)Start() {

	for {
		client := &http.Client{}
		url := "http://127.0.0.1:12560/print"
		reqest, _ := http.NewRequest("GET", url, nil)
		response, _ := client.Do(reqest)

		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		s := buf.String() 
		//bodybyte := readAll(response.Body, bytes.MinRead)
		a := strings.Split(s, "\n")
		for _, value := range a {
			if strings.Compare(value,"") != 0 {
				r.AddPeer(value)
			}
		}
		time.Sleep(time.Duration(3600)*time.Second)
		response.Body.Close()
	} 
}
