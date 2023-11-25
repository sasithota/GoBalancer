package roundRobinWorker

import "loadbalancer.com/worker/coreWorker"

type RoundRobinWorker struct {
	coreWorker.WorkerDetails
	NextWorker *RoundRobinWorker
	TotalWorkersLength int
}

func createWorker(ipAddress string, load int, status string) *RoundRobinWorker {
	return &RoundRobinWorker{
		WorkerDetails: *coreWorker.CreateWorkerDetails(ipAddress, load, status),
	}
}

func createDefaultWorker(ipAddress string) *RoundRobinWorker {
	return createWorker(ipAddress, 0, "ACTIVE")
}

func CreateWorkersFromFile() []*RoundRobinWorker {
	linesFromFile := coreWorker.GetLinesFromFile()
	rrWorkers := []*RoundRobinWorker{}
	for _, line := range linesFromFile {
		rrWorkers = append(rrWorkers, createDefaultWorker(line))
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
