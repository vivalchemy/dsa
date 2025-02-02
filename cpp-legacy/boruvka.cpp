// TODO: The code doesn't work at all. Need a complete overhaul
#include <algorithm>
#include <climits>
#include <iostream>
#include <unordered_map>
#include <vector>

#define VERTICES 6
#define EDGES 10

struct Edge {
    int source;
    int destination;
    int weight;
};

// add an edge to the graph
void addEdge(int graph[VERTICES + 1][VERTICES + 1], std::vector<Edge>& edges,
             int source, int destination, int distance) {
    if (source > VERTICES || destination > VERTICES || source < 1 ||
        destination < 1) {
        std::cerr << "Enter a valid source and destination" << std::endl;
        return;
    }

    graph[source][destination] = distance;
    graph[destination][source] = distance;
    edges.push_back({source, destination, distance});
    edges.push_back({destination, source, distance});
}

// get the minimum edge for the corresponding source
Edge getMinimum(int graph[VERTICES + 1][VERTICES + 1], int source) {
    int min = INT_MAX;
    int minIndex = -1;

    for (int i = 1; i <= VERTICES; i++) {
        if (graph[source][i] < min && i != source && graph[source][i] != 0) {
            min = graph[source][i];
            minIndex = i;
        }
    }

    return {source, minIndex, min};
}

// remove an edge from the graph
void removeEdge(int graph[VERTICES + 1][VERTICES + 1], std::vector<Edge>& edges,
                Edge edge) {
    if (edge.source > VERTICES || edge.destination > VERTICES ||
        edge.source < 1 || edge.destination < 1) {
        std::cerr << "Enter a valid source and destination" << std::endl;
        return;
    }

    graph[edge.source][edge.destination] = 0;
    graph[edge.destination][edge.source] = 0;
    edges.erase(
        std::remove_if(edges.begin(), edges.end(),
                       [&edge](const Edge& e) {
                           return (e.source == edge.source &&
                                   e.destination == edge.destination &&
                                   e.weight == edge.weight);
                       }),
        edges.end());
}

void printIncludedEdges(Edge includedEdges[]) {
    std::cout << "Edges included in the spanning tree:" << std::endl;
    for (int i = 0; i < VERTICES - 1; i++) {
        std::cout << "Source: " << includedEdges[i].source
                  << " Destination: " << includedEdges[i].destination
                  << " Weight: " << includedEdges[i].weight << std::endl;
    }
}

// print the graph
void printGraph(int graph[VERTICES + 1][VERTICES + 1]) {
    for (int i = 1; i <= VERTICES; i++) {
        for (int j = 1; j <= VERTICES; j++) {
            std::cout << graph[i][j] << " ";
        }
        std::cout << std::endl;
    }
}

int main() {
    std::unordered_map<int, int> setMap;
    std::vector<Edge> edges;
    int graph[VERTICES + 1][VERTICES + 1] = {0};
    Edge includedEdges[VERTICES - 1];
    int includedEdgesIterator = 0;

    addEdge(graph, edges, 1, 2, 3);
    addEdge(graph, edges, 1, 3, 4);
    addEdge(graph, edges, 2, 3, 2);
    addEdge(graph, edges, 2, 4, 1);
    addEdge(graph, edges, 3, 4, 5);
    addEdge(graph, edges, 3, 5, 4);
    addEdge(graph, edges, 4, 5, 6);
    addEdge(graph, edges, 4, 6, 7);
    addEdge(graph, edges, 5, 6, 8);

    for (includedEdgesIterator = 1; includedEdgesIterator <= VERTICES;
         includedEdgesIterator++) {
        Edge newEdge = getMinimum(graph, includedEdgesIterator);
        if (setMap.count(newEdge.destination) && setMap.count(newEdge.source) &&
            setMap[newEdge.source] == setMap[newEdge.destination]) {
            continue;
        }

        if (!setMap.count(newEdge.source)) {
            setMap[newEdge.source] = newEdge.source;
        }

        setMap[newEdge.destination] = setMap[newEdge.source];
        removeEdge(graph, edges, newEdge);
        includedEdges[includedEdgesIterator - 1] = newEdge;
    }

    printIncludedEdges(includedEdges);

    int totalCost = 0;
    for (int i = 0; i < VERTICES - 1; ++i) {
        totalCost += includedEdges[i].weight;
    }

    std::cout << "Cost of the spanning tree: " << totalCost << std::endl;

    return 0;
}
