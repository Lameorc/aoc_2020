package day10

import (
	"log"
	"strconv"
)

type adapters map[int]bool
type deviceRating int

func newAdapters(input []string) (adapters, deviceRating) {
	a := make(adapters)
	d := 0
	for _, l := range input {
		v, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		a[v] = true

		// dont want to bother with float casting
		if v > d {
			d = v
		}
	}
	d += 3
	return a, deviceRating(d)
}

func part1(a adapters) int {
	rating := 0
	diff1 := 0
	diff3 := 0
	for {
		if a[rating+1] {
			diff1++
			rating++
		} else if a[rating+3] {
			diff3++
			rating += 3
		} else {
			// out of options -> add device rating
			diff3++
			return diff1 * diff3
		}
	}
}

type graph map[int][]int

func newGraph(a adapters, r deviceRating) graph {
	g := make(graph)
	g[0] = make([]int, 0)
	for node := range a {
		e := make([]int, 0)
		for _, s := range []int{1, 2, 3} {
			next := node + s
			if a[next] {
				e = append(e, next)
			}
			if next == int(r) {
				e = append(e, int(r))
			}
			if node == s {
				g[0] = append(g[0], node)
			}
		}
		g[node] = e
	}
	return g
}

func (g graph) countPaths(idx int, r deviceRating, visited map[int]int) int {
	if idx == int(r) {
		return 1
	}
	valid := 0
	for _, n := range g[idx] {
		if visited[n] == 0 {
			visited[n] = g.countPaths(n, r, visited)
		}
		valid = valid + visited[n]
	}
	return valid
}

func part2(a adapters, r deviceRating) int {
	g := newGraph(a, r)
	// log.Print("Got graph:")
	// for k, v := range g {
	// 	log.Printf("\t%d: %d", k, v)
	// }

	v := make(map[int]int)
	return g.countPaths(0, r, v)
}

// Solve the day's problem
func Solve(input []string) {
	a, r := newAdapters(input)
	p1 := part1(a)
	log.Printf("Part 1: %d", p1)
	p2 := part2(a, r)
	log.Printf("Part 2: %d", p2)
}
