package roundRobinWorker

import "loadbalancer.com/balancer/core/worker/coreWorker"

type RoundRobinWorker struct {
	coreWorker.WorkerDetails
	NextWorker *RoundRobinWorker
	TotalWorkersLength int
}

func createWorker(workDetails coreWorker.WorkerDetails) *RoundRobinWorker {
	return &RoundRobinWorker{
		WorkerDetails: workDetails,
	}
}

func CreateWorkersFromFile() []*RoundRobinWorker {
	workerDetailsFromFile := coreWorker.GetWorkerDetails()
	rrWorkers := []*RoundRobinWorker{}
	for _, line := range workerDetailsFromFile {
		rrWorkers = append(rrWorkers, createWorker(line))
	}
	return rrWorkers
}

func (rrw *RoundRobinWorker) getNextWorker() *RoundRobinWorker {
	return rrw.NextWorker
}

func (rrw *RoundRobinWorker) GetNextAvailableWorker() (*RoundRobinWorker) {
	totalWorkersLength := rrw.TotalWorkersLength
	currWorker := rrw.getNextWorker()
	for totalWorkersLength > 0 {
		if currWorker.Status == "ACTIVE" {
			return currWorker
		}
		currWorker = currWorker.getNextWorker()
		totalWorkersLength--
	}
	return nil
}

func (rrw *RoundRobinWorker) GetIpAddress() string {
	return rrw.WorkerDetails.IpAddress
}


// [
//     {
//         "IpAddress": "127.9.32.1",
//         "Load": 0,
//         "status": "ACTIVE"
//     },
//     {
//         "IpAddress": "182.32.44.2",
//         "Load": 0,
//         "status": "ACTIVE"
//     },
//     {
//         "IpAddress": "132.4.2.33",
//         "Load": 0,
//         "status": "ACTIVE"
//     },
//     {
//         "IpAddress": "23.4.55.5",
//         "Load": 0,
//         "status": "ACTIVE"
//     }
// ]