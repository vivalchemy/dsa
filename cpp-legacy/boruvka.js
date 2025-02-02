class Graph {
  constructor(vertices) {
    this.vertices = vertices; // No. of vertices
    this.graph = []; // dictionary to store graph
  }

  // function to add an edge to graph
  addEdge(u, v, w) {
    this.graph.push([u, v, w]);
  }

  find(parent, i) {
    if (parent[i] === i) {
      return i;
    }
    return this.find(parent, parent[i]);
  }

  union(parent, rank, x, y) {
    const xroot = this.find(parent, x);
    const yroot = this.find(parent, y);

    if (rank[xroot] < rank[yroot]) {
      parent[xroot] = yroot;
    } else if (rank[xroot] > rank[yroot]) {
      parent[yroot] = xroot;
    } else {
      parent[yroot] = xroot;
      rank[xroot] += 1;
    }
  }

  //The main function to construct minimum spanning tree using boruvka's algorithm
  boruvka() {
    const parent = [];

    const rank = [];
    const cheapest = [];

    let numTrees = this.vertices;
    let MSTweight = 0;

    // Create vertices subsets with single elements
    for (let node = 0; node < this.vertices; node++) {
      parent.push(node);
      rank.push(0);
      cheapest[node] = -1;
    }

    while (numTrees > 1) {
      for (let i = 0; i < this.graph.length; i++) {
        const [u, v, w] = this.graph[i];
        const set1 = this.find(parent, u);
        const set2 = this.find(parent, v);

        if (set1 !== set2) {
          if (cheapest[set1] === -1 || cheapest[set1][2] > w) {
            cheapest[set1] = [u, v, w];
          }

          if (cheapest[set2] === -1 || cheapest[set2][2] > w) {
            cheapest[set2] = [u, v, w];
          }
        }
      }

      // Consider the above picked cheapest edges and add them to minimum spanning tree
      for (let node = 0; node < this.vertices; node++) {
        // Check if cheapest for current set exists
        if (cheapest[node] !== -1) {
          const [u, v, w] = cheapest[node];
          const set1 = this.find(parent, u);
          const set2 = this.find(parent, v);

          if (set1 !== set2) {
            MSTweight += w;
            this.union(parent, rank, set1, set2);
            console.log(`Edge: ${u}-${v} --> ${w}`);
            numTrees--;
          }
        }
      }

      for (let node = 0; node < this.vertices; node++) {
        // reset cheapest array
        cheapest[node] = -1;
      }
    }

    console.log(`\nTotal cost of minimum spanning tree is ${MSTweight}`);
  }
}

let g = new Graph(7);
g.addEdge(1, 2, 3);
g.addEdge(1, 3, 4);
g.addEdge(2, 3, 2);
g.addEdge(2, 4, 1);
g.addEdge(3, 4, 5);
g.addEdge(3, 5, 4);
g.addEdge(4, 0, 6);
g.addEdge(4, 6, 7);
g.addEdge(5, 6, 8);

console.log("The edges included in MST are: \n");
g.boruvka();

// This code is contributed by prajwal kandekar
