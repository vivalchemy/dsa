package leetcode
//
// import (
// 	"fmt"
// 	"sort"
// 	"testing"
// )
//
// func NumberOfAtoms(formula string) string {
// 	var soln string
// 	eleMap := make(map[string]int)
// 	openBracsInStckCounter := 0
// 	stk := make([]byte, 0, len(formula))
// 	for i := 0; i < len(formula); i++ {
// 		element := formula[i]
// 		if element >= 'A' && element <= 'Z' {
// 			compEle := string(element)
// 			var count int
// 			for i+1 < len(formula) && formula[i+1] >= 'a' && formula[i+1] <= 'z' {
// 				i++
// 				compEle += string(formula[i])
// 			}
// 			for i+1 < len(formula) && formula[i+1] >= '0' && formula[i+1] <= '9' {
// 				i++
// 				count = count*10 + int(formula[i]-'0')
// 			}
// 			if count == 0 {
// 				count = 1
// 			}
// 			eleMap[compEle] += count
// 		}
// 		if element == '(' {
// 			openBracsInStckCounter++
// 			for element != ')' {
// 				if element == '(' {
// 					openBracsInStckCounter++
// 				}
// 				stk = append(stk, element)
// 				i++
// 				element = formula[i]
// 			}
// 		} else if element == ')' {
// 			openBracsInStckCounter--
// 			stk = append(stk, ')')
// 		}
// 	}
// 	elements := make([]string, 0, len(eleMap))
// 	for element := range eleMap {
// 		elements = append(elements, element)
// 	}
// 	sort.Strings(elements)
// 	for _, element := range elements {
// 		count := eleMap[element]
// 		soln += element
// 		if count > 1 {
// 			soln += fmt.Sprintf("%d", count)
// 		}
// 	}
// 	fmt.Println(soln)
// 	return soln
// }
//
// // works for non braces
// // func NumberOfAtoms(formula string) string {
// // 	var soln string
// // 	eleMap := make(map[string]int)
// // 	// stk := make([]byte, 0, len(formula))
// // 	for i := 0; i < len(formula); i++ {
// // 		element := formula[i]
// // 		if element >= 'A' && element <= 'Z' {
// // 			compEle := string(element)
// // 			var count int
// // 			for i+1 < len(formula) && formula[i+1] >= 'a' && formula[i+1] <= 'z' {
// // 				i++
// // 				compEle += string(formula[i])
// // 			}
// // 			for i+1 < len(formula) && formula[i+1] >= '0' && formula[i+1] <= '9' {
// // 				i++
// // 				count = count*10 + int(formula[i]-'0')
// // 			}
// // 			if count == 0 {
// // 				count = 1
// // 			}
// // 			eleMap[compEle] += count
// // 		}
// // 	}
// // 	elements := make([]string, 0, len(eleMap))
// // 	for element := range eleMap {
// // 		elements = append(elements, element)
// // 	}
// // 	sort.Strings(elements)
// // 	for _, element := range elements {
// // 		count := eleMap[element]
// // 		soln += element
// // 		if count > 1 {
// // 			soln += fmt.Sprintf("%d", count)
// // 		}
// // 	}
// // 	fmt.Println(soln)
// // 	return soln
// // }
//
// func TestNumberOfAtoms(t *testing.T) {
// 	tests := []struct {
// 		formula  string
// 		expected string
// 	}{
// 		{"H2O", "H2O"},
// 		{"H2O2", "H2O2"},
// 		{"H2O2He3Mg4", "H2He3Mg4O2"},
// 		{"(H2O2)", "H2O2"},
// 		{"(H2O2)3", "H6O6"},
// 		{"K4(ON(SO3)2)2", "K4N2O14S4"},
// 		{"Mg(OH)2", "H2MgO2"},
// 		{"Be32", "Be32"},
// 		{"Fe2O3", "Fe2O3"},
// 		{"(NH4)2SO4", "H8N2O4S"},
// 	}
//
// 	for _, test := range tests {
// 		if result := NumberOfAtoms(test.formula); result != test.expected {
// 			t.Errorf("For formula %s, expected %s but got %s", test.formula, test.expected, result)
// 		}
// 	}
// }
