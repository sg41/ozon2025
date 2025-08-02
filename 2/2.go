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
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t, m, n int
	var A, B Point
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		in.ReadRune()
		field := make([][]rune, n)
		for j := 0; j < n; j++ {
			field[j] = make([]rune, m)
			for k := 0; k < m; k++ {
				field[j][k], _, _ = in.ReadRune()
				if field[j][k] == 'A' {
					A.Row = j
					A.Col = k
				}
				if field[j][k] == 'B' {
					B.Row = j
					B.Col = k
				}
			}
			in.ReadRune() //\n
		}
		if A.Row*A.Row+A.Col*A.Col <= B.Row*B.Row+B.Col*B.Col {
			MoveRobot(field, A, Point{0, 0}, 'a')
			MoveRobot(field, B, Point{n - 1, m - 1}, 'b')
		} else {
			MoveRobot(field, B, Point{0, 0}, 'b')
			MoveRobot(field, A, Point{n - 1, m - 1}, 'a')
		}
		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				fmt.Fprintf(out, "%c", field[j][k])
			}
			fmt.Fprintln(out)
		}
		out.Flush()
	}
}

func MoveRobot(field [][]rune, start, end Point, robot rune) {
	var rowFirst bool
	if start.Row%2 == 0 {
		rowFirst = true
	}
	if start.Row > end.Row {
		start.Row, end.Row = end.Row, start.Row
	}
	if start.Col > end.Col {
		start.Col, end.Col = end.Col, start.Col
	}
	if field[start.Row][start.Col] == '.' {
		field[start.Row][start.Col] = robot
	}
	for start.Row != end.Row || start.Col != end.Col {
		if rowFirst {
			if start.Row < end.Row {
				if field[start.Row+1][start.Col] == '.' {
					start.Row++
					field[start.Row][start.Col] = robot
					continue
				} else if field[start.Row+1][start.Col] != '#' {
					break
				}
			}
			if start.Col < end.Col {
				if field[start.Row][start.Col+1] == '.' {
					start.Col++
					field[start.Row][start.Col] = robot
				} else if field[start.Row][start.Col+1] != '#' {
					break
				}
			}
		} else {
			if start.Col < end.Col {
				if field[start.Row][start.Col+1] == '.' {
					start.Col++
					field[start.Row][start.Col] = robot
					continue
				} else if field[start.Row][start.Col+1] != '#' {
					break
				}
			}
			if start.Row < end.Row {
				if field[start.Row+1][start.Col] == '.' {
					start.Row++
					field[start.Row][start.Col] = robot
				} else if field[start.Row+1][start.Col] != '#' {
					break
				}
			}

		}
	}
}
