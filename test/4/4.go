package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	Row, Col int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		sourceMap := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &sourceMap[i])
		}

		regions := make(map[byte][]Point)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				c := sourceMap[i][j]
				if c != '.' {
					regions[c] = append(regions[c], Point{i, j})
				}
			}
		}

		valid := true
		for region, cells := range regions {
			if len(cells) == 0 {
				continue
			}

			visited := make(map[Point]bool)
			queue := []Point{cells[0]}
			visited[cells[0]] = true

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				for _, neighbor := range getNeighbors(current.Row, current.Col, n, m, sourceMap) {
					if !visited[neighbor] && sourceMap[neighbor.Row][neighbor.Col] == region {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}

			if len(visited) != len(cells) {
				valid = false
				break
			}
		}

		if valid {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func getNeighbors(i, j, n, m int, sourceMap []string) []Point {
	var neighbors []Point

	directions := []Point{{-1, -1}, {-1, 1}, {0, -2}, {0, 2}, {1, -1}, {1, 1}}

	for _, step := range directions {
		ni, nj := i+step.Row, j+step.Col
		if ni >= 0 && ni < n && nj >= 0 && nj < m && sourceMap[ni][nj] != '.' {
			neighbors = append(neighbors, Point{ni, nj})
		}
	}

	return neighbors
}
