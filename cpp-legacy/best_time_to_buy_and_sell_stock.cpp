//  https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
#include <algorithm>
#include <vector>

/*
1.
make finalProfit and finalBuy
iterate
  if the finalProfit is greater than if the item is sold at currentPrice ==> finalProfit = currentProfit
  if the current Price is less ==> finalBuy = curBuy
  */

class Solution {
public:
  int maxProfit(std::vector<int> &prices) {
    int finalProfit = 0;
    int finalBuy = prices[0];
    for (int i = 1; i < prices.size(); i++) {
      finalProfit = std::max(finalProfit, prices[i] - finalBuy);
      finalBuy = std::min(finalBuy, prices[i]);// since this is updated later it doesn't affect the profit
    }
    return finalProfit;
  }
};

// main function
int main(int argc, char *argv[]) {
  Solution s;
  std::vector<int> prices = {7, 1, 5, 3, 6, 4};
  s.maxProfit(prices);
  return 0;
}
