package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	COSURL "github.com/GGroups/svMisc/cosurl"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	sautx := COSURL.UrlObj{}
	epautx := COSURL.MakeObjUrlEndPoint(sautx)

	mysautxsvr := httpTransport.NewServer(epautx, COSURL.ObjUrlDecodeRequest, COSURL.CommEncodeResponse)

	routeSvr := mux.NewRouter()
	routeSvr.Handle(`/gpwm/misc/getCosUrls`, mysautxsvr).Methods("POST")

	//main loop
	ch := make(chan error, 2)
	go func() {
		log.Println("0.0.0.0:8008", `/gpwm/misc/**`)
		ch <- http.ListenAndServe("0.0.0.0:8008", routeSvr)
	}()
	go func() {
		log.Println("##wait for exit sigint...")
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		ch <- fmt.Errorf("%s", <-c)
	}()

	log.Fatal("MainSvr Terminated", <-ch)
}
