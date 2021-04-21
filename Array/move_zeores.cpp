#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    vector<int> A(n);
    for (auto &x : A)
        cin >> x;

    int count = 0;
    for (int i = 0; i < n; i++)
    {
        if (A[i] != 0)
        {
            swap(A[i], A[count]);
            count++;
        }
            
    }

    for (auto x : A)
        cout << x << " ";

    return 0;
}
