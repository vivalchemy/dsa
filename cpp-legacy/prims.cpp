#include <climits>
#include <iostream>
#include <unordered_map> // for dist and predecossor array
#include <unordered_set>
#include <vector>
// using namespace std;

std::string ColorOff = "\033[0m";
std::string CProcess = "\033[0;34m";
std::string CError = "\033[0;31m";
std::string CWarning = "\033[0;33m";
std::string CSuccess = "\033[0;32m";

struct Node {
  char data; // stores the data of the Node
  std::unordered_map<Node *, int>
      neighbours; // stores the nodes and their
                  // corresponding weight from the node
};

class Graph {
private:
  std::unordered_map<char, Node *> nodes; // list of nodes
  // used to give unique identification to the Node
  std::unordered_map<Node *, std::pair<int, Node *>>
      path;                             // final weight | pred array
  std::unordered_set<Node *> unvisited; // node | weight | predecessor
  std::vector<Node *> visited;
  Node *source = nullptr;
  int cost = 0;

  void _clearPaths() {
    cost = 0;
    path.clear();
    visited.clear();
    unvisited.clear();
    for (std::pair<const char, Node *> &node : nodes) {
      unvisited.insert(node.second);
    }
  }

  void _getCost() {
    if (path.empty()) { // Check if shortest paths have been calculated
      std::cerr
          << CError
          << "Error: Shortest paths have not been calculated. Run prims's "
             "algorithm first."
          << ColorOff << std::endl;
      return;
    }
    cost = 0; // so that when it is called again the cost is not added
    for (const std::pair<Node *const, std::pair<int, Node *>> &node : path) {
      cost += node.second.first;
    }
  }

public:
  ~Graph() {
    for (std::pair<const char, Node *> &pair : nodes) {
      delete pair.second;
    }
  }
  Node *getSource() { return source; }

  void setSource(const char &newSource) {
    if (nodes.count(newSource)) {
      source = nodes[newSource];
      _clearPaths();
      return;
    }
    std::cerr << CError << "Error: There is no element " << newSource
              << ColorOff << std::endl;
  }

  void addNode(const char &data) {
    if (!nodes.count(data)) {
      nodes[data] = new Node{data, {}};
    }
  }

  void addEdge(const char &from, const char &to, int weight) {
    if (nodes.count(from) && nodes.count(to)) {
      nodes[from]->neighbours[nodes[to]] = weight;
    } else {
      std::cerr << CError << "Error: Node(s) not found in the graph."
                << ColorOff << std::endl;
    }
  }

  void prims() {
    if (source == nullptr) {
      std::cerr << CError << "Error: Source node not set." << ColorOff
                << std::endl;
      return;
    }
    if (nodes.size() <= 1) {
      std::cerr << CError << "Error: There are either no nodes or only one node"
                << ColorOff << std::endl;
      return;
    }

    for (const std::pair<const char, Node *> &node : nodes) {
      // check if there is an edge from the source to the other nodes
      // If not initialize it to INT_MAX
      if (!source->neighbours.count(node.second)) {
        source->neighbours[node.second] = INT_MAX;
      };
    }

    _clearPaths();
    source->neighbours[source] = 0;
    for (const std::pair<const char, Node *> &node : nodes) {
      path[node.second] = {INT_MAX, nullptr}; // Initialize weight to infinity
                                              // and predecessor to nullptr
    }
    path[source] = {0, source}; // Weight from source to itself is 0

    while (!unvisited.empty()) {
      // minimum weight in the corresponding neighbours
      int minWeight = INT_MAX;
      Node *minNode = nullptr;
      for (Node *const &node : unvisited) {
        if (path.count(node) > 0 && path[node].first < minWeight) {
          minWeight = path[node].first;
          minNode = node;
        }
      }
      // All the nodes have been visited
      if (!minNode) {
        break;
      }

      // change the data in the visited and unvisited graph
      unvisited.erase(minNode);
      visited.push_back(minNode);

      for (const auto &[node, weight] : minNode->neighbours) {
        if (path[node].first > weight) {
          // if the new weight is lesser than the old weight then update it
          // in the path array
          path[node].second = minNode; // update the predecessor node
          path[node].first = weight;   // update the weight
        }
      }
    }
  }

  // print the node | weight | pred map
  void printPath() {
    if (path.empty()) { // Check if shortest paths have been calculated
      std::cerr
          << CError
          << "Error: Shortest paths have not been calculated. Run prims's "
             "algorithm first."
          << ColorOff << std::endl;
      return;
    }
    std::cout << "Source Node :\t" << source->data << std::endl;
    std::cout << "|Node\t|Weight\t|Predecessor\t|" << std::endl;
    std::cout << "|----\t|------\t|-----------\t|\n";
    for (const std::pair<Node *const, std::pair<int, Node *>> &node : path) {
      int weight = node.second.first == INT_MAX ? -1 : node.second.first;
      std::cout << "|" << node.first->data << "\t|" << weight << "\t|"
                << node.second.second->data << "\t\t|" << std::endl;
    }

    printCost();
  }

  // print to total cost of the spanning tree
  void printCost() {
    _getCost();
    std::cout << "\nThe cost of spanning tree is " << cost << std::endl;
  }
  // Print the graph representation
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
    _clearPaths();
  }
};

// main function
int main() {
  Graph graph;
  int op = 1;
  char source, dest;
  int weight;
  char vertex;
  while (op) {
    std::cout << "\n--------------------------------------------------------\n";
    std::cout << "Enter the option from the menu:\n\n1.Add a vertex\t2.Add an "
                 "edge\t3.Set source node\t4.Run Prims algorithm on the given "
                 "data\n5.Print graph\t6.Print the "
                 "weight-predecessor array\t7.Print the cost of spannig "
                 "tree\t8.Clear "
                 "graph\n0.Exit"
              << std::endl;
    std::cout << "Option :\t";
    std::cin >> op;
    switch (op) {
    case 1: // Add the vertex
      std::cout << "Enter the vertex :\t";
      std::cin >> vertex;
      graph.addNode(vertex);
      break;
    case 2: // Add the edge
      std::cout << "Enter the source :\t";
      std::cin >> source;
      std::cout << "Enter the destination :\t";
      std::cin >> dest;
      std::cout << "Enter the weight :\t";
      std::cin >> weight;
      graph.addEdge(source, dest, weight);
      break;
    case 3: // Set the source node
      std::cout << "Enter the source node :\t";
      std::cin >> source;
      graph.setSource(source);
      std::cout << "The source is " << graph.getSource()->data << std::endl;
      break;
    case 4: // apply prims
      graph.prims();
      break;
    case 5: // print the graph
      graph.printGraph();
      break;
    case 6: // print the weight | pred map
      graph.printPath();
      break;
    case 7: // print the total cost of spannig tree
      graph.printCost();
      break;
    case 8: // clear the data in graph
      graph.clearGraph();
      break;
    }
  }
  return 0;
}
