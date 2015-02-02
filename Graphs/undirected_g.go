type Graph interface {
    // NodeExists returns true when node is currently in the graph.
    NodeExists(Node) bool

    // NodeList returns a list of all nodes in no particular order, useful for
    // determining things like if a graph is fully connected. The caller is
    // free to modify this list. Implementations should construct a new list
    // and not return internal representation.
    NodeList() []Node

    // Neighbors returns all nodes connected by any edge to this node.
    Neighbors(Node) []Node

    // EdgeBetween returns an edge between node and neighbor such that
    // Head is one argument and Tail is the other. If no
    // such edge exists, this function returns nil.
    EdgeBetween(node, neighbor Node) Edge
}
