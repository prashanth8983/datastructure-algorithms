/*
 * @lc app=leetcode id=103 lang=cpp
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        
        if(root == nullptr) return vector<vector<int>>{};

        queue<TreeNode*> q;
        vector<vector<int>> result;

        q.push(root);

        int direction = 0;

        while(!q.empty()){
            int size = q.size();
            vector<int> level;
            while(size--){
                TreeNode *curr = q.front();
                q.pop();

                level.push_back(curr->val);

                if(curr->left != nullptr)
                    q.push(curr->left);
                if(curr->right != nullptr)
                    q.push(curr->right);
            }

            if(direction== 1){
                reverse(level.begin(),level.end());
            }
             direction = 1 - direction;

            result.push_back(level);
        }
        return result;
    }
};
// @lc code=end
