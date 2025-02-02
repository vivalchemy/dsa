#include <iostream>

int countStairs(int n) {
  if (n <= 0) {
    return 0;
  } else if (n == 1) {
    return 1;
  } else if (n == 2) {
    return 2;
  }
  return countStairs(n - 1) + countStairs(n - 2);
}

int main() {
  std::cout << countStairs(7);
  return 0;
}
