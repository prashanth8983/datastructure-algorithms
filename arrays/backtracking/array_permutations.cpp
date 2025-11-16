#include <bits/stdc++.h>

using namespace std;


vector<vector<int>> A;

void permutate(vector<int> S, int start, int end)
{
    if (start == end)
        A.push_back(S);

    else
    {
        for (int i = start; i <= end; i++)
        {
            swap(S[i], S[start]);
            permutate(S, start + 1, end);
            swap(S[start], S[i]);
        }
    }
}

int main()
{

    string s;
    cin >> s;

    permutate(s, 0, s.length() - 1);

    return 0;
}