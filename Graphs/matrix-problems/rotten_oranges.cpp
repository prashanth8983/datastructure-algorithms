class Solution {
public:
    int orangesRotting(vector<vector<int>>& grid) {

        ios_base::sync_with_stdio(false);
        cin.tie(NULL);
        
        int n = grid.size();
        int m = grid[0].size();
        int minutes = 0;

        

        queue<pair<int,int>> q;
        int total = 0;

      
        for(int i=0;i<n; i++){
            for(int j=0; j<m; j++){
                if(grid[i][j] == 2){
                    q.push({i,j});
                }
                if(grid[i][j] == 1){
                    total++;
                }
            }
        }

        if(total == 0){
            return 0;
        }

        while(!q.empty() && total != 0){
            int size = q.size();
            while(size--){
                pair<int,int> A = q.front();
                int i = A.first;
                int j = A.second;
                q.pop();

                if(i+1>=0 && i+1<n && grid[i+1][j] == 1){
                            grid[i+1][j] = 2;
                            total--;
                            q.push({i+1,j});
                    }
                                
                                
                if(i-1>=0 && i-1<n && grid[i-1][j] == 1){
                            grid[i-1][j] = 2;
                            total--;
                            q.push({i-1,j});
                        }
                                
                                
                if(j+1>=0 && j+1<m && grid[i][j+1] == 1)
                            {
                                grid[i][j+1] = 2;
                                total--;
                                q.push({i,j+1});
                            }
                                
                                
                if(j-1>=0 && j-1<m && grid[i][j-1] == 1)
                {
                                grid[i][j-1] = 2;
                                total--;
                                q.push({i,j-1});
                }
            }
            minutes++;
        }

        if (total != 0)
            return -1;
        return minutes;
    }
};
