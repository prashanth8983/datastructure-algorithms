#include <bits/stdc++.h>

using namespace std;

struct TreeNode
{

    TreeNode *left;
    TreeNode *right;
    int Data;

    TreeNode(int x)
    {
        Data = x;
        left = NULL;
        right = NULL;
    }
};

TreeNode *root;

void construct(string tree)
{
    if(tree == "")
        return;
    root = new TreeNode(tree[0]);
}

int main()
{
    string treestr;
    cin >> treestr;
    //4(2(3)(1))(6(5))

    return 0;
}