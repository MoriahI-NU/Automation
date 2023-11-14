package Copilot

import (
	"fmt"
	"math/rand"
	"time"
)

var nRuns = 500

type Response struct {
	Coefficient []float64 `json:"coefficients"`
	Time        float64   `json:"time"`
}

func linearRegression(x, y []float64) (float64, float64) {
	var sumX, sumY, sumXY, sumX2 float64
	n := float64(len(x))

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / n

	return intercept, slope
}

func runLinearRegression(dataX, dataY [][]float64) (float64, float64) {
	var totalIntercept, totalSlope []float64

	for i := 0; i < nRuns; i++ {
		index := rand.Intn(len(dataX))
		intercept, slope := linearRegression(dataX[index], dataY[index])
		totalIntercept = append(totalIntercept, intercept)
		totalSlope = append(totalSlope, slope)
	}

	avgIntercept := sum(totalIntercept) / float64(len(totalIntercept))
	avgSlope := sum(totalSlope) / float64(len(totalSlope))

	return avgIntercept, avgSlope
}

func sum(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func GoTest(set string, nRuns int) Response {

	start := time.Now()
	var responseGo Response

	var times []float64

	var anscombe = map[string]map[string][]float64{
		"One": {
			"x": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			"y": {8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		},
		"Two": {
			"x": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			"y": {9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
		},
		"Three": {
			"x": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			"y": {7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		},
		"Four": {
			"x": {8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
			"y": {6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
		},
	}

	for i := 0; i < nRuns; i++ {
		dataX := [][]float64{anscombe[set]["x"]}
		dataY := [][]float64{anscombe[set]["y"]}

		avgIntercept, avgSlope := runLinearRegression(dataX, dataY)

		responseGo.Coefficient = []float64{avgIntercept, avgSlope}
		times = append(times, time.Since(start).Seconds())
	}
	responseGo.Time = sum(times) / float64(len(times))
	fmt.Printf("%.7f\n", responseGo.Time)

	return responseGo
}

//////ORIGINAL COPILOT RESPONSE: (started out as a file, and not as a copilot package)///////
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// var nRuns = 500

// type Response struct {
// 	Coefficient []float64 `json:"coefficients"`
// 	Time        float64   `json:"time"`
// }

// func linearRegression(x, y []float64) (float64, float64) {
// 	var sumX, sumY, sumXY, sumX2 float64
// 	n := float64(len(x))

// 	for i := 0; i < len(x); i++ {
// 		sumX += x[i]
// 		sumY += y[i]
// 		sumXY += x[i] * y[i]
// 		sumX2 += x[i] * x[i]
// 	}

// 	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
// 	intercept := (sumY - slope*sumX) / n

// 	return intercept, slope
// }

// func runLinearRegression(dataX, dataY [][]float64) (float64, float64) {
// 	var totalIntercept, totalSlope []float64

// 	for i := 0; i < nRuns; i++ {
// 		index := rand.Intn(len(dataX))
// 		intercept, slope := linearRegression(dataX[index], dataY[index])
// 		totalIntercept = append(totalIntercept, intercept)
// 		totalSlope = append(totalSlope, slope)
// 	}

// 	avgIntercept := sum(totalIntercept) / float64(len(totalIntercept))
// 	avgSlope := sum(totalSlope) / float64(len(totalSlope))

// 	return avgIntercept, avgSlope
// }

// func sum(numbers []float64) float64 {
// 	sum := 0.0
// 	for _, number := range numbers {
// 		sum += number
// 	}
// 	return sum
// }

// func GoTest(set string, nRuns int) Response {
// 	fmt.Println("Hey there this is Go")
// 	start := time.Now()
// 	var responseGo Response

// 	var times []float64

// 	var anscombe = map[string]map[string][]float64{
// 		"One": {
// 			"x": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
// 			"y": {8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
// 		},
// 		"Two": {
// 			"x": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
// 			"y": {9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
// 		},
// 		"Three": {
// 			"x": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
// 			"y": {7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
// 		},
// 		"Four": {
// 			"x": {8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
// 			"y": {6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
// 		},
// 	}

// 	for i := 0; i < nRuns; i++ {
// 		dataX := [][]float64{anscombe[set]["x"]}
// 		dataY := [][]float64{anscombe[set]["y"]}

// 		avgIntercept, avgSlope := runLinearRegression(dataX, dataY)

// 		responseGo.Coefficient = []float64{avgIntercept, avgSlope}
// 		times = append(times, time.Since(start).Seconds())
// 	}
// 	responseGo.Time = sum(times) / float64(len(times))
// 	fmt.Printf("%.7f\n", responseGo.Time)

// 	return responseGo
// }

// func main() {
// 	fmt.Println("HEY THIS IS GO")

// 	anscombesets := []string{"One", "Two", "Three", "Four"}
// 	for _, set := range anscombesets {
// 		response := GoTest(set, 500)
// 		fmt.Println("anscombe set:", set, response)
// 	}

// 	return
// }
