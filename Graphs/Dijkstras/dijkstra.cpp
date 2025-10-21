#include <iostream>

using namespace std;

vector<vector<vector<int>>> construct_adj(vector<vector<int>> &edges, int V)
{
    vector<vector<vector<int>>> adj(V);
    for (const auto &edge : edges)
    {
        int u = edge[0];
        int v = edge[1];
        int wt = edge[2];
        adj[u].push_back({v, wt});
        adj[v].push_back({u, wt});
    }
    return adj;
}

vector<int> dijkstra(int V, vector<vector<int>> &edges, int src, int dest)
{
    vector<vector<vector<int>>> adj = construct_adj(edges, V);

    priority_queue<vector<int>, vector<vector<int>>, greater<vector<int>>> pq;
    vector<int> dist(V, INT_MAX);
    vector<int> parent(V, -1);

    pq.push({0, src});

    dist[src] = 0;

    while (!pq.empty())
    {
        int u = pq.top()[1];
        int current_dist = pq.top()[0];

        pq.pop();
        if (current_dist > dist[u])
            continue;

        for (auto x : adj[u])
        {
            int v = x[0];
            int weight = x[1];

            if (dist[v] > dist[u] + weight)
            {
                dist[v] = dist[u] + weight;
                parent[v] = u;
                pq.push({dist[v], v});
            }
        }
    }

    vector<int> path;
    if (dist[dest] == INT_MAX)
        return path; // No path exists

    for (int v = dest; v != -1; v = parent[v])
    {
        path.push_back(v);
    }
    reverse(path.begin(), path.end());

    return path;
}