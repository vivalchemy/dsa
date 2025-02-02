/*
 *1234
 *1234
 *1234 
 *1234
*/
#include <iostream>

int main (int argc, char *argv[]) {
  // get the user input
  int num;
  std::cout<<"Enter the number of rows";
  std::cin>>num;
  // for loop to print the rows
  for (int i = 1; i <= num; i++) {
  // for loop to print the columns
    for (int j = 1; j <= num; j++) {
    std::cout<<j;
    }
    std::cout<<std::endl;
  } 
  return 0;
}
