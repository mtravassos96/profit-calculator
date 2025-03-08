package main

import (
	"fmt"
	"os"
	"errors"
)

// Goals
// 1) Validate user input
//    - Show error message & exit if invalid input is provided
//    - No negative numbers
//    - Not 0
// 2) Store calculated results into file

const userFinancialsFile = "financials.txt"

func main() {
	revenue, expenses, taxRate, err := getUserInput()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	formattedEBT := fmt.Sprintf("Earnings Before Tax: USD %.2f\n", earningsBeforeTax)
	formattedProfit := fmt.Sprintf("Earnings After Tax: USD %.2f\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)

	writeFinancialsToFile(formattedEBT, formattedProfit, formattedRatio)

	// fmt.Printf("Earnings Before Tax: USD %.2f\nEarnings After Tax: USD %.2f\nRatio: %.2f\n", earningsBeforeTax, profit, ratio)
}

func getUserInput() (float64, float64, float64, error) {
	var revenue, expenses, taxRate float64
	
	fmt.Print("Inform your revenue (USD): ")
	fmt.Scan(&revenue)

	fmt.Print("Inform your expenses (USD): ")
	fmt.Scan(&expenses)

	fmt.Print("Inform the tax rate: ")
	fmt.Scan(&taxRate)

	if revenue <= 0 || expenses < 0 || taxRate <= 0 {
		return 0, 0, 0, errors.New("input values must be a positive number")
	}

	return revenue, expenses, taxRate, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio :=  profit / earningsBeforeTax

	return earningsBeforeTax, profit, ratio
}

func writeFinancialsToFile(formattedEBT, formattedProfit, formattedRatio string) {
	combineTexts := formattedEBT + formattedProfit + formattedRatio
	os.WriteFile(userFinancialsFile, []byte(combineTexts), 0644)
}