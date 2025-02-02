#include <unordered_set>
#include <vector>

class Solution {
public:
  bool containsDuplicate(std::vector<int> &nums) {
    std::unordered_set<int> uset;
    for (int i : nums) {
      if (uset.find(i) != uset.end()) {
        return true;
      } else {
        uset.insert(i);
      }
    }
    return false;
  }
};

// main function
int main(int argc, char *argv[]) {
  Solution s;
  std::vector<int> v = {1, 2, 3, 1};
  s.containsDuplicate(v);
  return 0;
}
