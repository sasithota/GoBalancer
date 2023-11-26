package handler

import (
	"net/http"
	"fmt"
	"loadbalancer.com/balancer/core/balancer/roundRobinBalancer"
)

func ServerRequest(w http.ResponseWriter, r *http.Request) {
	worker := roundRobinBalancer.RouteRequest(r)

	// Test load balancer for test ips
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(worker.WorkerDetails.IpAddress))

	// Actual Redirect
	ipAddress := fmt.Sprintf("%s%s", "http://", worker.GetIpAddress())
	http.Redirect(w, r, ipAddress, http.StatusTemporaryRedirect)
}

