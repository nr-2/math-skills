package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"math-skills/nor"
)

func main() {
	// get args from termial
	argement := os.Args[1:]
	if len(argement) != 1 {
		fmt.Println("Argement givin is not equel to one file! Plz try again.")
		return
	}
	// take the txt file from args
	file := argement[0]
	datafile, err := os.ReadFile(file)
	if err != nil {
		err.Error()
		return
	}
	// make array from the data
	datastr := strings.Split(string(datafile), "\n")
	data, err := nor.StringToInt(datastr)
	if err != nil {
		fmt.Println("Error in data givin!")
		return
	}
	sort.Ints(data)

	// appling the function to get the answer
	datamean := math.Round(nor.Mean(data))
	datamedian := math.Round(nor.Median(data))
	datavariance := math.Round(nor.Variance(data, datamean))
	datastandarddeviation := math.Round(math.Sqrt(datavariance))
	// print the answer
	// fmt.Println(data)
	// fmt.Println(len(data))
	fmt.Printf("Average: %d\n", int(datamean))
	fmt.Printf("Median: %d\n", int(datamedian))
	fmt.Printf("Variance: %d\n", int(datavariance))
	fmt.Printf("Standard Deviation: %d\n", int(datastandarddeviation))
}