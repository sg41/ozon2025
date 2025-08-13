package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	Row, Col int
}

type Hexagon struct {
	TopLeft       Point
	Height, Width int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t, m, n int
	var point1, point2 Point

	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {

		fmt.Fscan(in, &n, &m)
		in.ReadRune()
		fullMap := make([][]rune, n)
		for i := range fullMap {
			fullMap[i] = make([]rune, m)
			for j := range fullMap[i] {
				fullMap[i][j], _, _ = in.ReadRune()
				if fullMap[i][j] == ' ' {
					fullMap[i][j] = '~'
				}
			}
			in.ReadRune()
		}

		fmt.Fscan(in, &point1.Row, &point1.Col, &point2.Row, &point2.Col)
		point1.Row--
		point1.Col--
		point2.Row--
		point2.Col--

		var startHex, finishHex *Hexagon
		hexList := make([]*Hexagon, 0)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if fullMap[i][j] == '_' && (fullMap[i][j-1] == '~' || fullMap[i][j-1] == ' ' || fullMap[i][j-1] == '\\') {
					hex := CreateHexagon(fullMap, i, j, n, m)
					if hex != nil {
						if point1.IsInside(hex) {
							startHex = hex
						}
						if point2.IsInside(hex) {
							finishHex = hex
						}
						hexList = append(hexList, hex)
					}
				}
			}
		}
		if startHex == finishHex {
			fmt.Fprintln(out, "YES")
			continue
		}

		colOffset := startHex.Height
		colStep := startHex.Width + startHex.Height
		rowOffset := 0
		rowStep := startHex.Height

		miniMap := make([][]rune, n/rowStep+1)
		for i := range miniMap {
			miniMap[i] = make([]rune, m/colStep+2)
		}

		for _, hex := range hexList {
			miniMap[(hex.TopLeft.Row-rowOffset)/rowStep][(hex.TopLeft.Col-colOffset)/colStep] = 8
		}

		startPoint := Point{Row: (startHex.TopLeft.Row - rowOffset) / rowStep, Col: (startHex.TopLeft.Col - colOffset) / colStep}
		finishPoint := Point{Row: (finishHex.TopLeft.Row - rowOffset) / rowStep, Col: (finishHex.TopLeft.Col - colOffset) / colStep}

		if findPath(miniMap, startPoint, finishPoint, n/rowStep+1, m/colStep+2) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}

	}
}

func findPath(miniMap [][]rune, startPoint, finishPoint Point, n, m int) bool {

	visited := make(map[Point]bool)
	queue := []Point{startPoint}
	visited[startPoint] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range getNeighbors(current.Row, current.Col, n, m, miniMap) {
			if neighbor == finishPoint {
				return true
			}
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return false
}

func getNeighbors(i, j, n, m int, sourceMap [][]rune) []Point {
	var neighbors []Point

	directions := []Point{{-1, -1}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 1}, {-2, 0}, {2, 0}}

	for _, step := range directions {
		ni, nj := i+step.Row, j+step.Col
		if ni >= 0 && ni < n && nj >= 0 && nj < m && sourceMap[ni][nj] == 8 {
			neighbors = append(neighbors, Point{ni, nj})
		}
	}

	return neighbors
}

func CreateHexagon(field [][]rune, leftTopCornerRow, leftTopCornerCol, n, m int) (h *Hexagon) {
	row := leftTopCornerRow
	col := leftTopCornerCol
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
		return nil
	}

	if !CheckHexagon(field, row, col-height, width, height) {
		return nil
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

	return &Hexagon{Point{Row: leftTopCornerRow, Col: leftTopCornerCol}, height, width}
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

func (p *Point) IsInside(h *Hexagon) bool {
	topLeftCornerRow := h.TopLeft.Row
	topLeftCornerCol := h.TopLeft.Col
	width := h.Width
	height := h.Height

	rowOk := p.Row > topLeftCornerRow && p.Row < topLeftCornerRow+2*height
	if !rowOk {
		return false
	}
	relativeRow := p.Row - topLeftCornerRow
	var rowStart, rowWidth int
	if relativeRow <= height {
		rowWidth = width + 2*(relativeRow-1)
		rowStart = topLeftCornerCol - (relativeRow - 1)
	} else {
		rowWidth = width + 2*(height-1) - (relativeRow - height - 1)
		rowStart = topLeftCornerCol - (height - 1) + (relativeRow - height - 1)
	}
	if p.Col >= rowStart && p.Col < rowStart+rowWidth {
		return true
	}
	return false

}
