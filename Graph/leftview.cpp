vector<int> leftView(Node *root)
{
   vector<int> leftview;
    if(root == NULL)
        return leftview;
    queue<Node *> A;
    A.push(root);
    A.push(NULL);
    
    Node *prev;
    while(A.empty() == false)
    {
        Node *curr = A.front();
        
        if(curr != NULL)
        {
            prev = A.front();
            if(curr->right != NULL)
                A.push(curr->right);
            if(curr->left != NULL)
                A.push(curr->left);
            
            A.pop();
            
        }
        else
        {
            A.pop();
            if(A.size() > 0)
                A.push(NULL);
            leftview.push_back(prev->data);
        }
        
        
    }
    
    return leftview;
}