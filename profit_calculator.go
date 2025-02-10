package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Inform your revenue (USD): ")
	fmt.Scan(&revenue)

	fmt.Print("Inform your expenses (USD): ")
	fmt.Scan(&expenses)

	fmt.Print("Inform the tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	formattedEBT := fmt.Sprintf("Earnings Before Tax: USD %.2f\n", earningsBeforeTax)
	formattedProfit := fmt.Sprintf("Earnings After Tax: USD %.2f\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)

	// fmt.Printf("Earnings Before Tax: USD %.2f\nEarnings After Tax: USD %.2f\nRatio: %.2f\n", earningsBeforeTax, profit, ratio)
}
