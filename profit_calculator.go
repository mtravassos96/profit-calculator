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

	fmt.Println("Earnings Before Tax: USD ", earningsBeforeTax)
	fmt.Println("Earnings After Tax: USD ", profit)
	fmt.Println("Ratio: ", ratio)
}
