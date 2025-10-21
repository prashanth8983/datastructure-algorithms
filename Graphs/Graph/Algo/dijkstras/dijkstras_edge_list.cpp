# include <bits/stdc++.h>

void addEdge(vector<pair<int, int>> adj[], int u, int v, int weight){
    adj[u].push_back(make_pair(v, weight));
    adj[v].push_back(make_pair(u, weight));
}

void shortest_distance(vector<pair<int,int>> adj[], int v, int source){

    priority_queue<pair<int, int>, vector<pair<int,int>>, greater<pair<int,int>>> pq;

    vector<int> dist(v, INT_MAX);
    pq.push(make_pair(0, source));
    dist[source] = 0;

    while(!pq.empty()){
        int u = pq.top().second;
        pq.pop();

        for(auto x: adj[u]){
            int v = x.first;
            int weight = x.second;
            if(dist[v] > dist[u] + weight){
                dist[v] = dist[u] + weight;
                pq.push(make_pair(dist[v], v));
            }
        }
    }

    cout<<"Vertex \tDistance from source\n";
    for(int i=0; i<v; ++i){
        printf("%d\t\t%d\n", i, dist[i]);
    }
}

int main(){
    int V = 9;
    vector<pair<int,int>> adj[V];

    addEdge(adj, 0, 1, 4);
    addEdge(adj, 0, 7, 8);
    addEdge(adj, 1, 2, 8);
    addEdge(adj, 1, 7, 11);
    addEdge(adj, 2, 3, 7);
    addEdge(adj, 2, 8, 2);
    addEdge(adj, 2, 5, 4);
    addEdge(adj, 3, 4, 9);
    addEdge(adj, 3, 5, 14);
    addEdge(adj, 4, 5, 10);
    addEdge(adj, 5, 6, 2);
    addEdge(adj, 6, 7, 1);
    addEdge(adj, 6, 8, 6);
    addEdge(adj, 7, 8, 7);

    shortest_distance(adj, V, 0);

    return 0;
    
}