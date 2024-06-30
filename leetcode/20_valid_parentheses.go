package leetcode

import "fmt"

// using switch case exhaustively
// func ValidParentheses(s string) bool {
// 	braces := make([]rune, 10000)
// 	var bracesPointer int
// 	for _, val := range s {
// 		fmt.Printf("%v %v %T %#v\n", string(val), val, val, val)
// 		switch val {
// 		case '{':
// 			// fmt.Println("{")
// 			braces[bracesPointer] = '{'
// 			bracesPointer++
// 		case '[':
// 			// fmt.Println("[")
// 			braces[bracesPointer] = '['
// 			bracesPointer++
// 		case '(':
// 			// fmt.Println("(")
// 			braces[bracesPointer] = '('
// 			bracesPointer++
// 		case '}':
// 			// fmt.Println("}")
// 			bracesPointer--
// 			if bracesPointer < 0 || braces[bracesPointer] != '{' {
// 				// fmt.Println("} got fucked")
// 				return false
// 			}
// 		case ']':
// 			// fmt.Println("]")
// 			bracesPointer--
// 			if bracesPointer < 0 || braces[bracesPointer] != '[' {
// 				// fmt.Printf("%v %#v %T", braces[bracesPointer], braces[bracesPointer], braces[bracesPointer])
// 				// fmt.Println("] got fucked")
// 				return false
// 			}
// 		case ')':
// 			// fmt.Println(") ", braces[bracesPointer], bracesPointer, braces)
// 			bracesPointer--
// 			if bracesPointer < 0 || braces[bracesPointer] != '(' {
// 				// fmt.Printf("%v %#v %T", braces[bracesPointer], braces[bracesPointer], braces[bracesPointer])
// 				// fmt.Println(") got fucked")
// 				return false
// 			}
// 		}
// 	}
// 	if bracesPointer != 0 {
// 		// fmt.Println("bracesPointer got fucked at ", bracesPointer)
// 		return false
// 	}
// 	return true
// }

// using maps
func ValidParentheses(s string) bool {
	stack := make([]rune, 0, 10000)
	pair := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, val := range s {
		fmt.Printf("%v %v %T %#v\n", string(val), val, val, val)
		switch val {
		case '(', '[', '{':
			stack = append(stack, val)
		case ')', ']', '}':
			noOfOpeners := len(stack)
			if noOfOpeners <= 0 || stack[noOfOpeners-1] != pair[val] {
				return false
			}
			stack = stack[:noOfOpeners-1]
		}
	}
	return len(stack) == 0
}
