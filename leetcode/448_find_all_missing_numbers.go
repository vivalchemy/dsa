package leetcode

// counting sort
func FindDisappearedNumbers(nums []int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}
	solution := make([]int, 0, 1)
	var i int
	for i < numsLen {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	for index, element := range nums {
		if index+1 != element {
			solution = append(solution, index+1)
		}
	}
	return solution
}

// negate the first occurence of the elements and then find the index of the non negated
func FindDisappearedNumbersUsingNegation(nums []int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}
	var soln []int
	for _, v := range nums {
		index := abs(v) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	for index, element := range nums {
		if element > 0 {
			soln = append(soln, index+1)
		}
	}
	return soln
}

// helper functions
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FindDisappearedNumbersUsingSubstitutionAndGrouping(nums []int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}
	soln := make([]int, numsLen)
	// get all the values of nums
	for _, value := range nums {
		soln[value-1] = value
	}

	var j, index int
	// can't use range due to inplace editing
	for ; index < numsLen; index++ {
		if soln[index] == 0 {
			soln[j] = index + 1
			j++
		}
	}
	return soln[:j]
}
