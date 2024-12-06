#include <iostream>
#include <vector>
#include <map>

using namespace std;

class Solution {
public:
    int maxCount(vector<int>& banned, int n, int maxSum) {
        map<int, bool> bannedMap;
        for (int v : banned) {
            if (v < n) {
                bannedMap[v] = true;
            }
        }

        int totalSum = 0;
        int total = 0;

        for (int i = 1; i <= n; i++){
            if (bannedMap.find(i) != bannedMap.end()) {
                continue;
            }
            totalSum += i;
            if (totalSum > maxSum){
                return total;
            }
            total++;
        }
        return total;
    }
};

int main() {
    // Define test cases as a struct
    struct TestCase {
        int id;
        vector<int> banned;
        int n;
        int maxSum;
        int want;
    };

    // Initialize test cases
    vector<TestCase> cases = {
        {1, {1, 6, 5}, 5, 6, 2},
        {2, {1, 2, 3, 4, 5, 6, 7}, 8, 1, 0},
        {3, {11}, 7, 50, 7}
    };

    Solution solution;

    // Iterate over test cases
    for (auto& test : cases) {
        int res = solution.maxCount(test.banned, test.n, test.maxSum);
        if (res != test.want) {
            cout << "Test " << test.id << " failed: expected " << test.want 
                 << ", got " << res << endl;
        } else {
            cout << "Test " << test.id << " passed." << endl;
        }
    }

    return 0;
}