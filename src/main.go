package main

import (
	// "fmt"
	// "loadbalancer.com/balancer/core/balancer/roundRobinBalancer"
	// "loadbalancer.com/balancer/core/request"
	"loadbalancer.com/balancer/server/httpServer"
)

func main() {

	// balancerService := roundRobinBalancer.CreateRoundRobinBalancer()
	// requests := 20
	// for requests > 0 {
	// 	requestData := request.InternalRequest{
	// 		Path: "hello",
	// 		Body: map[string]string{},
	// 		Headers: map[string]string{},
	// 		Method: "world",
	// 	}
	// 	response := balancerService.BalanceRequest(requestData)
	// 	fmt.Println(response)
	// 	requests--
	// }
	// fmt.Println(roundRobinBalanccer.GetAvlWorkersForDebug(balancerService))

	httpServer.StartServer()
}