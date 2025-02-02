#include <iostream>

void pattern4(int num) {
  for (int i = 1; i <= num; i++) {
    for (int j = 0; j < i; j++) {
      std::cout << i << " ";
    }
    std::cout << std::endl;
  }
}

int main(int argc, char *argv[]) {
  int num;
  std::cout << "Enter the number of rows";
  std::cin >> num;
  pattern4(num);
  return 0;
}
