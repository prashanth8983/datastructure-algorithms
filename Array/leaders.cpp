#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    vector<int> A(n);
    for (auto &x : A)
        cin >> x;

    vector<int> leader;

    int current = A[n - 1];
    leader.push_back(current);
    for (int i = n - 2; i >= 0; i--)
    {
        if (A[i] > current)
        {
            leader.push_back(A[i]);
            current = A[i];
        }
    }

    for (auto x : leader)
        cout << x << " ";
    return 0;
}