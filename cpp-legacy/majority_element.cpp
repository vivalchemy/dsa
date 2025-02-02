#include <iostream>
#include <unordered_map>
#include <vector>

class Solution {
public:
  int majorityElement(std::vector<int> &nums) {
    std::unordered_map<int, int> umap;
    int i;
    for (int i : nums) {
      if (umap[i] < nums.size() / 2) {
        umap[i]++;
      } else {
        return i;
      }
    }
    return 0;
  }
};

// main function
int main(int argc, char *argv[]) {
  Solution s;
  std::vector<int> nums = {3, 2, 3};
  s.majorityElement(nums);
  return 0;
}
