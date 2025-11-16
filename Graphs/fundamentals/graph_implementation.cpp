// { Driver Code Starts
#include<bits/stdc++.h>
using namespace std;

 // } Driver Code Ends
class Solution {
public:
	vector<int>bfsOfGraph(int V, vector<vector<int>> adj)
	{
	    vector<bool> visited(V,false);
	    vector<int> A;
	    queue<int> q;
	    
	    q.push(adj[0][0]);
	    
	    visited[adj[0][0]] = true;
	    
	    while(!q.empty())
	    {
	        int s = q.front();
	        q.pop();
	        A.push_back(s);
	        
	        for(auto u: adj[s])
	        {
	            if(visited[u] == false)
	            {
	                visited[u] = true;
	                q.push(u);
	                
	            }
	        } 
	    }
        return A;
	}
};

// { Driver Code Starts.
int main(){
	int tc;
	cin >> tc;
	while(tc--){
		int V, E;
    	cin >> V >> E;

    	vector<vector<int>> adj(V);

    	for(int i = 0; i < E; i++)
    	{
    		int u, v;
    		cin >> u >> v;
    		adj[u].push_back(v);
    // 		adj[v].push_back(u);
    	}
        // string s1;
        // cin>>s1;
        Solution obj;
        vector<int>ans=obj.bfsOfGraph(V, adj);
        for(int i=0;i<ans.size();i++){
        	cout<<ans[i]<<" ";
        }
        cout<<endl;
	}
	return 0;
}  // } Driver Code Ends