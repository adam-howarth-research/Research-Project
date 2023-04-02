package main

import (
	"encoding/json"
	"fmt"
)

type InputData struct {
	Level    string  `json:"level"`
	Start    int64   `json:"start"`
	Nodes    int     `json:"nodes"`
	Factor   int     `json:"factor"`
	Elapsed  float64 `json:"elapsed"`
	Time     int64   `json:"time"`
	Message  string  `json:"message"`
}

func main() {
	inputDataJSONA := []string{
		// Add your input data as JSON strings here for both algorithms
	}

	var inputDataListA []InputData
	for _, data := range inputDataJSONA {
		var inputData InputData
		err := json.Unmarshal([]byte(data), &inputData)
		if err != nil {
			fmt.Println("Error unmarshalling input data:", err)
			return
		}
		inputDataListA = append(inputDataListA, inputData)
	}

	inputDataJSONB := []string{
		// Add your input data as JSON strings here for both algorithms
	}

	var inputDataListB []InputData
	for _, data := range inputDataJSONB {
		var inputData InputData
		err := json.Unmarshal([]byte(data), &inputData)
		if err != nil {
			fmt.Println("Error unmarshalling input data:", err)
			return
		}
		inputDataListB = append(inputDataListB, inputData)
	}

	// Assuming that inputDataList is sorted by time and has an even number of entries
	dominanceScore := 0
	for i := 0; i < len(inputDataListA); i++ {
		algorithmA := inputDataListA[i]
		algorithmB := inputDataListB[i]


		if algorithmA.Elapsed < algorithmB.Elapsed {
			dominanceScore++
		} else if algorithmA.Elapsed > algorithmB.Elapsed {
			dominanceScore--
		}
	}

	fmt.Printf("Dominance Score: %d\n", dominanceScore)
	if dominanceScore > 0 {
		fmt.Println("Algorithm A is generally more resource-efficient than Algorithm B.")
	} else if dominanceScore < 0 {
		fmt.Println("Algorithm B is generally more resource-efficient than Algorithm A.")
	} else {
		fmt.Println("Both algorithms have similar resource consumption profiles.")
	}
}
