/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-04
 * @fileoverview How long it will take the user to pay off post-secondary education.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// declare variables
	var startingAmountString string
	var startingAmountNumber float64
	var rateString string
	var rateNumber float64
	var requiredAmountString string
	var requiredAmountNumber float64
	var years int = 0

	reader := bufio.NewReader(os.Stdin)

	// input
	// starting amount in bank account
	fmt.Print("Enter the starting bank account amount: ")
	startingAmountString, _ = reader.ReadString('\n')
	startingAmountString = strings.TrimSpace(startingAmountString)
	startingAmountNumber, _ = strconv.ParseFloat(startingAmountString, 64)

	// interest rate
	fmt.Print("Enter the yearly interest rate (as a percentage): ")
	rateString, _ = reader.ReadString('\n')
	rateString = strings.TrimSpace(rateString)
	rateNumber, _ = strconv.ParseFloat(rateString, 64)

	// amount needed for post-secondary
	fmt.Print("Enter the amount needed for post-secondary education: ")
	requiredAmountString, _ = reader.ReadString('\n')
	requiredAmountString = strings.TrimSpace(requiredAmountString)
	requiredAmountNumber, _ = strconv.ParseFloat(requiredAmountString, 64)

	// processing
	// turn rate into decimal
	rateNumber = rateNumber / 100

	// loop starts
	for startingAmountNumber < requiredAmountNumber {
		startingAmountNumber = startingAmountNumber + (startingAmountNumber * rateNumber)
		years++
	}

	// output
	fmt.Printf("It will take %d years for the starting bank account to reach the required amount for post-secondary education. ", years)

	fmt.Println("\nDone.")
}
