package roundRobinBalancer

import (
	"loadbalancer.com/request"
	"loadbalancer.com/response"
	"loadbalancer.com/worker/roundRobinWorker"
)

type RoundRobinBalancer struct {
	Workers []*roundRobinWorker.RoundRobinWorker
	NextWorker *roundRobinWorker.RoundRobinWorker
}

// Load workers from File
func (rrb *RoundRobinBalancer) loadWorkers() error {
	rrb.Workers = roundRobinWorker.CreateWorkersFromFile()
	return nil
}

// Get the next available worker node that can handle the request
func (rrb *RoundRobinBalancer) getWorkerForRequest(request request.InternalRequest) *roundRobinWorker.RoundRobinWorker {
	nextWorker := rrb.NextWorker.GetNextAvailableWorker()
	rrb.NextWorker = nextWorker
	return nextWorker
}

// Ping the worker with the request and return the response
func (rrb *RoundRobinBalancer) getResponseFromWorker(rrw roundRobinWorker.RoundRobinWorker, request request.InternalRequest) response.InternalResponse {
	// TODO
	return response.CreateFakeResponse(rrw.IpAddress)
}

func (rrb *RoundRobinBalancer) BalanceRequest(request request.InternalRequest) response.InternalResponse {
	rrw := rrb.getWorkerForRequest(request)
	return rrb.getResponseFromWorker(*rrw, request)
}

// For Round robin - load the workers and link the works in circular fashion
func (rrb *RoundRobinBalancer) initaliseLeastConnectionWorkers() error {
	rrb.loadWorkers()
	totalLength := len(rrb.Workers)
	for indx, _ := range rrb.Workers {
		rrb.Workers[indx].NextWorker = rrb.Workers[(indx+1)%totalLength]
		rrb.Workers[indx].TotalWorkersLength = totalLength
	}
	rrb.NextWorker = rrb.Workers[0]
	return nil
}

// contructor to create RoundRobingBalancer
func CreateRoundRobinBalancer() *RoundRobinBalancer {
	rrb := &RoundRobinBalancer{}
	rrb.initaliseLeastConnectionWorkers()
	return rrb
}

// For debugging purpose
func GetAvlWorkersForDebug(rrb *RoundRobinBalancer) []roundRobinWorker.RoundRobinWorker {
	workers := []roundRobinWorker.RoundRobinWorker{}
	for _, workerNode := range rrb.Workers {
		workers = append(workers, *workerNode)
	}
	return workers
}

