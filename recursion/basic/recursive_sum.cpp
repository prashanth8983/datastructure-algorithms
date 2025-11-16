#include <bits/stdc++.h>

using namespace std;

int recursion_sum(int n)
{
    if(n == 0)
        return 0;
    return n+recursion_sum(n-1);
}

int main()
{
    int n;
    cin>>n;
    cout<<recursion_sum(n);
    return 0;
}