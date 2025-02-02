
/*
 *123
 *456
 *789
*/
#include <iostream>
#include <ostream>

int main (int argc, char *argv[]) {
  // get the user input
  int num;
  std::cout<<"Enter the number of rows";
  std::cin>>num;
  // for loop to print the rows
  for (int i = 0; i < num; i++) {
  // for loop to print the columns
    for (int j = 1; j <= num; j++) {
    std::cout<<j+i*num;
    }
    std::cout<<std::endl;
  } 
  return 0;
}
