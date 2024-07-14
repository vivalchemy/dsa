package strivera2zdsa

func UnionOfArrays(nums1, nums2 []int) []int {
	var pointer1, pointer2 int
	nums1Len, nums2Len := len(nums1), len(nums2)
	soln := make([]int, 0, nums1Len+nums2Len)
	if nums1Len == 0 {
		return nums2
	} else if nums2Len == 0 {
		return nums1
	}
	for pointer1 < nums1Len && pointer2 < nums2Len {
		if nums1[pointer1] <= nums2[pointer2] {
			if len(soln) != 0 || soln[len(soln)-1] != nums2[pointer2] {
				soln = append(soln, nums1[pointer1])
			}
			pointer1++
		} else {
			if len(soln) != 0 || soln[len(soln)-1] != nums2[pointer2] {
				soln = append(soln, nums2[pointer2])
			}
			pointer2++
		}
	}

	if pointer1 < nums1Len {
		if soln[len(soln)-1] != nums2[pointer2] {
			soln = append(soln, nums1[pointer1])
		}
		pointer1++
	}

	if pointer2 < nums2Len {
		if soln[len(soln)-1] != nums2[pointer2] {
			soln = append(soln, soln[pointer2])
		}
		pointer2++
	}
	return soln
}
