package models

type DfsRequest struct{
   Edges [][]int `json:"edges,omitempty"`
   Start int `json:"start,omitempty"`
   End int `json:"end,omitempty"`
}

type DfsResponse struct{
   Path [][]int `json:"Path,omitempty"`
}