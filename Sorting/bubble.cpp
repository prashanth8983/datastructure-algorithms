#include <bits/stdc++.h>

using namespace std;
void bubble_sort(vector<int> &A)
{
    for (int i = 0; i < A.size(); i++)
        for (int j = i + 1; j < A.size(); j++)
            if (A[j] < A[i])
                swap(A[i], A[j]);
}

int main()
{
    int n;
    cin >> n;
    vector<int> A(n);
    for (auto &x : A)
        cin >> x;
    bubble_sort(A);
    for (auto x : A)
        cout << x << " ";

    return 0;
}