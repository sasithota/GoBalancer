package roundRobinBalancer

import (
	"net/http"
	"loadbalancer.com/balancer/core/worker/roundRobinWorker"
	"loadbalancer.com/balancer/core/request"
)

func RouteRequest(req *http.Request) *roundRobinWorker.RoundRobinWorker {
	loadbalancer := GetRoundRobinLoadBalancer()
	internalRequest := request.GenerateInternalRequest(req)
	return GetWorkerForRequest(loadbalancer, internalRequest)
}
