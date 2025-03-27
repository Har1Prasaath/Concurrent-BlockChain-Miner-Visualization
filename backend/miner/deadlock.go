package miner

import (
	"fmt"
	"sync"
)

// Resource represents a system resource that can be allocated
type Resource struct {
	ID int
}

// DeadlockDetector implements a simple deadlock detection algorithm
type DeadlockDetector struct {
	// Which process holds which resources
	allocations map[int][]int
	// Which process is waiting for which resources
	waitFor map[int][]int
	mutex   sync.Mutex
}

// NewDeadlockDetector creates a new deadlock detector
func NewDeadlockDetector() *DeadlockDetector {
	return &DeadlockDetector{
		allocations: make(map[int][]int),
		waitFor:     make(map[int][]int),
	}
}

// AddAllocation records that process holds resource
func (d *DeadlockDetector) AddAllocation(process, resource int) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.allocations[process] = append(d.allocations[process], resource)
}

// AddWaitFor records that process is waiting for resource
func (d *DeadlockDetector) AddWaitFor(process, resource int) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.waitFor[process] = append(d.waitFor[process], resource)
}

// DetectDeadlocks checks for deadlocks in the system
func (d *DeadlockDetector) DetectDeadlocks() [][]int {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Create wait-for graph: process -> processes it's waiting for
	waitForGraph := make(map[int][]int)

	// For each process waiting for resources
	for process, resources := range d.waitFor {
		for _, resource := range resources {
			// Find which processes hold the resource
			for holder, held := range d.allocations {
				for _, r := range held {
					if r == resource && holder != process {
						// Process is waiting for holder
						waitForGraph[process] = append(waitForGraph[process], holder)
						fmt.Printf("▸▸▸ Process %d waits for Process %d (which holds resource %d)\n",
							process, holder, resource)
					}
				}
			}
		}
	}

	// Now detect cycles in the wait-for graph using DFS
	deadlocks := [][]int{}
	visited := make(map[int]bool)

	for node := range waitForGraph {
		if !visited[node] {
			path := []int{}
			d.findDeadlockCycles(node, waitForGraph, visited, path, &deadlocks)
		}
	}

	return deadlocks
}

// findDeadlockCycles is a helper function to find cycles in the wait-for graph
func (d *DeadlockDetector) findDeadlockCycles(current int, graph map[int][]int, visited map[int]bool, path []int, cycles *[][]int) {
	// Mark the current node as visited
	visited[current] = true
	path = append(path, current)

	// Check all adjacent vertices
	for _, neighbor := range graph[current] {
		// Check if neighbor is already in path (cycle found)
		for i, node := range path {
			if node == neighbor {
				// Found a cycle
				cycle := append([]int{}, path[i:]...)
				cycle = append(cycle, neighbor)
				*cycles = append(*cycles, cycle)
				fmt.Printf("▸▸▸ Cycle detected: ")
				for _, p := range cycle[:len(cycle)-1] {
					fmt.Printf("%d → ", p)
				}
				fmt.Printf("%d\n", cycle[len(cycle)-1])
				return
			}
		}

		// If neighbor hasn't been visited yet
		if !visited[neighbor] {
			d.findDeadlockCycles(neighbor, graph, visited, path, cycles)
		}
	}
}

func (d *DeadlockDetector) PrintDeadlocks() {
	fmt.Println("▸▸▸ DEADLOCK DETECTION CHECK ▸▸▸")
	deadlocks := d.DetectDeadlocks()

	if len(deadlocks) == 0 {
		fmt.Println("▸▸▸ No deadlocks detected in the system ▸▸▸")
		return
	}

	fmt.Printf("▸▸▸ ALERT: Detected %d deadlocks in the system! ▸▸▸\n", len(deadlocks))
	for i, cycle := range deadlocks {
		fmt.Printf("▸▸▸ Deadlock #%d: Process", i+1)
		for _, process := range cycle[:len(cycle)-1] {
			fmt.Printf(" %d →", process)
		}
		fmt.Printf(" %d\n", cycle[len(cycle)-1])
	}
	fmt.Println("▸▸▸ END DEADLOCK DETECTION ▸▸▸")
}
