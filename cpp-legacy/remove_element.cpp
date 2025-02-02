// https://leetcode.com/problems/remove-element/description/
#include <iostream>
#include <vector>

class Solution {
public:
  int removeElement(std::vector<int> &nums, int val) {
    int valCount = 0;        // the amount of times val occurs
    // int k = nums.size() - 1; // the place where val will be shifted
    // while (nums[k] == val) {
      // k--;
      // valCount++;
    // }
    for (int i = 0; i < nums.size(); i++) {
      if (nums[i] != val) {
        // no swap needed since we already have the value in val
        // nums[i] = nums[valCount];
        nums[valCount] = nums[i];
        nums[i] = val;
        // k--;
        valCount++;
      }
    }
    // std::cout << nums.size() - valCount << " " << k +1 << "\n";
    // for (int element : nums) {
    //   std::cout << element << " ";
    // }
    // return nums.size() - valCount;
    return valCount;
  }
};

// main function
int main(int argc, char *argv[]) {
  Solution s;
  // std::vector<int> nums = {3, 2, 2, 2, 3, 3};
  std::vector<int> nums = {};
  int val = 0;
  s.removeElement(nums, val);
  return 0;
}
