package httpServer


import (
	"net/http"
	"log"
	"loadbalancer.com/balancer/server/handler"
	"fmt"
)


func StartServer() error {
	http.HandleFunc("/", handler.ServerRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}