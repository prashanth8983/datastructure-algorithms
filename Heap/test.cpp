#include <bits/stdc++.h>

using namespace std;

int main()
{

    int n;
    cin >> n;
    vector<int> B(n);
    for (auto &x : B)
        cin >> x;
    priority_queue<int, vector<int>, greater<int>> A;
    for (auto x : B)
        A.push(x);

    while (A.empty() == false)
    {
        cout << A.top() << " ";
        A.pop();
    }
    return 0;
}