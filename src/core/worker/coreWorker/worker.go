package coreWorker


type Worker interface {
	createWorker(ipAddress string, load int, status string) *Worker
	createDefaultWorker(ipAddress string) *Worker
	CreateWorkersFromFile() []*Worker
}

type WorkerDetails struct {
	IpAddress string
	Load int
	Status string
}

func CreateWorkerDetails(ipAddress string, load int, status string) *WorkerDetails {
	return &WorkerDetails{
		IpAddress: ipAddress,
		Load: load,
		Status: status,
	}
}

func GetWorkerDetails() []WorkerDetails{
	return loadJson()
}
