// { Driver Code Starts
#include <bits/stdc++.h>
using namespace std;

// } Driver Code Ends
class Solution
{
public:
    bool DFSCycle(int s, int parent, vector<int> adj[], vector<bool> &visited)
    {
        visited[s] = true;

        for(auto x: visited)
            cout<<x<<" ";
        cout<<"("<<s<<", "<<parent<<")\n";

        for (auto u : adj[s])
        {
            if (!visited[u])
            {
                if (DFSCycle(u, s, adj, visited))
                    return true;
            }
            else if (u != parent)
                return true;
        }

        return false;
    }

    bool isCycle(int V, vector<int> adj[])
    {
        vector<bool> visited(V, false);
        for (int i = 0; i < V; i++)
        {
            if (!visited[i])
                if (DFSCycle(i, -1, adj, visited))
                    return true;
        }

        return false;
    }
};


int main()
{
    int tc;
    cin >> tc;
    while (tc--)
    {
        int V, E;
        cin >> V >> E;
        vector<int> adj[V];
        for (int i = 0; i < E; i++)
        {
            int u, v;
            cin >> u >> v;
            adj[u].push_back(v);
            adj[v].push_back(u);
        }
        Solution obj;
        bool ans = obj.isCycle(V, adj);
        if (ans)
            cout << "1\n";
        else
            cout << "0\n";
    }
    return 0;
}