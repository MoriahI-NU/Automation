package main

import (
	Copilot "Copilot/AnsTestGo"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Results struct {
	Coefficients string
	Time         string
}

func ExperimentPython(set string) Results {

	cmdPy := exec.Command("python", "AnscombeTest.py", set)
	outPy, errPy := cmdPy.Output()

	py_replaced := strings.Replace(string(outPy), "\"Coefficients\": ", "", -1)

	// Split the string by ","
	parts := strings.Split(py_replaced, ", \"Time\": ")

	// Assign values to separate variables
	var firstPart, secondPart string
	if len(parts) >= 2 {
		firstPart = parts[0]
		secondPart = parts[1]
	}

	if errPy != nil {
		fmt.Println(errPy)
		return Results{Coefficients: "", Time: ""}
	}

	return Results{Coefficients: firstPart, Time: secondPart}
}

func ExperimentR(set string) Results {

	cmdR := exec.Command("Rscript", "AnscombeTest.R", set)
	outR, errR := cmdR.Output()

	r_replaced := strings.Replace(string(outR), "\"Coefficients\": ", "", -1)

	// Split the string by ","
	parts := strings.Split(r_replaced, "\"Time\": ")

	// Assign values to separate variables
	var firstPart, secondPart string
	if len(parts) >= 2 {
		firstPart = parts[0]
		secondPart = parts[1]
	}

	if errR != nil {
		fmt.Println(errR)
		return Results{Coefficients: "", Time: ""}
	}

	return Results{Coefficients: firstPart, Time: secondPart}
}

func ExperimentGo(set string) Results {

	response := Copilot.GoTest(set, 500)

	response_concat := strings.Join(response.Coefficient, " ")

	return Results{Coefficients: response_concat, Time: response.Time}
}

func main() {
	anscombeSets := []string{"One", "Two", "Three", "Four"}

	var totalPyTime float64
	var totalRTime float64
	var totalGoTime float64

	for _, set := range anscombeSets {

		fmt.Println("\n///////////////////////////ANSCOMBE SET ", set, "///////////////////////////")

		pyResults := ExperimentPython(set)
		fmt.Println("Python Results for Set", set, ":\n", "Coefficients:", pyResults.Coefficients, "Time:", pyResults.Time)

		pyTime_numonly := strings.Replace(pyResults.Time, "}\r\n", "", -1)

		// Convert the time from string to float64
		pyTime, err := strconv.ParseFloat(pyTime_numonly, 64)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			continue
		}

		totalPyTime += pyTime

		fmt.Println("-----------------------------------------------------")

		rResults := ExperimentR(set)
		fmt.Println("R Results for Set", set, ":\n", "Coefficients:", rResults.Coefficients, "Time:", rResults.Time)

		rTime_numonly := strings.Replace(rResults.Time, "\r\n", "", -1)
		rTime_nobracket := strings.Replace(rTime_numonly, "} ", "", -1)

		// Convert the time from string to float64
		rTime, err := strconv.ParseFloat(rTime_nobracket, 64)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			continue
		}

		totalRTime += rTime

		fmt.Println("-----------------------------------------------------")

		goResults := ExperimentGo(set)
		fmt.Println("Go Results for Set", set, ":\n", "Coefficients:", goResults.Coefficients, "Time:", goResults.Time)

		// Convert the time from string to float64
		goTime, err := strconv.ParseFloat(goResults.Time, 64)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			continue
		}

		totalGoTime += goTime

	}

	// Calculate the average times
	avgPyTime := totalPyTime / float64(len(anscombeSets))
	avgRTime := totalRTime / float64(len(anscombeSets))
	avgGoTime := totalGoTime / float64(len(anscombeSets))

	fmt.Println("\n//////////////////AVERAGE RUNTIMES AND RUNTIME RATIOS FOR ALL SETS//////////////////")

	//Print
	fmt.Println("\nAverage Python Runtime:", avgPyTime)
	fmt.Println("Average R Runtime:", avgRTime)
	fmt.Println("Average Go Runtime:", avgGoTime)

	// Print the tables
	fmt.Println("\n-------------------------PYTHON COMPARISON RATIOS-------------------------")
	printRatioTable("PyRatio", avgPyTime/avgPyTime, avgRTime/avgPyTime, avgGoTime/avgPyTime)
	fmt.Println("\n-------------------------R COMPARISON RATIOS-------------------------")
	printRatioTable("RRatio", avgPyTime/avgRTime, avgRTime/avgRTime, avgGoTime/avgRTime)
	fmt.Println("\n-------------------------GO COMPARISON RATIOS-------------------------")
	printRatioTable("GoRatio", avgPyTime/avgGoTime, avgRTime/avgGoTime, avgGoTime/avgGoTime)
}

func printRatioTable(label string, ratio1 float64, ratio2 float64, ratio3 float64) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)

	// Center-align the strings
	pythonHeader := fmt.Sprintf("%*s", 5, "Python")
	rHeader := fmt.Sprintf("%*s", 3, "R")
	goHeader := fmt.Sprintf("%*s", 4, "Go")

	// Print the centered headers
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", "", pythonHeader, rHeader, goHeader)

	fmt.Fprintf(w, "%s\t%.2f\t%.2f\t%.2f\n", label, ratio1, ratio2, ratio3)

	w.Flush()

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
