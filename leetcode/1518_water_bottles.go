package leetcode

func NumWaterBottles(numBottles int, numExchange int) int {
	var numDrankBottles int
	var excessEmptyBottles int
	for numBottles+excessEmptyBottles >= numExchange {
		numDrankBottles += numBottles
		numBottles, excessEmptyBottles = (numBottles+excessEmptyBottles)/numExchange, (numBottles+excessEmptyBottles)%numExchange
	}
	numDrankBottles += numBottles
	return numDrankBottles
}
