# include <bits/stdc++.h>

using namespace std;

struct Node{
    Node *left;
    Node *right;

    int data;

    Node()
    {
        data = 0;
        left = NULL;
        right = NULL;
    }

    Node(int item)
    {
        data = item;
        left = NULL;
        right = NULL;
    }
};


void inorder(Node* root)
{
    if(root != NULL)
    {
        inorder(root->left);
        cout<<root->data<<" ";
        inorder(root->right);
    }
}

int main()
{
    Node *root;

    root = new Node(10);
    root->left = new Node(20);
    root->left->left = new Node(40);
    root->left->right = new Node(50);
    root->left->right->left = new Node(70);
    root->left->right->right = new Node(80);
    root->right = new Node(30);
    root->right->right = new Node(60);

    inorder(root);

    return 0;
}