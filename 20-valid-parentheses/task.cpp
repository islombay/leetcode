#include <string>
#include <stack>

using namespace std;

class Solution {
public:
    bool isValid(string s) {
        stack<char> st;
        for (char ch : s) {
            if (ch == '(' || ch == '{' || ch == '[') {
                st.push(ch);
                continue;
            }

            if (st.empty()) {
                return false;
            }

            char top = st.top();

            if (top == '(' && ch == ')' || top == '[' && ch == ']' || top == '{' && ch == '}') {
                st.pop();
            } else {
                return false;
            }
        }

        return st.empty();
    }
};