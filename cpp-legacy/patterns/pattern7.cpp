

/*
 */
#include <iostream>

void pattern5(int num) {
  for (int i = 1; i <= num; i++) {
    for (int j = 0; j < i; j++) {
      std::cout << i - j << " ";
    }
    std::cout << std::endl;
  }
}

int main(int argc, char *argv[]) {
  int num;
  std::cout << "Enter the number of rows";
  std::cin >> num;
  pattern5(num);
  return 0;
}
