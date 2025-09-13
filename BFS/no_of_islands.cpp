class Solution {
public:
    int numIslands(vector<vector<char>>& grid) {
        
        ios_base::sync_with_stdio(false);
        cin.tie(NULL);


        int n = grid.size();
        int m = grid[0].size();

        vector<vector<bool>> visited(n, vector<bool>(m,false));
      
        queue<pair<int,int>> q;

        int island = 0;

        for(int x=0; x<n; x++){
            for(int y=0; y<m; y++){
                if(!visited[x][y] and grid[x][y] == '1'){
                    q.push({x,y});
                    visited[x][y] = true;
                    while(!q.empty()){
                        pair<int, int> curr = q.front();
                        int i = curr.first;
                        int j = curr.second;

                        q.pop();

                        if(i+1>=0 && i+1<n && grid[i+1][j] == '1' && visited[i+1][j] == false){
                            visited[i+1][j] = true;
                            q.push({i+1,j});
                        }
                                
                                
                        if(i-1>=0 && i-1<n && grid[i-1][j] == '1' && visited[i-1][j] == false){
                            visited[i-1][j] = true;
                            q.push({i-1,j});
                        }
                                
                                
                        if(j+1>=0 && j+1<m && grid[i][j+1] == '1' && visited[i][j+1] == false)
                            {
                                visited[i][j+1] = true;
                                q.push({i,j+1});
                            }
                                
                                
                            if(j-1>=0 && j-1<m && grid[i][j-1] == '1' && visited[i][j-1] == false)
                            {
                                visited[i][j-1] = true;
                                q.push({i,j-1});
                            }
                    }
                    island++;
                }
            }
        }

        return island;


    }
};
