package main

/*
 Рекомендуется использовать быстрый (буферизованный) ввод и вывод
var in *bufio.Reader
var out *bufio.Writer
in = bufio.NewReader(os.Stdin)
out = bufio.NewWriter(os.Stdout)
defer out.Flush()

var a, b int
fmt.Fscan(in, &a, &b)
fmt.Fprint(out, a + b)
*/

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
	var t, width, height int
	var spaceBuffer, dashBuffer string
	for i := 0; i < 200; i++ {
		spaceBuffer += " "
		dashBuffer += "_"
	}
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &width, &height)
		for i := 0; i < height*2+1; i++ {
			if i == 0 {
				fmt.Fprint(out, spaceBuffer[0:height]+dashBuffer[0:width]+"\n")
				continue
			}
			if i == height*2 {
				fmt.Fprint(out, spaceBuffer[0:height-1]+"\\"+dashBuffer[0:width]+"/\n")
				continue
			}
			if i <= height {
				fmt.Fprint(out, spaceBuffer[0:height-i]+"/"+spaceBuffer[0:width+2*i-2]+"\\\n")
				continue
			}
			fmt.Fprint(out, spaceBuffer[0:i-height-1]+"\\"+spaceBuffer[0:width+(2*height-i)*2]+"/\n")
		}
		out.Flush()
	}
}
