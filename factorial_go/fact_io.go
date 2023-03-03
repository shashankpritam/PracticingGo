package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func factorial(n int64) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("factorial is undefined for negative integers")
	}
	result := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result, nil
}

func main() {
	// Open input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()
	
	// Create output file
	outputFile, err := os.Create("output.md")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()
	
	// Write Markdown table header
	_, err = fmt.Fprintln(outputFile, "| Input | Factorial |")
	if err != nil {
		fmt.Println("Error writing output to file:", err)
		return
	}
	_, err = fmt.Fprintln(outputFile, "| --- | --- |")
	if err != nil {
		fmt.Println("Error writing output to file:", err)
		return
	}
	
	// Read input file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// Parse input line
		inputLine := strings.TrimSpace(scanner.Text())
		if inputLine == "" {
			continue // skip empty lines
		}
		n, err := strconv.ParseInt(inputLine, 10, 64)
		if err != nil {
			// Write error to output file
			_, err = fmt.Fprintf(outputFile, "| %s | Error converting input to integer: %s |\n", inputLine, err)
			if err != nil {
				fmt.Println("Error writing output to file:", err)
				return
			}
			continue
		}
		
		// Compute factorial
		result, err := factorial(n)
		if err != nil {
			// Write error to output file
			_, err = fmt.Fprintf(outputFile, "| %d | Error computing factorial: %s |\n", n, err)
			if err != nil {
				fmt.Println("Error writing output to file:", err)
				return
			}
			continue
		}
		
		// Write output to file
		_, err = fmt.Fprintf(outputFile, "| %d | %v |\n", n, result)
		if err != nil {
			fmt.Println("Error writing output to file:", err)
			return
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
}
