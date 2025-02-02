


/*
 */
#include <iostream>

void pattern5(int num) {
  for (int i = 0; i < num; i++) {
    for (int j = 0; j < num; j++) {
      char c = 'A'+i ;
      std::cout << c << " ";
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
