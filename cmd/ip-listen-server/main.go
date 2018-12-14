package main
import (
    "fmt"
	"net"
	"net/http"
	
   // "strings"
	//"log"
	//"github.com/julienschmidt/httprouter"
	//"go.uber.org/zap"
)

func RegisterAddr(addr string) {
	client := &http.Client{}
	url := "http://127.0.0.1:12560/add/" + addr
	fmt.Println(url)
	reqest, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(reqest)
	fmt.Println(response.Body)
	defer response.Body.Close()
}

func Start() error {

	for {
		ln, err := net.Listen("tcp", "localhost:15222")
		if err != nil {
			return fmt.Errorf("peer init: listening %s: %s", "localhost:15222", err)
		}

		//p2pLog.Info("Accepting connection", zap.String("listen", r.listeningAddress))

		for {
			conn, err := ln.Accept()
			if err != nil {
				return fmt.Errorf("peer init: listening %s: %s", "localhost:15222", err)
				break
			}
			//p2pLog.Info("Connected to", zap.Stringer("remote", conn.RemoteAddr()))
			fmt.Println("Get Connect ",conn.RemoteAddr())
			RegisterAddr(conn.RemoteAddr().String())
		}
	}

	return nil
}


func main() {
	Start()
}
