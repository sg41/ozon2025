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
				if result[i][j] == ' ' {
					result[i][j] = '~'
				}
			}
			in.ReadRune()
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if result[i][j] == '_' && (result[i][j-1] == '~' || result[i][j-1] == ' ' || result[i][j-1] == '\\') {
					MarkHexagon(result, i, j, n, m)
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

func MarkHexagon(field [][]rune, row, col, n, m int) bool {
	width := 0
	height := 0
	for i := col; i < m; i++ {
		if field[row][i] == '_' {
			width++
		} else {
			break
		}
	}
	for i := 1; i < n-row && col-i >= 0; i++ {
		if field[row+i][col-i] == '/' {
			height++
		} else {
			break
		}
	}

	if height == 0 || row+2*height >= n {
		return false
	}

	if !CheckHexagon(field, row, col-height, width, height) {
		return false
	}
	for h := 0; h < height; h++ {
		for w := 0; w < width+2*h; w++ {
			field[row+h+1][col-h+w] = ' '
			field[row+2*height-h-1][col-h+w] = ' '
		}
		if h < height-1 {
			field[row+2*height-h-1][col-h-1] = ' '
			field[row+2*height-h-1][col+width+h] = ' '
		}
	}

	return true
}

func CheckHexagon(field [][]rune, row, col, width, height int) bool {
	for j := col + height; j < col+height+width; j++ {
		if field[row][j] != '_' || field[row+2*height][j] != '_' {
			return false
		}
	}

	for h := 0; h < height; h++ {
		if field[row+h+1][col+height-h-1] != '/' ||
			field[row+height+h+1][col+h] != '\\' ||
			field[row+h+1][col+height+width+h] != '\\' ||
			field[row+height+h+1][col+width+2*height-h-1] != '/' {
			return false
		}
	}
	return true
}
