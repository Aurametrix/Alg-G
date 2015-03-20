package cluster

import (
        "sort"
)

// Single linkage hierarchical clustering using Minimum Spanning Tree (MST) algorithm
type HClustersSingle struct {
        HClusters

        // minimum distances from each element to the growing MST
        minDistances []float64
}

func NewHClustersSingle(X Matrix, metric MetricOp, d *Distances) *HClustersSingle {
        if d == nil {
                d = NewDistances(X, metric)
        }
        return &HClustersSingle{
                HClusters: HClusters{
                        X: X,
                        Metric: metric,
                        Method: single_linkage,
                        D: d,
                },
        }
}

func (c *HClustersSingle) Cluster(k int) (classes *Classes) {
        if c.D == nil { return }

        c.Hierarchize()

        c.K = k
        c.CutTree(k)

        // copy classification information
        classes = &Classes{
                make([]int, len(c.X)), k, c.Cost }
        copy(classes.Index, c.Index)

        return
}

func (c *HClustersSingle) Hierarchize() Linkages {
        if c.D == nil { return nil }
        c.initialize()
        c.cluster()
        return c.Dendrogram
}

func (c *HClustersSingle) initialize() {
        m := len(c.X)

        c.Dendrogram = make([]Linkage, m-1)

        c.actives =  NewActiveSet(m)

        c.minDistances = make([]float64, m)
        for i := range c.minDistances {
                c.minDistances[i] = maxValue
        }
}

func (c *HClustersSingle) cluster() {
        m := len(c.X)
        
        // Simplifed MST method based on Prim's algorithm
        
        // start with the first element (any will be fine)
        current := c.actives.Begin()

        // (m - 1) merges will occur in the main loop
        for i := 0; i < m-1; i++ {
                // Remove current node from active set
                // ('Add' it to the growing MST; not actually done for the sake of efficiency)
                c.actives.Remove(current)
                
                // Calculate min distance of each element to MST,
                // while keeping track of minimum
                min, minIdx := maxValue, 0
                for s := c.actives.Begin(); s < m; s = c.actives.Next(s) {
                        d := c.D.rep[s][current]
                        if d < c.minDistances[s] {
                                c.minDistances[s] = d
                        }
                        // keep track of minimum
                        if c.minDistances[s] < min {
                                min, minIdx = c.minDistances[s], s
                        }
                }

                // Add the element with the minimum minDistance to the growing MST
                c.Dendrogram[i] = Linkage{ First:current, Second:minIdx, Distance:min }
                
                // Procceed onto the newly added element
                current = minIdx
        }

        // *Stably* sort the stepwise dendorgram, which is currently not in order
        // Ties in distance will cause problems (e.g. merging a node to a new cluster
        // which has yet been formed); therefore, a stable sorting algorithm is required
        // TODO Replace standard library sort with a stable sort
        sort.Sort(c.Dendrogram)
}


// Find element labels from cluster representatives for the dendrogram
func (c *HClustersSingle) label() {
        
}
Hide details
Change log
0bb5140bbda1 by David Jh Shih <djh.shih> on May 17, 2012   Diff
Added Hierarchizer interface.
Added test case for checking whether
kmeans returns same result after
permutation.
Added InvPerm(x): inverses a permutation.
Added constructor for Hopach struct.
Go to: 	
Project members, sign in to write a code review
Older revisions
 d2b8ef161853 by David Jh Shih <djh.shih> on May 17, 2012   Diff 
 18fdcd0b4921 by David Jh Shih <djh.shih> on May 12, 2012   Diff 
 f971a25b6296 by David Jh Shih <djh.shih> on May 9, 2012   Diff 
All revisions of this file
File info
Size: 2583 bytes, 111 lines
View raw file

