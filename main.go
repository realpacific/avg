package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PriceCount struct {
	price string
	count string
}

func init() {

}

// Takes arguments and calculates average or calculates the amount (i.e. count) to get that specific average.
//
// The first argument without ":" is considered as the expected price. If provided, it calculates the amount to be bought to get a specific average.
// If all arguments provided is price:count pair, then calculates the average.
func main() {
	arguments := os.Args[1:]
	expectedResult := ""

	if !strings.Contains(arguments[0], ":") {
		// case of expected result provided
		expectedResult = arguments[0]
		arguments = arguments[1:]

		fmt.Printf("Expected result: %v\n", expectedResult)

		// Validate that one of the remaining arguments must contain the variable "x" in its count segment
		variableCount := 0
		for _, arg := range arguments {
			if strings.Contains(arg, ":x") {
				variableCount++
			}
		}
		if variableCount != 1 {
			panic(any(fmt.Sprintln("invalid input. must provide only one variable. given=", arguments)))
		}
	}
	fmt.Printf("Arguments: %v\n", arguments)

	priceCountPairs := make([]PriceCount, len(arguments))

	for i, arg := range arguments {
		trim := strings.TrimSpace(arg)
		segments := strings.Split(trim, ":")
		priceStr := segments[0]
		count := segments[1]
		priceCountPairs[i] = PriceCount{price: priceStr, count: count}
	}

	if expectedResult == "" {
		calculateAverage(priceCountPairs)
	} else {
		calculateCountForExpectedPrice(priceCountPairs, expectedResult)
	}
}

func calculateCountForExpectedPrice(priceCountPairs []PriceCount, expectedResult string) {
	var unknownPriceCountPair PriceCount
	var knownPriceCountPair []PriceCount
	separateUnknownAndKnowns(priceCountPairs, &unknownPriceCountPair, &knownPriceCountPair)

	expectedPrice := toFloat(expectedResult)

	denominator := toFloat(unknownPriceCountPair.price) - expectedPrice
	var numerator float64 = 0
	fmt.Printf("PriceCountPairs: %v\n", knownPriceCountPair)

	for _, pc := range knownPriceCountPair {
		numerator += expectedPrice * toFloat(pc.count)
	}
	for _, pc := range knownPriceCountPair {
		numerator -= toFloat(pc.count) * toFloat(pc.price)
	}
	fmt.Println("-------------------")
	fmt.Printf(
		"You need %.2f @ %.2f to get %.2f\n",
		numerator/denominator,
		toFloat(unknownPriceCountPair.price),
		expectedPrice,
	)
}

func toFloat(price string) float64 {
	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		panic(any(fmt.Sprintf("invalid argument provided: %v %v\n\n", price, err)))
	}
	return p
}

func separateUnknownAndKnowns(priceCountMap []PriceCount, unknownPriceCountPair *PriceCount, knownPriceCountPair *[]PriceCount) {
	for _, pc := range priceCountMap {
		if pc.count == "x" {
			*unknownPriceCountPair = pc
		} else {
			*knownPriceCountPair = append(*knownPriceCountPair, pc)
		}
	}
}

func calculateAverage(priceCountMap []PriceCount) {
	var totalAmount float64
	var totalCount int
	for _, pc := range priceCountMap {
		count, _ := strconv.Atoi(pc.count)
		totalCount += count
		price := toFloat(pc.price)
		totalAmount += price * float64(count)
	}
	fmt.Printf("Price: %v\n", totalAmount)
	fmt.Printf("Count: %v\n", totalCount)
	fmt.Println("-------------------")
	fmt.Printf("Average: %.2f\n", totalAmount/float64(totalCount))
}
