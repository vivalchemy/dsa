// https://leetcode.com/problems/palindrome-number/description/
#include <iostream>
/*
 1.
 convert the integer to string
 use forloop two pointer while head < tail
   if not same return false
 return true

 2.
 extract all the digits by mod 10 and store them in a vector(since we don't know
 the size of the number,can't use array)
 use forloop two pointer while head < tail
   if not same return false
 return true

3.
reverse the number and then compare

*/
class Solution {
public:
  bool isPalindrome(int x) {
    // negative numbers and numbers divisible by 10
    if (x < 0 || x != 0 && x % 10 == 0) {
      // the or statement is required cause 
      std::cout << 0;
      return false;
    }
    if (x < 10) {
      std::cout << 1;
      return true;
    }
    int rev = 0;
    while (x > rev) {
      // first iteration: rev=0 ==> only x get inside
      rev = rev * 10 + x % 10;
      x /= 10;
    }
    // rev == x of even and rev / 10 == x for odd no digits as it gets rid of
    // the middle no stored as last digit in rev
    std::cout << (rev == x) || (rev == x / 10);
    return ((x == rev) || (x == rev / 10));
  }
};

// main function
int main(int argc, char *argv[]) {
  Solution s;
  s.isPalindrome(121);
  return 0;
}
