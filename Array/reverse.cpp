# include <bits/stdc++.h>

using namespace std;

int main(){
    int n;
    cin>>n;
    vector<int> A(n);
    for(auto &x: A)
        cin>>x;
    for(int i=0;i<n/2;i++)
    {
        int temp = A[i];
        A[i] = A[n-i-1];
        A[n-i-1] = temp;
    }

    for(auto x: A)
        cout<<x<<" ";
    return 0;
}