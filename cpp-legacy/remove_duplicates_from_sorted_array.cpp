#include <iostream>
#include <vector>

class Solution {
public:
  int removeDuplicates(std::vector<int> &nums) {
    int j = 1;
    for (int i = 1; i < nums.size(); i++) {
      if (nums[i] != nums[i - 1]) {
        nums[j] = nums[i];
        j++;
      }
    }
    return j;
  }
};

// main function
int main(int argc, char *argv[]) {
  Solution s;
  std::vector<int> nums = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
  s.removeDuplicates(nums);
  return 0;
}
