package coreWorker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func loadJson() []WorkerDetails{
	// Specify the path to the JSON file
	filePath := "core/worker/workers.json"

	// Read the entire file content
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	// Create a variable of type Person to unmarshal the JSON data
	var workerDetails []WorkerDetails

	// Unmarshal the JSON data into the Person struct
	err = json.Unmarshal(fileContent, &workerDetails)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	return workerDetails
}
