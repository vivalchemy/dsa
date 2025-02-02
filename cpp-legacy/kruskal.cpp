#include <algorithm>
#include <iostream>
#include <unordered_map> // for dist and predecossor array
#include <unordered_set>
#include <vector>
using namespace std;

std::string ColorOff = "\033[0m";
std::string CProcess = "\033[0;34m";
std::string CError = "\033[0;31m";
std::string CWarning = "\033[0;33m";
std::string CSuccess = "\033[0;33m";

struct Node {
  char data; // stores the data of the Node
  std::unordered_map<Node *, int>
      neighbours; // stores the nodes and their
                  // corresponding distance from the node
};

class Graph {
private:
  std::unordered_map<char, Node *> nodes; // list of nodes
  // used to give unique identification to the Node
  std::vector<std::pair<std::pair<Node *, Node *>, int>> edges;
  std::unordered_map<char, std::unordered_set<Node *>>
      mapOfSets; // relate the set to the nodes in that set
  std::unordered_map<char, char>
      nodeSetRel; // relate a node to its corresponding set
  std::vector<std::pair<std::pair<Node *, Node *>, int>> includedEdges;
  int cost = 0;

  // compare two edges based on their distance
  static bool
  _compareDistances(const std::pair<std::pair<Node *, Node *>, int> &a,
                    const std::pair<std::pair<Node *, Node *>, int> &b) {
    return a.second < b.second; // Compare edges based on distance
  }

  // sort all the edges based on their distances
  void _sortEdges() {
    std::sort(edges.begin(), edges.end(), _compareDistances);
  }

public:
  // destructor to free all the memory used by the graph
  ~Graph() {
    for (auto &pair : nodes) {
      delete pair.second; // Deallocate memory for each node
    }
  }

  int getCost() {
    if (!cost) {
      std::cout << CError
                << "Please run the kruskal algorithm before accessing the cost"
                << ColorOff << std::endl;
    }
    return cost;
  }
  // add new node to the graph
  void addNode(const char &data) {
    if (nodes.count(data)) {
      std::cout << CError << "Node already exists in the graph" << ColorOff
                << std::endl;
      return;
    }
    nodes[data] = new Node{data, {}}; // create a node
    nodeSetRel[data] = data; // add it to the nodes set relation database
    mapOfSets[data].insert(nodes[data]); // add it to the list of sets
  }

  // connect two edges with direction
  void addEdge(const char &from, const char &to, int distance) {
    if (nodes.count(from) && nodes.count(to)) {
      nodes[from]->neighbours[nodes[to]] = distance;
      edges.push_back({{nodes[from], nodes[to]}, distance});
    } else {
      std::cerr << CError << "Error: Node(s) not found in the graph."
                << ColorOff << std::endl;
    }
  }

  // krsukal algorithm
  void krsukal() {
    _sortEdges();
    while (!edges.empty() && mapOfSets.size() > 1) {
      std::pair<std::pair<Node *, Node *>, int> edge = edges.front();
      edges.erase(edges.begin());
      std::cout << "Current Edge: " << edge.first.first->data << " -> "
                << edge.first.second->data << " : " << edge.second << std::endl;
      // if the edge causes a cycle then ignore it
      if (nodeSetRel[edge.first.first->data] ==
          nodeSetRel[edge.first.second->data]) {

        std::cerr << CWarning
                  << "Warning: Edge causes a cycle. Ignoring the edge"
                  << ColorOff << std::endl;
        continue;
      }
      // add the cost of the edge to the final cost
      cost += edge.second;

      char secondNodeSetName = nodeSetRel[edge.first.second->data];
      std::unordered_set secondSet = mapOfSets[secondNodeSetName];

      // make the set of the first node as the set of all the nodes in the
      // second set
      for (Node *node : secondSet) {
        nodeSetRel[node->data] = nodeSetRel[edge.first.first->data];
      }
      // combine the sets
      mapOfSets[nodeSetRel[edge.first.first->data]].insert(secondSet.begin(),
                                                           secondSet.end());

      mapOfSets.erase(secondNodeSetName);
      // add the edge to the list of included edges
      includedEdges.push_back(edge);
    }
  }

  void printEdges() {
    for (std::pair<std::pair<Node *, Node *>, int> edge : edges) {
      std::cout << edge.first.first->data << " - " << edge.first.second->data
                << " : " << edge.second << std::endl;
    }
  }

  // print the list of included edges
  void printIncludedEdges() {
    if (includedEdges.empty()) {
      std::cout << CError
                << "Please run the kruskal algorithm before accessing the edges"
                << ColorOff << std::endl;
      return;
    }
    for (std::pair<std::pair<Node *, Node *>, int> edge : includedEdges) {
      std::cout << edge.first.first->data << " - " << edge.first.second->data
                << " : " << edge.second << std::endl;
    }
  }

  // print the map of sets
  void printMapOfSets() {
    std::cout << "Map of sets printing..." << std::endl;
    for (auto set : mapOfSets) {
      std::cout << set.first << " : ";
      for (Node *node : set.second) {
        std::cout << node->data << " ";
      }
      std::cout << std::endl << std::endl;
    }
  }

  // print the node set relation
  void printNodeSetRel() {
    std::cout << "Node set Relation printing..." << std::endl << std::endl;
    for (auto nodeSet : nodeSetRel) {
      std::cout << nodeSet.first << " : " << nodeSet.second << std::endl;
    }
  }

  // print all the nodes and corresponding edges in the graph
  void printGraph() {
    for (std::pair<char, Node *> node : nodes) {
      std::cout << node.first << "->\t";
      for (std::pair<Node *const, int> neighbour : node.second->neighbours) {
        std::cout << neighbour.first->data << ":" << neighbour.second << "\t";
      }
      std::cout << std::endl;
    }
  }

  void clearGraph() {
    nodes.clear();
    edges.clear();
    mapOfSets.clear();
    nodeSetRel.clear();
    includedEdges.clear();
    cost = 0;
  }
};

// main function
int main() {
  Graph graph;
  int op = 1;
  char source, dest;
  int dist;
  char vertex;
  while (op) {
    std::cout << "\n--------------------------------------------------------\n";
    std::cout
        << "Enter the option from the menu:\n\n1.Add a vertex\t2.Add an "
           "edge\t3.Run the kruskal algorithm\t4.Get the cost of the spanning "
           "tree\n5.Print all the edges of the graph\t6.Print all the edges of "
           "the spanning tree \t7.Print the graph\t8.Clear the graph \n0.Exit"
        << std::endl;
    std::cout << "Option :\t";
    std::cin >> op;
    switch (op) {
    case 1: // Add the vertex
      std::cout << "Enter the vertex :\t";
      std::cin >> vertex;
      graph.addNode(vertex);
      break;
    case 2: // Add the vertex
      std::cout << "Enter the source :\t";
      std::cin >> source;
      std::cout << "Enter the destination :\t";
      std::cin >> dest;
      std::cout << "Enter the distance :\t";
      std::cin >> dist;
      graph.addEdge(source, dest, dist);
      break;
    case 3:
      std::cout << CProcess << "Running the kruskal algorithm..." << ColorOff
                << std::endl;
      graph.krsukal();
      std::cout << CSuccess << "The spanning tree is generated" << ColorOff
                << std::endl;
      break;
    case 4:
      std::cout << CSuccess << "The cost of spanning tree is "
                << graph.getCost() << ColorOff << std::endl;
      break;
    case 5:
      std::cout << CProcess << "Printing all the edges of the graph" << ColorOff
                << std::endl;
      graph.printEdges();
      break;
    case 6:
      std::cout << CProcess << "Printing all the edges of the spanning tree"
                << ColorOff << std::endl;
      graph.printIncludedEdges();
      break;
    case 7:
      std::cout << CProcess << "Printing the graph" << ColorOff << std::endl;
      graph.printGraph();
      break;
    case 8:
      std::cout << CWarning << "Clearing the graph" << ColorOff << std::endl;
      graph.clearGraph();
      break;
    }
  }
  return 0;
}
