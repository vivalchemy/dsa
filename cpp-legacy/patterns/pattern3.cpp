#include <iostream>
#include <ostream>
/*
* *
* **
* ***
* ****
*/
void pattern3(int num){
  for(int i = 0; i < num; i++){
    for(int j = 0; j <= i; j++){
      std::cout<<"* ";
    }
    std::cout<<std::endl;
  }
}

int main (int argc, char *argv[]) {
  int num;
  std::cout << "Enter the number of rows";
  std::cin>>num;
  pattern3(num);
  return 0;
}
