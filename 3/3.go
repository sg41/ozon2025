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
	var t, k, m, n int

	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {

		fmt.Fscan(in, &k, &n, &m)
		in.ReadRune() //\n

		result := make([][]rune, n)

		for l := 0; l < n; l++ {
			result[l] = make([]rune, m)
			for r := 0; r < m; r++ {
				result[l][r] = '.'
			}
		}

		for j := 0; j < k; j++ {
			for l := 0; l < n; l++ {
				for r := 0; r < m; r++ {
					relief, _, _ := in.ReadRune()
					if result[l][r] == '.' && relief != '.' {
						result[l][r] = relief
					}
				}
				in.ReadRune() //\n
			}
			if j < k-1 {
				in.ReadRune() //\n
			}
		}

		for l := 0; l < n; l++ {
			for r := 0; r < m; r++ {
				fmt.Fprintf(out, "%c", result[l][r])
			}
			out.WriteRune('\n')
		}
		out.WriteRune('\n')
		out.Flush()
	}
}
