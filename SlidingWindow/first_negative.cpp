#include <bits/stdc++.h>

using namespace std;

void first_negative(vector<int> A, int k)
{
    for(int i=0;i<A.size();i++)
        if(A[i] > 0)
            A[i] = 0;

    for(int i=0;i<A.size()-1;i++)
        A[i+1] += A[i];
    
    A.insert(A.begin(),0);
    for(int i=0;i<A.size()-k;i++)
    {
        if(A[i+k]-A[i]);
    }

}

int main()
{
    int n, k;
    cin >> n >> k;
    vector<int> A(n);
    for (auto &x : A)
        cin >> x;
    first_negative(A, k);
    return 0;
}