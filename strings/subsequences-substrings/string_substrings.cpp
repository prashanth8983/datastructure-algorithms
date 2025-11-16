#include <bits/stdc++.h>

using namespace std;

void compute(vector<string> &res, string s, int i)
{
    if (i == s.length())
    {
        res.push_back(s);
        return;
    }

    for (int j = i; j < s.length(); j++)
    {
        swap(s[i], s[j]);
        compute(res, s, i + 1);
        swap(s[i], s[j]);
    }
}

int main()
{
    string T;
    cin >> T;
    vector<string> res;
    compute(res, T, 0);
    sort(res.begin(), res.end());
    for (auto x : res)
        cout << "[" << x << "]\n";
    return 0;
}