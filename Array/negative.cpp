#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin>>n;
    vector<int> A(n);
    for(auto &x: A)
        cin>>x;
    
    int mid = n/2;
    int low = 0;
    int high = n-1;

    while(low<= high){
        mid = (high+low)/2;
        if(A[mid] < 0)
            swap(A[low++],A[mid]);
        else if(A[mid] > 0)
            swap(A[high--],A[mid]);
    }

    for(int i=0;i<n;i++)
        cout<<A[i]<<" ";

    return 0;
}