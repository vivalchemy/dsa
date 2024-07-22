package strivera2zdsa

import (
	"reflect"
	"testing"
)

func StockBuyAndSell(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	minPriceSoFar := nums[0]
	soln := 0
	for _, todaysPrice := range nums {
		if minPriceSoFar > todaysPrice {
			minPriceSoFar = todaysPrice
			continue
		}
		if todaysPrice-minPriceSoFar > soln {
			soln = todaysPrice - minPriceSoFar
		}
	}
	return soln
}

func TestStockBuyAndSell(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5}, // Buy at 1, sell at 6
		{[]int{7, 6, 4, 3, 1}, 0},    // No profit possible
		{[]int{1, 2, 3, 4, 5}, 4},    // Buy at 1, sell at 5
		{[]int{2, 1, 4, 3, 5}, 4},    // Buy at 1, sell at 5
		{[]int{5, 4, 3, 2, 1}, 0},    // No profit possible
		{[]int{3, 3, 3, 3, 3}, 0},    // No profit possible
		{[]int{1}, 0},                // Single day, no profit possible
		{[]int{}, // Empty array
			0},
		{[]int{1, 5, 3, 6, 4, 7, 1, 2}, 6}, // Buy at 1, sell at 7
	}

	for _, test := range tests {
		result := StockBuyAndSell(test.prices)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For prices %v, expected %d but got %d", test.prices, test.expected, result)
		}
	}
}
