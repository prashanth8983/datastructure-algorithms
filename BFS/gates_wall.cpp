class Solution {
public:
    void wallsAndGates(vector<vector<int>>& rooms) {
        ios_base::sync_with_stdio(false);
        cin.tie(NULL);

        int n = rooms.size();
        int m = rooms[0].size();

        queue<pair<int, int>> q;

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (rooms[i][j] == 0) {
                    q.push({i, j});
                }
            }
        }

        while (!q.empty()) {
            int size = q.size();
            while (size--) {
                int dist = 1;
                pair<int, int> A = q.front();
                int i = A.first;
                int j = A.second;
                q.pop();

                if (i + 1 >= 0 && i + 1 < n && rooms[i + 1][j] != -1) {
                    if (rooms[i + 1][j] >= rooms[i][j] + 1) {
                        rooms[i + 1][j] = rooms[i][j] + 1;
                        q.push({i + 1, j});
                    }
                }

                if (i - 1 >= 0 && i - 1 < n && rooms[i - 1][j] != -1) {
                    if (rooms[i - 1][j] >= rooms[i][j] + 1) {
                        rooms[i - 1][j] = rooms[i][j] + 1;
                        q.push({i - 1, j});
                    }
                }

                if (j + 1 >= 0 && j + 1 < m && rooms[i][j + 1] != -1) {
                    if (rooms[i][j + 1] >= rooms[i][j] + 1) {
                        rooms[i][j + 1] = rooms[i][j] + 1;
                        q.push({i, j + 1});
                    }
                }

                if (j - 1 >= 0 && j - 1 < m && rooms[i][j - 1] != -1) {
                    if (rooms[i][j - 1] >= rooms[i][j] + 1) {
                        rooms[i][j - 1] = rooms[i][j] + 1;
                        q.push({i, j - 1});
                    }
                }
            }
        }
    }
};
