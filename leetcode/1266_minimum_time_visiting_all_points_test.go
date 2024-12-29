package leetcode

import (
  "testing"
  "math"
)

func MinimunTimeVisitingAllPoints(points [][]int)int {
  hop := 0
  // assuming that the starting point is the first point in the array
  currentPoint := points[0]
  for _, point := range points{
    hop += int(
      math.Max(
        math.Abs(float64(currentPoint[0]) - float64(point[0])),
        math.Abs(float64(currentPoint[1]) - float64(point[1]))))
    currentPoint = point
  }
  return hop
}

func TestMinimunTimeVisitingAllPoints(t *testing.T) {
	tests := []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{1, 1}, {3, 4}, {-1, 0}}, 7},
		{[][]int{{3, 2}, {-2, 2}}, 5},
		{[][]int{{0, 0}, {1, 1}, {1, 2}}, 2},
		{[][]int{{0, 0}, {0, 0}}, 0},
	}

	for _, test := range tests {
		result := MinimunTimeVisitingAllPoints(test.points)
		if result != test.expected {
			t.Errorf("For points %v, expected %d but got %d", test.points, test.expected, result)
		}
	}
}
