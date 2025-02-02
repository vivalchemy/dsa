#include <iostream>

void sayDigits(int n) {
  if (n == 0) {
    return;
  }
  sayDigits(n / 10);
  switch (n % 10) {
  case 0:
    std::cout << "Zero ";
    break;
  case 1:
    std::cout << "One ";
    break;
  case 2:
    std::cout << "Two";
    break;
  case 3:
    std::cout << "Three ";
    break;
  case 4:
    std::cout << "Four ";
    break;
  case 5:
    std::cout << "Five ";
    break;
  case 6:
    std::cout << "Six ";
    break;
  case 7:
    std::cout << "Seven ";
    break;
  case 8:
    std::cout << "Eight ";
    break;
  case 9:
    std::cout << "Nine ";
    break;
  }
}

int main() {
  sayDigits(543);
  return 0;
}
