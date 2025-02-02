// https://leetcode.com/problems/maximum-subarray/description/
#include <algorithm>// std::max()
#include <climits>// INT_MIN
#include <iostream>
#include <vector>

class Solution {
public:
  int maxSubArray(std::vector<int> &nums) {
    int maxSum = nums[0];
    int curSum = 0;// nums[0] not written cause for single input curSum becomes double

    for (int i : nums) {
      curSum = std::max(i, curSum + i);// if the current sum is less than 0 then i will become the cur sum
      maxSum = std::max(maxSum, curSum);// if the curSum surpasses the max sum the replace it 
    }
    std::cout<< maxSum;
    return maxSum;
  }
};

// main function
int main(int argc, char *argv[]) {
  Solution s;
  std::vector<int> nums = {1};
  s.maxSubArray(nums);
  return 0;
}
