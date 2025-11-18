/*
 * @lc app=leetcode id=637 lang=cpp
 *
 * [637] Average of Levels in Binary Tree
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
    vector<double> averageOfLevels(TreeNode* root) {
       if(root == NULL)
            return vector<double>{};
        
        vector<double> output;
        queue<TreeNode*> q;

        q.push(root);

        while(!q.empty()){

            int size = q.size();
            double total = 0;
            int num = size;
            while(size--){
                TreeNode *curr = q.front();
                q.pop();
                total += curr->val;
                if(curr->left != NULL){
                    q.push(curr->left);
                }
                if(curr->right != NULL){
                    q.push(curr->right);
                }
            }
            output.push_back(total/num);
        }

        return output;
         
    }
};
// @lc code=end

