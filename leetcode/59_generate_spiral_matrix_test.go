package leetcode

import (
"testing"
  "reflect"
)

func GenerateSpiralMatrix(n int)[][]int {
  // generate a matrix of size n*n
  soln := make([][]int, n)
  for i := range soln {
    soln[i] = make([]int, n)
  }

  counter := 1

  left := 0
  right := n - 1
  top := 0
  bottom := n - 1


  // PrintMatrix(matrix)

  for top <= bottom && left <= right{

    // go from left to right
    // top ++
    for i := left; i <= right; i++{
      soln[top][i] = counter
      counter++
    }
    top++

    // go from top to bottom
    // right --
    for i := top; i <= bottom; i++{
      soln[i][right] = counter
      counter++
      // fmt.Print(matrix[i][right])
    }
    right--

    // go from right to left
    // bottom --
    if top <= bottom {

      for i := right; i >= left; i--{
        soln[bottom][i] = counter
        counter++
      }
      bottom--
    }

    // go from bottom to top
    // left ++
    if left <= right{
      for i := bottom; i >= top; i--{
        soln[i][left] = counter
        counter++
      }
      left++
    }
  }

  return soln
}

func TestGenerateSpiralMatrix(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected [][]int
	}{
		{
			name: "3x3 matrix",
			n:    3,
			expected: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "1x1 matrix",
			n:    1,
			expected: [][]int{
				{1},
			},
		},
		{
			name: "4x4 matrix",
			n:    4,
			expected: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateSpiralMatrix(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

