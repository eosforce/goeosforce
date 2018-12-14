package main
import (
    "fmt"
    "net/http"
   // "strings"
	//"log"
	"github.com/julienschmidt/httprouter"
)

type regAddress struct {
	RegAddress string
}

var regAddresses []regAddress

func find(addr string) bool {
	for _, value := range regAddresses {
		if addr == value.RegAddress {
			return true
		}
	}
	return false 
}

func register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bFind := find(p.ByName("name"))
	if bFind == false {
		var getAddress regAddress
		getAddress.RegAddress = p.ByName("name")
		regAddresses = append(regAddresses,getAddress)
	}
	
	print(w,r,p)
}

func print(w http.ResponseWriter, r *http.Request,p httprouter.Params) {
	for _, value := range regAddresses {
		fmt.Fprintf(w,"%s\n",value.RegAddress);
	} 
}

func delete(w http.ResponseWriter, r *http.Request,p httprouter.Params) {
	for index, value := range regAddresses {
		if p.ByName("name") == value.RegAddress {
			regAddresses = append(regAddresses[:index], regAddresses[index+1:]...)
			break;
		}
	} 
	print(w,r,p)
}

func main() {
	mux := httprouter.New()
	mux.GET("/add/:name",register)
	mux.GET("/print",print)
	mux.GET("/delete/:name",delete)
	server := http.Server{
		Addr:"127.0.0.1:12560",
		Handler:mux,
	}
    server.ListenAndServe()
}
