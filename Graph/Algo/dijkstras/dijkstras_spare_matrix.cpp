#include <bits/stdc++.h>

int min_distance(vector<int> &dist, vector<bool> &sptSet){
    int min = INT_MAX, idx = -1;
    for(int i=0; i< dist.size(); i++){
        if(!sptSet[i] && dist[i] < min)
            min = dist[i], idx = i;
    }
    return idx;
}

void dijkstra(vector<vector<int>> &graph, int src){
    int V = graph.size();
    vector<int> dist(V, INT_MAX);

    vector<bool> sptset(V, false);
    dist[src] = 0;

    for(int count =0; count < V-1; count++){
        int u = min_distance(dist, sptset);
        if(u == -1)
            break;
        sptset[u] = true;

        for(int v=0; v<V; v++){
            if(!sptset[v] && graph[u][v] && dist[u] != INT_MAX && dist[u] + graph[u][v] < dist[v])
                dist[v] = dist[u] + graph[u][v];
        }

        cout<< "Verex \t distance from source \n";
        for(int i=0; i<V; i++){
            cout<<i<<" \t\t "<<dist[i]<<"\n";
        } 
    }
}

int main(){
    vector<vector<int>> graph = {
        {0, 4, 0, 0, 0, 0, 0, 8, 0},  {4, 0, 8, 0, 0, 0, 0, 11, 0}, {0, 8, 0, 7, 0, 4, 0, 0, 2},
        {0, 0, 7, 0, 9, 14, 0, 0, 0}, {0, 0, 0, 9, 0, 10, 0, 0, 0}, {0, 0, 4, 14, 10, 0, 2, 0, 0},
        {0, 0, 0, 0, 0, 2, 0, 1, 6},  {8, 11, 0, 0, 0, 0, 1, 0, 7}, {0, 0, 2, 0, 0, 0, 6, 7, 0}};

    dijkstra(graph, 0);
    return 0;
}