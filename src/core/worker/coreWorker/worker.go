package coreWorker

import (
	"bufio"
	"os"
	"fmt"
)

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

func GetLinesFromFile() []string {
	file, err := os.Open("core/worker/workers.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return result
}
