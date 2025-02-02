#include <algorithm>
#include <iostream>
#include <unordered_map>
#include <unordered_set>
#include <vector>

class Solution {
public:
  void printSet(std::unordered_set<char> &set) {
    for (char c : set) {
      std::cout << c << " ";
    }
  }

  void printVector(std::vector<char> &set) {
    for (char c : set) {
      std::cout << c << " ";
    }
  }

  int lengthOfLongestSubstring(std::string s) {
    std::unordered_map<char, int> umap;
    int maxSubstringLength = 0, start = 0, end;
    for (end = 0; end < s.length(); end++) {
      if (umap.find(s[end]) != umap.end() && umap[s[end]] >= start) {
        maxSubstringLength = std::max(maxSubstringLength, end - start);
        start = umap[s[end]] + 1;// where the current substring starts
      }
      umap[s[end]] = end; // last visited index of that character
    }
    return std::max(maxSubstringLength, end - start);
  }

  int lengthOfLongestSubstringUsingSet(std::string s) {
    std::unordered_set<char> set;
    std::vector<char> helper;
    int final = 0, deletionCounter = 0;
    for (char c : s) {
      if (set.count(c)) {
        final = set.size() > final ? set.size() : final;
        for (char i : helper) {
          set.erase(i);
          deletionCounter++;
          if (i == c) {
            break;
          }
        }
        // you shouldn't update vector while it i being traversed
        helper.erase(helper.begin(), helper.begin() + deletionCounter);
        deletionCounter = 0;
      }
      set.insert(c);
      helper.push_back(c);
    }
    final = set.size() > final ? set.size() : final;
    return final;
  }
};

// main function
int main(int argc, char *argv[]) {
  Solution s;
  std::cout << s.lengthOfLongestSubstring(" ");
  return 0;
}
