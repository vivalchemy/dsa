package leetcode

import (
  "testing"
  "fmt"
)

func PrintMatrix(matrix [][]int) {
  // Find the maximum width needed for any number
  maxWidth := 0
  for _, row := range matrix {
    for _, value := range row {
      width := len(fmt.Sprintf("%d", value))
      if width > maxWidth {
        maxWidth = width
      }
    }
  }

  // Print the matrix with proper formatting
  for _, row := range matrix {
    for _, value := range row {
      fmt.Printf("%*d ", maxWidth, value)
    }
    fmt.Println()
  }
}

func SpiralMatrix(matrix [][]int)[]int {
  soln := []int{}
  rows := len(matrix)
  columns := len(matrix[0])

  left := 0
  right := columns - 1
  top := 0
  bottom := rows - 1


  // PrintMatrix(matrix)

  for top <= bottom && left <= right{

    // go from left to right
    // top ++
    for i := left; i <= right; i++{
      soln = append(soln,matrix[top][i])
    }
    top++

    // go from top to bottom
    // right --
    for i := top; i <= bottom; i++{
      soln = append(soln, matrix[i][right])
      // fmt.Print(matrix[i][right])
    }
    right--

    // go from right to left
    // bottom --
    if top <= bottom {

      for i := right; i >= left; i--{
        soln = append(soln, matrix[bottom][i])
      }
      bottom--
    }

    // go from bottom to top
    // left ++
    if left <= right{
      for i := bottom; i >= top; i--{
        soln = append(soln, matrix[i][left])
      }
      left++
    }
  }

  return soln
}

func TestSpiralMatrix(t *testing.T) {
  testCases := []struct {
    matrix   [][]int
    expected []int
  }{
    {
      matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
      expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
    },
    {
      matrix:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
      expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
    },
  }

  for _, tc := range testCases {
    result := SpiralMatrix(tc.matrix)
    if !equal(result, tc.expected) {
      t.Errorf("For matrix %v, expected %v, but got %v", tc.matrix, tc.expected, result)
    }
  }
}

// Helper function to compare two slices
func equal(a, b []int) bool {
  if len(a) != len(b) {
    return false
  }
  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}


/*
Reason for the if conditions:

* when the rows > columns then the first for loop will have an extra traversal
* when the columns > rows then the first and the second for loop will have an 
  extra traversal
* however if the rows and columns are not equal then the last two for loop 
  shouldn't be activated hence we check first before printing the values
*/
