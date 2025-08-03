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
	"sort"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t, n int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		fmt.Fscanln(in)

		scores := make(map[string]int)
		var action string
		for j := 0; j < n; j++ {
			points := 1
			var sentence string
			sentence, _ = in.ReadString('\n')
			words := strings.Split(sentence, " ")
			name := strings.Trim(words[0], ":")
			scores[name] += 0
			suspect := strings.Trim(words[1], " ")
			if suspect == "I" && words[2] == "am" {
				suspect = name
				points = 2
			}
			action = strings.Trim(words[len(words)-1], "!\n")
			did := true
			negative := strings.Trim(words[len(words)-2], " ")
			if negative == "not" {
				did = false
			}
			if did {
				scores[suspect] += points
			} else {
				scores[suspect]--
			}
		}
		sortedSuspects := sortMapKeysByValue(scores)
		maxScore := scores[sortedSuspects[0]]
		topSuspects := make([]string, 0)
		for _, suspect := range sortedSuspects {
			if scores[suspect] == maxScore {
				topSuspects = append(topSuspects, suspect)
			} else {
				break
			}
		}
		sort.Strings(topSuspects)
		for _, suspect := range topSuspects {
			fmt.Fprintf(out, "%s is %s.\n", suspect, action)
		}
	}
}

func sortMapKeysByValue(unsortedMap map[string]int) []string {
	keys := make([]string, 0, len(unsortedMap))
	for k := range unsortedMap {
		keys = append(keys, k)
	}

	// Сортируем ключи по значениям
	sort.SliceStable(keys, func(i, j int) bool {
		return unsortedMap[keys[i]] > unsortedMap[keys[j]]
	})

	return keys
}
