/*
 * @lc app=leetcode id=515 lang=cpp
 *
 * [515] Find Largest Value in Each Tree Row
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
    vector<int> largestValues(TreeNode* root) {
        if(root == NULL)
            return vector<int>{};
        
        vector<int> output;
        queue<TreeNode*> q;

        q.push(root);

        while(!q.empty()){

            int size = q.size();
            int max_num = INT_MIN;
            while(size--){
                TreeNode *curr = q.front();
                q.pop();
                max_num = max(max_num, curr->val);
                if(curr->left != NULL){
                    q.push(curr->left);
                }
                if(curr->right != NULL){
                    q.push(curr->right);
                }
            }
            output.push_back(max_num);
        }

        return output;
    }
};
// @lc code=end

