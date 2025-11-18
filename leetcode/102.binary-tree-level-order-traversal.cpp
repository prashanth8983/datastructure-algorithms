/*
 * @lc app=leetcode id=102 lang=cpp
 *
 * [102] Binary Tree Level Order Traversal
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
    vector<vector<int>> levelOrder(TreeNode* root) {
        
        if(root == NULL)
            return vector<vector<int>>{};
        
        vector<vector<int>> output;
        queue<TreeNode*> q;

        q.push(root);

        while(!q.empty()){

            int size = q.size();
            vector<int> curr_level;
            while(size--){
                TreeNode *curr = q.front();
                q.pop();
                curr_level.push_back(curr->val);
                if(curr->left != NULL){
                    q.push(curr->left);
                }
                if(curr->right != NULL){
                    q.push(curr->right);
                }
            }
            output.push_back(curr_level);
        }

        return output;
        
    }
};
// @lc code=end

