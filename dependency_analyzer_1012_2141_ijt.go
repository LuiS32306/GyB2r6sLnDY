// 代码生成时间: 2025-10-12 21:41:57
package main

import (
    "context"
    "fmt"
    "log"
    "strings"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

// DependencyGraph defines the structure of the dependency graph.
type DependencyGraph struct {
    nodes map[string][]string
}

// NewDependencyGraph creates a new DependencyGraph instance.
func NewDependencyGraph() *DependencyGraph {
    return &DependencyGraph{
        nodes: make(map[string][]string),
    }
}

// AddNode adds a node to the graph.
func (g *DependencyGraph) AddNode(node string) {
    if _, exists := g.nodes[node]; !exists {
        g.nodes[node] = []string{}
    }
}

// AddEdge adds an edge between two nodes in the graph.
func (g *DependencyGraph) AddEdge(from, to string) {
    if _, exists := g.nodes[from]; !exists {
        g.nodes[from] = []string{}
    }
    g.nodes[from] = append(g.nodes[from], to)
}

// ResolveDependencies recursively resolves dependencies for a given node.
func (g *DependencyGraph) ResolveDependencies(node string) []string {
    if _, exists := g.nodes[node]; !exists {
        return []string{}
    }
    
    resolved := make(map[string]bool)
    g.resolve(resolved, node)
    return keys(resolved)
}

func (g *DependencyGraph) resolve(resolved map[string]bool, node string) {
    if _, exists := resolved[node]; exists {
        return
    }
    resolved[node] = true
    for _, dep := range g.nodes[node] {
        g.resolve(resolved, dep)
    }
}

func keys(m map[string]bool) []string {
    var keys []string
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

// DependencyAnalyzerService defines the service interface for dependency analysis.
type DependencyAnalyzerService struct {
    Graph *DependencyGraph
}

// ResolveDependency is a method to resolve dependencies for a given node.
func (s *DependencyAnalyzerService) ResolveDependency(ctx context.Context, in *emptypb.Empty) (*ResolveDependencyResponse, error) {
    if ctx.Err() != nil {
        return nil, status.Errorf(codes.Canceled, "context canceled")
    }
    node := ctx.Value("node").(string) // Assume node is passed in the context
    dependencies := s.Graph.ResolveDependencies(node)
    return &ResolveDependencyResponse{Dependencies: dependencies}, nil
}

// ResolveDependencyResponse defines the response for dependency resolution.
type ResolveDependencyResponse struct {
    Dependencies []string `protobuf:"bytes,1,rep,name=dependencies"`
}

func main() {
    // Create a new dependency graph.
    graph := NewDependencyGraph()
    
    // Add nodes and edges to the graph.
    graph.AddNode("A")
    graph.AddNode("B")
    graph.AddNode("C")
    graph.AddEdge("A", "B")
    graph.AddEdge("A", "C")
    graph.AddEdge("B", "C")
    
    // Create a new service with the dependency graph.
    service := &DependencyAnalyzerService{Graph: graph}
    
    // Define a gRPC server.
    server := grpc.NewServer()
    
    // Register the service with the server.
    RegisterDependencyAnalyzerServiceServer(server, service)
    
    // Listen on port 50051.
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    // Serve the server.
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
