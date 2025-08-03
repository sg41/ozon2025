package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t, m, n int

	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {

		fmt.Fscan(in, &n, &m)
		in.ReadRune()
		result := make([][]rune, n)
		for i := range result {
			result[i] = make([]rune, m)
			for j := range result[i] {
				result[i][j], _, _ = in.ReadRune()
			}
			in.ReadRune()
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if result[i][j] == ' ' && !checkHexagon(result, i, j, n, m) {
					result[i][j] = '~'
				}
			}
		}

		for _, row := range result {
			fmt.Fprintln(out, string(row))
		}
		fmt.Fprintln(out)
	}
}

type Point struct {
	Row, Col int
}

func checkHexagon(field [][]rune, i, j, n, m int) bool {

	directions := []Point{{-1, 0}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, step := range directions {
		ni, nj := i+step.Row, j+step.Col
		if ni >= 0 && ni < n && nj >= 0 && nj < m && (field[ni][nj] == ' ' || field[ni][nj] == '~') {
			return false
		} else if ni < 0 || ni >= n || nj < 0 || nj >= m {
			return false
		}
	}
	return true
}
