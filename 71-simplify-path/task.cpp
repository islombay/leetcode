#include <iostream>
#include <stack>
#include <vector>
#include <sstream>

using namespace std;

class Solution
{
public:
    string simplifyPath(string path)
    {
        vector<string> pathArr = split(path, '/');

        Node *stack = nullptr;

        for (string ch : pathArr)
        {
            if (ch == "" || ch == ".")
            {
                continue;
            }

            if (ch == "..")
            {
                if (stack != nullptr)
                {
                    stack->pop();
                    continue;
                }
            }

            stack->push(ch);
        }

        path = "";
    }

private:
    class Node
    {
        string val;
        Node *head;
        Node *next;

        Node()
        {
            val = "";
            head = nullptr;
            next = nullptr;
        }

        Node(string val, Node *next)
        {
            val = val;
            head = nullptr;
            next = next;
        }

    public:
        void push(string path)
        {
            if (head == nullptr)
            {
                head = new Node(path, nullptr);
            }
            else
            {
                head = new Node(path, head->next);
            }
        }

        void pop()
        {
            if (head != nullptr)
            {
                head = head->next;
            }
        }

        string top()
        {
            return head->val;
        }
    };

    vector<string> split(const string &path, char delimiter)
    {
        vector<string> result;
        stringstream ss(path);
        string item;

        while (getline(ss, item, delimiter))
        {
            result.push_back(item);
        }
        return result;
    }
};