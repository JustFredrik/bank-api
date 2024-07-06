package main

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/justfredrik/bank-api/internal/initializer"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Running on %d CPU(s)\n", numCPU)

	camtDoc, err := initializer.LoadLocalData("./data/camt053.xml")
	if err != nil {
		fmt.Println(err)
	}

	jsonBytes, _ := json.MarshalIndent(camtDoc.BankStatement, "", "	")

	fmt.Print(string(jsonBytes))

}
