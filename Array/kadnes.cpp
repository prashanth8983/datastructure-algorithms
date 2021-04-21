#include <bits/stdc++.h>

using namespace std;

int maxSubarraySum(int arr[], int n)
{
    int sum = INT_MIN;
    int curr = 0;
    for (int i = 0; i < n; i++)
    {
        curr = max(arr[i], curr + arr[i]);
        sum = max(sum, curr);
    }
    return sum;
}

int main()
{
    int n;
    cin >> n;
    int A[n];
    for (auto &x : A)
        cin >> x;
    cout << maxSubarraySum(A, n);
    return 0;
}