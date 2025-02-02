#include <cstdio>

int coins[] = {1, 2, 3}, sum = 4;
int numberofCoins = 3;
// Function to initialize 1st column of dynamicprogTable with 1
void initdynamicprogTable(int dynamicprogTable[][5]) {
  int i;
  // First row to 0
  for (i = 1; i <= sum + 1; i++) {
    dynamicprogTable[0][i] = 0;
  }
  // First column to 1
  for (i = 1; i <= numberofCoins; i++) {
    dynamicprogTable[i][0] = 1;
  }
}

int solution(int dynamicprogTable[][5]) {
  int coinindex, dynamicprogSum;
  for (coinindex = 1; coinindex < numberofCoins + 1; coinindex++) {
    for (dynamicprogSum = 1; dynamicprogSum < 5; dynamicprogSum++) {
      // value of coin should be less than or equal to sum value to consider it
      if (coins[coinindex - 1] > dynamicprogSum)
        dynamicprogTable[coinindex][dynamicprogSum] =
            dynamicprogTable[coinindex - 1][dynamicprogSum];
      else
        dynamicprogTable[coinindex][dynamicprogSum] =
            dynamicprogTable[coinindex - 1][dynamicprogSum] +
            dynamicprogTable[coinindex][dynamicprogSum - coins[coinindex - 1]];
    }
  }
  // return final row and column value
  return dynamicprogTable[numberofCoins][sum];
}

int main() {
  int dynamicprogTable[numberofCoins + 1][5];
  initdynamicprogTable(dynamicprogTable);
  printf("Total Solutions: %d", solution(dynamicprogTable));
  return 0;
}
