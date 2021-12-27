package palindrome

import (
	"fmt"
	"strconv"
)

type Product struct {
	Value   int
	Factorizations [][2]int // An array of 2-length-arrays of ints
}
// If the product exists, find it's index in the array
func ProductIndex(value int, products []Product) int {
	for i, product := range products {
		if product.Value == value {
			return i
		}
	}
	return -1
}
// Find if a multiplication factor already exists in array
func FactorInArray(factors [][2]int, num1, num2 int) bool{
	for _, factor := range factors {
		if [2]int{num2, num1} == factor { // If factor is in array, it must be the other permutation
			return true
		}
	}
	return false
}
// Append to products slice if palindrome
func appendPalindromes(num1, num2 int, products *[]Product) {
	value := num1 * num2
	if isPalindrome(value) {
		// 1. If product already exists in array, append factor to it
		// 	1a. Append factor only if it is unique - considering flipped permutations
		// 2. If product is new, append it to array
		existing_index := ProductIndex(value, *products)
		if existing_index != -1 { // Since product already exists in array, append factor to it
			if !FactorInArray((*products)[existing_index].Factorizations, num1, num2) { // Check if unique
				(*products)[existing_index].Factorizations = append((*products)[existing_index].Factorizations, [2]int{num1, num2})
			}
		} else { // Product is new, append it to array
			product := Product{
				Value:   value,
				Factorizations: [][2]int{{num1, num2}},
			}
			*products = append(*products, product)
		}
	}
}
// Get all products within a given range (inclusive)
func getProducts(fmin, fmax int) []Product {
	products := make([]Product, 0)
	var num1, num2 int
	// Start from left
	for num1 = fmin; num1 <= fmax; num1++ {
		for num2 = fmin; num2 <= fmax; num2++ {
			appendPalindromes(num1, num2, &products)
		}
	}
	return products
}
// Reverse a string - used for checking Palindromes
func reverseString(input string) string {
	var output string
	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}
	return output
}
// Check if an integer is palindrome via String reversal
func isPalindrome(value int) bool {
	if value < 10 { // Single digit is automatically palindrome
		return true
	}
	value_string := strconv.Itoa(value) // String representation
	return value_string == reverseString(value_string)
}
// Find min and max products from the completed slice
func findMinAndMax(products[]Product) (Product, Product) {
	// Initial guesses
	min_product := products[0]
	max_product := products[len(products)-1]
	// Update initial guesses
	for _, product := range(products) {
		if product.Value < min_product.Value {
			min_product = product
		}
		if product.Value > max_product.Value {
			max_product = product 
		}
	}
	return min_product, max_product
}
// Main function called from test cases
func Products(fmin, fmax int) (Product, Product, error) {
	// Error handling
	if fmax <= fmin {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax")
	}
	// Get all Palindrome Products and Factorizations
	palindrome_products := getProducts(fmin, fmax)
	if len(palindrome_products) != 0 {
		min_product, max_product := findMinAndMax(palindrome_products)
		return min_product, max_product, nil
	} else {
		return Product{}, Product{}, fmt.Errorf("no palindromes")
	}
}
