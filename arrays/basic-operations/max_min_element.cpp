#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin>>n;
    vector<int> A(n);
    int MAX,MIN;
    for(auto &x: A){
        cin>>x;
        MAX = max(MAX,x);
        MIN = min(MIN,x);
    }

    cout<<MAX<<" "<<MIN;
    return 0;
}