package leetcode

// import (
// 	"fmt"
// 	"time"
// )

// inefficient solution just laying out the idea with rough code, two pointer approach
// func SortedSquares(nums []int) []int {
// 	n := len(nums)
// 	soln := make([]int, n)
// 	frontPointer, backPointer := 0, n-1
// 	solnPointer := n - 1
// 	for frontPointer <= backPointer {
// 		// time.Sleep(time.Millisecond * 500)
// 		frontNum, backNum := nums[frontPointer], nums[backPointer]
// 		// fmt.Println("Front Pointer", frontPointer, "\tBack Pointer", backPointer, "\tFront Num", frontNum, "\tBack Num", backNum, "\tSoln Pointer", solnPointer)
// 		if frontNum < 0 {
// 			frontNum *= -1
// 		}
// 		if backNum < 0 {
// 			backNum *= -1
// 		}
// 		if frontPointer == backPointer {
// 			soln[solnPointer] = frontNum
// 			solnPointer--
// 			break
// 		}
// 		if frontNum < backNum {
// 			soln[solnPointer] = backNum
// 			solnPointer--
// 			backPointer--
// 		} else if frontNum > backNum {
// 			soln[solnPointer] = frontNum
// 			solnPointer--
// 			frontPointer++
// 		} else {
// 			soln[solnPointer], soln[solnPointer-1] = frontNum, backNum
// 			solnPointer -= 2
// 			frontPointer++
// 			backPointer--
// 		}
// 		// fmt.Println("Solution\n", soln)
// 	}
// 	for index := range soln {
// 		soln[index] = soln[index] * soln[index]
// 	}
// 	return soln
// }

// efficient solution by tiding things up
func SortedSquares(nums []int) []int {
	n := len(nums)
	soln := make([]int, n)
	frontPointer, backPointer := 0, n-1

	for solnIndex := n - 1; solnIndex >= 0; solnIndex-- {
		if nums[frontPointer]*nums[frontPointer] < nums[backPointer]*nums[backPointer] {
			soln[solnIndex] = nums[backPointer] * nums[backPointer]
			backPointer--
		} else {
			soln[solnIndex] = nums[frontPointer] * nums[frontPointer]
			frontPointer++
		}
	}
	return soln
}
