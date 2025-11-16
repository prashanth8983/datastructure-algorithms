#include <bits/stdc++.h>

using namespace std;

int knapsack(vector<int> wt, vector<int> val, int size)
{
    int m = size;
    int n = wt.size();

    int T[n + 1][m + 1];
    for (int i = 0; i <= n; i++)
    {
        for (int j = 0; j <= m; j++)
        {
            if (i == 0 || j == 0)
                T[i][j] = 0;
            else if (wt[i - 1] <= j)
                T[i][j] = max(T[i - 1][j], val[i - 1] + T[i - 1][j - wt[i - 1]]);
            else
                T[i][j] = T[i - 1][j];
        }
    }
    return T[n][m];
}

int main()
{
    int n;
    cin >> n;
    int cap;
    cin >> cap;
    vector<int> wt(n);
    vector<int> val(n);
    for (auto &x : wt)
        cin >> x;

    for (auto &x : val)
        cin >> x;
    cout << knapsack(wt, val, cap);
    return 0;
}