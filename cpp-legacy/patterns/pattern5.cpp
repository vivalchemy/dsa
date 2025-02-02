
/*
 */
#include <iostream>

void pattern5(int num) {
  int count = 1;
  for (int i = 0; i < num; i++) {
    for (int j = 0; j <= i; j++, count++) {
      std::cout << count << " ";
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
