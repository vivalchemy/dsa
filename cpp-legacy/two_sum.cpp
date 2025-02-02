// https://leetcode.com/problems/two-sum/description/
#include <iostream>
#include <unordered_map>
#include <vector>

int main() {

  std::vector<int> nums = {3, 3};
  int target = 6;
  std::unordered_map<int, int> umap; // <element,index>
  // iterate through all the elements in the array
  // solution one
  // for (int i = 0; i < nums.size(); i++) {
  //   // check if the element exists in the array
  //   if (umap.find(nums[i]) == umap.end()) {
  //     // std::cout<<"Element not found"
  //     // if element not found add its complement to the array
  //     umap[target - nums[i]] = i;
  //   } else {
  //     std::cout << "First element\t" << nums[i] << " at index " << i;
  //     std::cout << "\nSecond element\t" << target - nums[i] << " at index "
  //               << umap[nums[i]];
  //   }
  // }

  for (int i = 0; i < nums.size(); i++) {
    // check if the element exists in the array
    if (umap.count(nums[i])) {
      // if count non-zero then exist
      std::cout << "First element\t" << nums[i] << " at index " << i;
      std::cout << "\nSecond element\t" << target - nums[i] << " at index "
                << umap[nums[i]];

    } else {
      //if count zero then doesn't exist hence add 
      umap[target - nums[i]] = i;
    }
  }
}
