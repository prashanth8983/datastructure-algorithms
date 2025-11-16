# include <bits/stdc++.h>
using namespace std;

void selection_sort(vector<int> &A)
{
    int min = 0;
    for(int i=0;i < A.size()-1;i++){
        for(int j=0;j<A.size()-i-1;j++){
            if(A[j] < A[min])
                min = j;
        }
        swap(A[min],A[i]);
    }
}

int main(){
    int n;
    cin>>n;
    vector<int> A(n);
    for(auto &x: A)
        cin>>x;
    selection_sort(A);
    for(auto x: A)
        cout<<x<<" ";
    return 0;
}