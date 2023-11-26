package roundRobinBalancer

import (
	"sync"
	"loadbalancer.com/balancer/core/request"
	"loadbalancer.com/balancer/core/worker/roundRobinWorker"
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

// For Round robin - load the workers and link the works in circular fashion
func (rrb *RoundRobinBalancer) initaliseRoundRobinWorkers() error {
	rrb.loadWorkers()
	totalLength := len(rrb.Workers)
	for indx, _ := range rrb.Workers {
		rrb.Workers[indx].NextWorker = rrb.Workers[(indx+1)%totalLength]
		rrb.Workers[indx].TotalWorkersLength = totalLength
	}
	rrb.NextWorker = rrb.Workers[0]
	return nil
}

var (
	loadbalancer *RoundRobinBalancer
	once sync.Once
)

func GetRoundRobinLoadBalancer() *RoundRobinBalancer {
	once.Do(func() {
		loadbalancer = &RoundRobinBalancer{}
		loadbalancer.initaliseRoundRobinWorkers()
	})
	return loadbalancer
}

func GetWorkerForRequest(rrb *RoundRobinBalancer, request request.InternalRequest) *roundRobinWorker.RoundRobinWorker {
	return rrb.getWorkerForRequest(request)
}

// contructor to create RoundRobingBalancer
func CreateRoundRobinBalancer() *RoundRobinBalancer {
	rrb := &RoundRobinBalancer{}
	rrb.initaliseRoundRobinWorkers()
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

