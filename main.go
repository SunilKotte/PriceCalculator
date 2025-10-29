package main

import (
	"sunilkotte/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	// var result map[float64][]float64 =make(map[float64][]float64)
	// result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		pricesJob := prices.NewTaxIncludedPriceJob(taxRate)
		pricesJob.Process()
	}
 
}
