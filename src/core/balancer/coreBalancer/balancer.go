package coreBalancer

import (
	"loadbalancer.com/request"
	"loadbalancer.com/response"
	"loadbalancer.com/worker/coreWorker"
)

type AbstractBalancer interface {
	getAvailableWorkers() []coreWorker.Worker
	getWorkerForRequest(request request.InternalRequest) coreWorker.Worker
	getResponseFromWorker(worker coreWorker.Worker, request request.InternalRequest) response.InternalResponse
	BalanceRequest(request request.InternalRequest) response.InternalResponse
	loadWorkers() error
}


