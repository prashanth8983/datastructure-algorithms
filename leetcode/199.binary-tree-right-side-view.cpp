/*
 * @lc app=leetcode id=199 lang=cpp
 *
 * [199] Binary Tree Right Side View
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
    vector<int> rightSideView(TreeNode* root) {
        
        if(root==nullptr) return vector<int>{};

        vector<int> result;
        queue<TreeNode*> q;

        q.push(root);

        while(!q.empty()){
            int size = q.size();
            int pushed = false;
            while(size--){

                TreeNode *curr = q.front();

                if(!pushed){
                    result.push_back(q.back()->val);
                    pushed = true;
                }
                //
                q.pop();

                if(curr->left){
                    q.push(curr->left);
                }

                if(curr->right){
                    q.push(curr->right);
                }

            }
        }

        return result;


    }
};
// @lc code=end

