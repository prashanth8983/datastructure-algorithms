# include <bits/stdc++.h>

using namespace std;
int sliding_window(vector<int> A, int k)
{
    A.insert(A.begin(),0);
    for(int i=1;i<A.size()-1;i++)
        A[i+1]+=A[i];
    int sum = -1;
    for(int i=0;i<A.size()-k;i++)
        sum = max(sum,A[i+k]-A[i]);
    return sum;
}


int main(){

    int n,k;
    cin>>n>>k;
    vector<int> A(n);
    for(int &x : A)
        cin>>x;
    cout<<sliding_window(A,k);
    return 0;
}