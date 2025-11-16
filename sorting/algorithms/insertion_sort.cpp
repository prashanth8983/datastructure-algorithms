#include <bits/stdc++.h>

using namespace std;

void insertion_sort(vector<int> &A)
{
    int key;
    for(int i=1;i<A.size();i++)
    {
        key = A[i];
        int j = i-1;
        while(A[j] > key && j >= 0){
            A[j+1] = A[j];
            j--;
        }
        A[j+1] = key;
    }
}

int main()
{
    int n;
    cin >> n;
    vector<int> A(n);
    for (auto &x : A)
        cin >> x;

    insertion_sort(A);
    for(auto x: A)
        cout<<x<<" ";
    return 0;
}