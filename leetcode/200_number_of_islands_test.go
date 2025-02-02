package leetcode

import "testing"

// The following solution won't work
// for e.g.:
// ["1","1","1","1","0"],
// ["1","1","0","1","0"],
// ["1","1","0","0","1"], <-- this is increment the numberOfIslands
// ["0","0","0","1","1"]
//               ^
//               |_ even though we don't want to increment at here it will 
//                  increment for the following solution
// func NumberOfIslands(grid [][]byte)int {
//   numberOfIslands := 0
//   for row, curRow := range(grid){
//     for col, val := range(curRow){
//       if val == '1' {
//         if (row == 0 || grid[row-1][col] != '1') && (col == 0 || grid[row][col-1] != '1') {
//           numberOfIslands++
//         }
//       }
//     }
//   }
//
//   return numberOfIslands
// }


func NumberOfIslands(grid [][]byte)int {
  numberOfIslands := 0
  return numberOfIslands
}

func TestNumberOfIslands(t *testing.T) {

}

