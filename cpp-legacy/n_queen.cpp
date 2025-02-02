#include <cstdlib>
#include <iostream>

#define N 5 // number of queens
#define ColorOff "\033[0m"
#define CSuccess "\033[0;32m"
#define CError "\033[0;31m"

void printBoard(int placedQueens[N]) {
  std::cout << "Solution :  " << std::endl;
  for (int i = 0; i < N; i++) {
    for (int j = 0; j < N; j++) {
      if (j == placedQueens[i]) {
        std::cout << CSuccess << "󰡚 " << ColorOff;
        continue;
      }
      std::cout << CError << "x " << ColorOff;
    }
    std::cout << std::endl;
  }
  std::cout << std::endl;
  return;
}

bool checkPos(int placedQueens[N], int curRow, int curCol) {
  for (int i = 0; i < curRow; i++) {
    // abs cause the direction of the diagonal is X so can be ++ +- -+ --
    if (placedQueens[i] == curCol ||
        abs(curRow - i) == abs(curCol - placedQueens[i])) {
      return false;
    }
  }
  return true;
}

// check the valid queens array
void nQueen(int placedQueens[N], int curRow) {
  // print the n queens if the it is the last row
  if (curRow == N) {
    printBoard(placedQueens);
  }
  for (int i = 0; i < N; i++) {
    if (checkPos(placedQueens, curRow, i)) {
      // if we place the queen in the current row
      placedQueens[curRow] = i;
      nQueen(placedQueens, curRow + 1); // go for the next row
    }
  }
}

int main() {
  int placedQueens[N] = {0}; // keeps the position of the placed queens
  nQueen(placedQueens, 0);
  return 0;
}
