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
	var m, n, width, height, k int

	fmt.Fscan(in, &m, &n, &width, &height, &k)

	result := make([][]byte, n+2)
	for i := range result {
		result[i] = make([]byte, m+2)
		for j := range result[i] {
			result[i][j] = ' '
		}
	}

	result[0][0] = '+'
	result[0][m+1] = '+'
	result[n+1][0] = '+'
	result[n+1][m+1] = '+'
	for j := 1; j <= m; j++ {
		result[0][j] = '-'
		result[n+1][j] = '-'
	}
	for i := 1; i <= n; i++ {
		result[i][0] = '|'
		result[i][m+1] = '|'
	}

	hexWidth := width + 2*height
	hexHeight := 2 * height

	count := 0
	for i := 1; count < k && i < n; i += hexHeight / 2 {
		start := 1
		if (i-1)%hexHeight != 0 {
			start = height + width + 1
		}
		for j := start; count < k; j += hexWidth + width {
			if i+hexHeight > n+2 || j+hexWidth >= m+2 {
				break
			}
			drawHexagon(result, i, j, width, height)
			count++
		}
	}

	for _, row := range result {
		fmt.Fprintln(out, string(row))
	}
}

func drawHexagon(field [][]byte, row, col, width, height int) {
	for j := col + height; j < col+height+width; j++ {
		field[row][j] = '_'
		field[row+2*height][j] = '_'
	}

	for h := 0; h < height; h++ {
		field[row+h+1][col+height-h-1] = '/'
		field[row+height+h+1][col+h] = '\\'
		field[row+h+1][col+height+width+h] = '\\'
		field[row+height+h+1][col+width+2*height-h-1] = '/'
	}
}
