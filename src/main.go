package main

import (
	"fmt"
	"loadbalancer.com/balancer/roundRobinBalancer"
	"loadbalancer.com/request"
)

func main() {

	balancerService := roundRobinBalancer.CreateRoundRobinBalancer()
	requests := 20
	for requests > 0 {
		requestData := request.InternalRequest{
			Path: "hello",
			Body: map[string]string{},
			Headers: map[string]string{},
			Method: "world",
		}
		response := balancerService.BalanceRequest(requestData)
		fmt.Println(response)
		requests--
	}
	// fmt.Println(roundRobinBalancer.GetAvlWorkersForDebug(balancerService))
}