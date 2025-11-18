/*
 * @lc app=leetcode id=107 lang=cpp
 *
 * [107] Binary Tree Level Order Traversal II
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
class Solution
{
public:
    vector<vector<int>> levelOrderBottom(TreeNode *root)
    {

        if (root == NULL) return vector<vector<int>>{};

        queue<TreeNode *> q;
        vector<vector<int>> result;

        q.push(root);
        int size = 0;
        while (!q.empty())
        {
            size = q.size();
            vector<int> level;
            while(size--){

                TreeNode *curr = q.front();
                q.pop();

                level.push_back(curr->val);
                
                if(curr->left !=NULL)
                    q.push(curr->left);
                if(curr->right != NULL)
                    q.push(curr->right);
                
            }
            result.push_back(level);
        }

        reverse(result.begin(), result.end());

        return result;
    }
};
// @lc code=end


/*
================================================================================
                    INTERVIEW QUESTIONS AND SOLUTIONS
================================================================================

Q1: How would you modify this solution to return the level order traversal from
    top to bottom (normal order) instead of bottom-up?

Answer:
Simply remove the reverse() call at line 52. The BFS naturally produces
top-to-bottom order.

vector<vector<int>> levelOrder(TreeNode* root) {
    // Same code as above but without reverse()
    if (root == nullptr) return {};

    queue<TreeNode*> q;
    vector<vector<int>> result;
    q.push(root);

    while (!q.empty()) {
        int size = q.size();
        vector<int> level;

        while(size--) {
            TreeNode* curr = q.front();
            q.pop();
            level.push_back(curr->val);

            if(curr->left) q.push(curr->left);
            if(curr->right) q.push(curr->right);
        }
        result.push_back(level);
    }

    return result; // No reverse needed
}

--------------------------------------------------------------------------------

Q2: Can you solve this problem using DFS (Depth-First Search) instead of BFS?

Answer:
Yes! We can use DFS with a level parameter to track depth. This approach uses
recursion and builds levels as we traverse.

class Solution {
public:
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int>> result;
        dfs(root, 0, result);
        reverse(result.begin(), result.end());
        return result;
    }

private:
    void dfs(TreeNode* node, int level, vector<vector<int>>& result) {
        if (!node) return;

        // Ensure we have a vector for this level
        if (level >= result.size()) {
            result.push_back(vector<int>());
        }

        result[level].push_back(node->val);
        dfs(node->left, level + 1, result);
        dfs(node->right, level + 1, result);
    }
};

Time: O(n), Space: O(h) where h is height of tree (recursion stack)

--------------------------------------------------------------------------------

Q3: What if we want to return a zigzag level order traversal (alternating
    left-to-right and right-to-left)?

Answer:
We can add a flag to track direction and reverse alternate levels:

vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
    if (!root) return {};

    queue<TreeNode*> q;
    vector<vector<int>> result;
    q.push(root);
    bool leftToRight = true;

    while (!q.empty()) {
        int size = q.size();
        vector<int> level(size); // Pre-allocate size

        for (int i = 0; i < size; i++) {
            TreeNode* curr = q.front();
            q.pop();

            // Place value based on direction
            int index = leftToRight ? i : (size - 1 - i);
            level[index] = curr->val;

            if(curr->left) q.push(curr->left);
            if(curr->right) q.push(curr->right);
        }

        result.push_back(level);
        leftToRight = !leftToRight; // Toggle direction
    }

    return result;
}

--------------------------------------------------------------------------------

Q4: How would you find the maximum value in each level of the tree?

Answer:
Modify the level processing to track the maximum instead of all values:

vector<int> largestValues(TreeNode* root) {
    if (!root) return {};

    queue<TreeNode*> q;
    vector<int> result;
    q.push(root);

    while (!q.empty()) {
        int size = q.size();
        int maxVal = INT_MIN;

        while(size--) {
            TreeNode* curr = q.front();
            q.pop();
            maxVal = max(maxVal, curr->val);

            if(curr->left) q.push(curr->left);
            if(curr->right) q.push(curr->right);
        }

        result.push_back(maxVal);
    }

    return result;
}

--------------------------------------------------------------------------------

Q5: Can you find the rightmost node at each level (right side view)?

Answer:
Keep track of the last node processed at each level:

vector<int> rightSideView(TreeNode* root) {
    if (!root) return {};

    queue<TreeNode*> q;
    vector<int> result;
    q.push(root);

    while (!q.empty()) {
        int size = q.size();

        for (int i = 0; i < size; i++) {
            TreeNode* curr = q.front();
            q.pop();

            // Only add the rightmost node (last in level)
            if (i == size - 1) {
                result.push_back(curr->val);
            }

            if(curr->left) q.push(curr->left);
            if(curr->right) q.push(curr->right);
        }
    }

    return result;
}

--------------------------------------------------------------------------------

Q6: What's the space complexity difference between BFS and DFS approaches?

Answer:
- BFS Space: O(w) where w is the maximum width of the tree
  In worst case (complete tree), the last level has n/2 nodes, so O(n)

- DFS Space: O(h) where h is the height of the tree
  In worst case (skewed tree), h = n, so O(n)
  In best case (balanced tree), h = log(n), so O(log n)

For a balanced tree, DFS is more space-efficient.
For a very wide tree, DFS might be preferable.

--------------------------------------------------------------------------------

Q7: How would you connect nodes at the same level with next pointers?

Answer:
Use level order traversal and link nodes within each level:

class Node {
    int val;
    Node* left;
    Node* right;
    Node* next;
};

Node* connect(Node* root) {
    if (!root) return nullptr;

    queue<Node*> q;
    q.push(root);

    while (!q.empty()) {
        int size = q.size();
        Node* prev = nullptr;

        for (int i = 0; i < size; i++) {
            Node* curr = q.front();
            q.pop();

            if (prev) {
                prev->next = curr;
            }
            prev = curr;

            if(curr->left) q.push(curr->left);
            if(curr->right) q.push(curr->right);
        }
        // Last node in level points to nullptr (default)
    }

    return root;
}

================================================================================
                            COMMON FOLLOW-UP QUESTIONS
================================================================================

1. How would you handle very large trees that don't fit in memory?
   - Use iterative deepening or external memory algorithms
   - Process tree in chunks using disk storage

2. Can you optimize the space usage for the bottom-up traversal?
   - Instead of reversing at the end, insert each level at the beginning
   - result.insert(result.begin(), level) - but this is O(n²) time
   - Or use deque and push_front() for O(1) insertion

3. How would you parallelize this algorithm?
   - Each level can be processed in parallel after the previous completes
   - Use parallel queue processing with synchronization barriers

4. What if nodes have parent pointers?
   - Could traverse upward from leaves
   - Find all leaves first, then work upward level by level

5. How to find the minimum depth of the tree using level order?
   - Return the level number when you first encounter a leaf node
   - This gives the shortest path from root to any leaf

*/
