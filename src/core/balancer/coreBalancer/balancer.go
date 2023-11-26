package coreBalancer

import (
	"loadbalancer.com/balancer/core/request"
	"loadbalancer.com/balancer/core/worker/coreWorker"
)

type AbstractBalancer interface {
	getAvailableWorkers() []coreWorker.Worker
	getWorkerForRequest(request request.InternalRequest) coreWorker.Worker
	loadWorkers() error
}


