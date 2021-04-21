#include <bits/stdc++.h>

using namespace std;

int main()
{
    int n;
    cin >> n;
    vector<int> A(n);
    for (auto &x : A)
        cin >> x;

    unordered_map<int, int> val;

    for (int i = 0; i < n; i++)
    {
        for (int j = i + 1; j < n; j++)
        {
            if (A[i] > A[j])
                val[A[j]]++;
        }
    }
    for (auto x : val)
        cout << x.first << " " << x.second << "\n";

    return 0;
}