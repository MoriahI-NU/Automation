package main

import (
	Copilot "Copilot/AnsTestGo"
	"fmt"
	"os/exec"
)

type Results struct {
	Output string
	Error  error
}

func ExperimentPython(set string) Results {
	cmdPy := exec.Command("python", "AnscombeTest.py", set)
	outPy, errPy := cmdPy.Output()
	if errPy != nil {
		fmt.Println(errPy)
		return Results{Output: "", Error: errPy}
	}

	return Results{Output: string(outPy), Error: nil}
}

func ExperimentR(set string) Results {
	cmdR := exec.Command("Rscript", "AnscombeTest.R", set)
	outR, errR := cmdR.Output()
	if errR != nil {
		fmt.Println(errR)
		return Results{Output: "", Error: errR}
	}

	return Results{Output: string(outR), Error: nil}
}

func ExperimentGo(set string) Copilot.Response {
	response := Copilot.GoTest(set, 500)
	return response
}

func main() {
	anscombeSets := []string{"One", "Two", "Three", "Four"}

	for _, set := range anscombeSets {
		pyResults := ExperimentPython(set)
		fmt.Println("Python Results for Set", set, ":", pyResults)

		rResults := ExperimentR(set)
		fmt.Println("R Results for Set", set, ":", rResults)

		goResults := ExperimentGo(set)
		fmt.Println("Go Results for Set", set, ":", goResults)
	}
}

/////Original Copilot Response////
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os/exec"
// )

// func main() {
// 	anscombeSets := []string{"1", "2", "3", "4"}

// 	for _, set := range anscombeSets {
// 		// Run AnscombeTest.py
// 		cmdPy := exec.Command("python", "AnscombeTest.py", set)
// 		outPy, errPy := cmdPy.Output()
// 		if errPy != nil {
// 			fmt.Println(errPy)
// 		}
// 		var pyResults map[string]interface{}
// 		json.Unmarshal(outPy, &pyResults)
// 		fmt.Println("Python Results for Set", set, ":", pyResults)

// 		// Run AnscombeTest.R
// 		cmdR := exec.Command("Rscript", "AnscombeTest.R", set)
// 		outR, errR := cmdR.Output()
// 		if errR != nil {
// 			fmt.Println(errR)
// 		}
// 		var rResults map[string]interface{}
// 		json.Unmarshal(outR, &rResults)
// 		fmt.Println("R Results for Set", set, ":", rResults)

// 		// Run AnscombeTest.go
// 		cmdGo := exec.Command("go", "run", "AnscombeTest.go", set)
// 		outGo, errGo := cmdGo.Output()
// 		if errGo != nil {
// 			fmt.Println(errGo)
// 		}
// 		var goResults map[string]interface{}
// 		json.Unmarshal(outGo, &goResults)
// 		fmt.Println("Go Results for Set", set, ":", goResults)
// 	}
// }
