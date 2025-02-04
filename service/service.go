package service

import (
    "streakai/errors"
    "streakai/models"
)

func addedges(a int, b int, graph map[int][]int) {
    graph[a] = append(graph[a], b)
    graph[b] = append(graph[b], a)
}


func CalculateDfs(payload *models.DfsRequest) (*models.DfsResponse, *errors.ErrorDetails) {
    if payload == nil {
        return nil, errors.MakeErrorMessage("PAYLOAD_CANNOT_BE_EMPTY", "Payload Cannot be Empty")
    }
    if len(payload.Edges) == 0 {
        return nil, errors.MakeErrorMessage("EDGES_CANNOT_BE_EMPTY", "Edges Cannot be Empty")
    }
    
    Graph := make(map[int][]int) 
    
    for _, val := range payload.Edges {
        if len(val) != 2 {
            return nil, errors.MakeErrorMessage("INVALID_EDGE_FORMAT", "Each edge must contain exactly two nodes")
        }
        addedges(val[0], val[1], Graph)
    }
    
    if _, exists := Graph[payload.Start]; !exists {
        return nil, errors.MakeErrorMessage("START_NODE_NOT_FOUND", "Start node does not exist in the graph")
    }
    if _, exists := Graph[payload.End]; !exists {
        return nil, errors.MakeErrorMessage("END_NODE_NOT_FOUND", "End node does not exist in the graph")
    }
    
    paths := AllPaths(payload.Start, payload.End, Graph)
    
    return &models.DfsResponse{Path: paths}, nil
}

func AllPaths(start int, end int, graph map[int][]int) [][]int {
    if start == end {
        return [][]int{{start}} 
    }
    
    visited := make(map[int]bool)
    var paths [][]int
    var path []int

    dfs(start, end, visited, path, &paths, graph)
    return paths
}


func dfs(s int, e int, visited map[int]bool, path []int, paths *[][]int, graph map[int][]int) {
    path = append(path, s)
    
    if s == e {
        copyPath := make([]int, len(path))
        copy(copyPath, path)
        *paths = append(*paths, copyPath)
        return
    }
    
    visited[s] = true 
    
    for _, neighbor := range graph[s] {
        if !visited[neighbor] {
            dfs(neighbor, e, visited, path, paths, graph)
        }
    }
    
    visited[s] = false 
}