package leetcode

// using for loop
// func PassThePillow(n int, t int) int {
// 	var towardsN bool = true
// 	curPilAt := 0
// 	for range t {
// 		time.Sleep(time.Millisecond * 0)
// 		fmt.Print(curPilAt, "\t")
// 		if towardsN {
// 			curPilAt++
// 		} else {
// 			curPilAt--
// 		}
// 		if curPilAt == 0 || curPilAt == n-1 {
// 			towardsN = !towardsN
// 			fmt.Println()
// 		}
// 	}
// 	return curPilAt + 1
// }

// using math
func PassThePillow(n int, time int) int {

	if time < n {
		return time + 1
	}
	noOfRotations := time / (n - 1)
	excess := time % (n - 1)
	if noOfRotations%2 == 0 {
		// fmt.Println("no of Rotations", noOfRotations, "\nexcess", excess)
		return excess + 1
	} else {
		// fmt.Println("no of Rotations", noOfRotations, "\nexcess n", excess, "\nn - excess", n-excess)
		return n - excess
	}
}

// wizardry in math
// func PassThePillow(n int, time int) int {
//
// 	if time < n {
// 		return time + 1
// 	}
// 	ans := time%((n-1)*2) + 1
// 	if ans < n {
// 		return ans
// 	}
// 	return n - time%(n-1)
//
// }
