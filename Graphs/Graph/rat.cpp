#include <bits/stdc++.h>

class Solution
{
public:
    vector<string> findPath(vector<vector<int>> &m, int n)
    {
        bool visited[n][n];
        memset(visited, false, sizeof(visited));

        vector<string> A;
        queue<pair<vector<int>, char>> q;
        pair<vector<int>, char> temp;
        q.push({{0, 0}, '\0'});

        int row, column;
        string temp_path = "";
        while (!q.empty())
        {
            int size = q.size();

            while (size--)
            {
                temp = q.front();
                q.pop();
                temp_path += temp.second;
                row = temp.first[0];
                column = temp.first[1];

                if (row == n - 1 && column == n - 1)
                {
                    A.push_back(temp_path);
                    string temp_path = "";
                    break;
                }

                if (!visited[row][column])
                {
                    visited[row][column] = true;
                    if (row < n - 1)
                        if (m[row + 1][column] == 1 && !visited[row + 1][column])
                            q.push({{row + 1, column}, 'R'});

                    if (row > 0)
                        if (m[row - 1][column] == 1 && !visited[row - 1][column])
                            q.push({{row - 1, column}, 'L'});

                    if (column < n - 1)
                        if (m[row][column + 1] == 1 && !visited[row][column + 1])
                            q.push({{row, column + 1}, 'D'});

                    if (column > 0)
                        if (m[row][column - 1] == 1 && !visited[row][column - 1])
                            q.push({{row, column - 1}, 'U'});
                }
            }
        }

        return A;
    }
};

// { Driver Code Starts.

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        int n;
        cin >> n;
        vector<vector<int>> m(n, vector<int>(n, 0));
        for (int i = 0; i < n; i++)
            for (int j = 0; j < n; j++)
                cin >> m[i][j];

        Solution obj;
        vector<string> result = obj.findPath(m, n);
        if (result.size() == 0)
            cout << -1;
        else
            for (int i = 0; i < result.size(); i++)
                cout << result[i] << " ";
        cout << endl;
    }
    return 0;
} 