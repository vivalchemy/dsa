

/*
1
2 3
3 4 5
4 5 6 7
5 6 7 8 9
6 7 8 9 10 11
*/
#include <iostream>

void pattern5(int num) {
  for (int i = 0; i < num; i++) {
    for (int j = 1; j <= i + 1; j++) {
      std::cout << j + i << " ";
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
