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
                  // corresponding distance from the node
};

class Graph {
private:
  std::unordered_map<char, Node *> nodes; // list of nodes
  // used to give unique identification to the Node
  std::unordered_map<Node *, std::pair<int, Node *>>
      path;                             // final dist | pred array
  std::unordered_set<Node *> unvisited; // node | distance | predecessor
  std::vector<Node *> visited;
  Node *source = nullptr;

  void _clearPaths() {
    path.clear();
    visited.clear();
    unvisited.clear();
    for (std::pair<const char, Node *> &node : nodes) {
      unvisited.insert(node.second);
    }
  }

  int _getDistance(const char &dest) {
    if (path.empty()) { // Check if shortest paths have been calculated
      return -1;
    }
    return path[nodes[dest]].first;
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

  void addEdge(const char &from, const char &to, int distance) {
    if (nodes.count(from) && nodes.count(to)) {
      nodes[from]->neighbours[nodes[to]] = distance;
    } else {
      std::cerr << CError << "Error: Node(s) not found in the graph."
                << ColorOff << std::endl;
    }
  }

  void dijkstra() {
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
      path[node.second] = {INT_MAX, nullptr}; // Initialize distance to infinity
                                              // and predecessor to nullptr
    }
    path[source] = {0, source}; // Distance from source to itself is 0

    while (!unvisited.empty()) {
      // minimum distance in the corresponding neighbours
      int minDistance = INT_MAX;
      Node *minNode = nullptr;
      for (Node *const &node : unvisited) {
        if (path.count(node) > 0 && path[node].first < minDistance) {
          minDistance = path[node].first;
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

      for (const auto &[node, distance] : minNode->neighbours) {
        // if the node exist in the path
        int newDistance = path[minNode].first + distance;
        if (path[node].first > newDistance) {
          // if the new distance is lesser than the old distance the update it
          // in the path array
          path[node].second = minNode;    // update the predecessor node
          path[node].first = newDistance; // update the distance
        }
      }
    }
  }

  void printPath() {
    if (path.empty()) { // Check if shortest paths have been calculated
      std::cerr
          << CError
          << "Error: Shortest paths have not been calculated. Run Dijkstra's "
             "algorithm first."
          << ColorOff << std::endl;
      return;
    }
    std::cout << "Source Node :\t" << source->data << std::endl;
    std::cout << "|Node\t|Distance\t|Predecessor\t|" << std::endl;
    std::cout << "|----\t|--------\t|-----------\t|\n";
    for (const std::pair<Node *const, std::pair<int, Node *>> &node : path) {
      int dist = node.second.first == INT_MAX ? -1 : node.second.first;
      std::cout << "|" << node.first->data << "\t|" << dist << "\t\t|"
                << node.second.second->data << "\t\t|" << std::endl;
    }
  }
  // prints the route from the source to the destination
  void printRoute(const char &dest) {
    if (path.empty()) { // Check if shortest paths have been calculated
      std::cerr
          << CError
          << "Error: Shortest paths have not been calculated. Run Dijkstra's "
             "algorithm first."
          << ColorOff << std::endl;
      return;
    }
    if (dest != source->data) {
      printRoute(path[nodes[dest]].second->data);
    } else {
      // when you reach the source
      std::cout << "The route is:\t" << dest;
      return;
    }
    std::cout << " -> " << dest;
  }

  void printDistance(const char &dest) {
    std::cout << "\nThe distance of " << dest << " from " << source->data
              << " is " << _getDistance(dest) << std::endl;
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
  int dist;
  char vertex;
  while (op) {
    std::cout << "\n--------------------------------------------------------\n";
    std::cout
        << "Enter the option from the menu:\n\n1.Add a vertex\t2.Add an "
           "edge\t3.Set source node\t4.Run dijstra algorithm of the given "
           "data\n5.Print graph\t6.Print the "
           "shortest distance\t7.Find route to the destination\t8.Clear "
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
      std::cout << "Enter the source node :\t";
      std::cin >> source;
      graph.setSource(source);
      std::cout << "The source is " << graph.getSource()->data << std::endl;
      break;
    case 4:
      graph.dijkstra();
      break;
    case 5:
      graph.printGraph();
      break;
    case 6:
      graph.printPath();
      break;
    case 7:
      std::cout << "Enter the destination :\t";
      std::cin >> dest;
      graph.printRoute(dest);
      graph.printDistance(dest);
      break;
    case 8:
      graph.clearGraph();
      break;
    }
  }
  return 0;
}
