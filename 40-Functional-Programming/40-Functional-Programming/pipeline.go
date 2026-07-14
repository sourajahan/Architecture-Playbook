package main

import (
	"fmt"
)

// Map transforms a slice of type T to a slice of type R using the provided function.
// This is a "Pure Function" - it does not modify the input.
func Map[T any, R any](input []T, transform func(T) R) []R {
	result := make([]R, len(input))
	for i, v := range input {
		result[i] = transform(v)
	}
	return result
}

// Filter keeps only the elements that satisfy the predicate.
func Filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Data-in-Motion: We never mutate 'orders'. We create new views of the data.
	orders := []float64{10.0, 50.0, 100.0, 25.0}

	// Declarative Pipeline: Define "What" we want, not "How" to loop.
	highValueOrders := Filter(orders, func(val float64) bool {
		return val > 40.0
	})

	discountedPrices := Map(highValueOrders, func(val float64) float64 {
		return val * 0.9 // 10% discount
	})

	fmt.Printf("Original: %v\n", orders)
	fmt.Printf("Processed: %v\n", discountedPrices)
}
