package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// open file
	contents, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	contentsString := string(contents)
	contentsSplit := strings.Split(contentsString, "\n")

	var solution = make(map[int]int64)

	var elf = 1

	for _, line := range contentsSplit {
		if line == "" {
			elf++
			continue
		}
		parsed, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		solution[elf] += parsed
	}

	// sort
	keys := make([]int, 0, len(solution))
	for key := range solution {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return solution[keys[i]] > solution[keys[j]]
	})

	fmt.Printf("The elf with the highest value is: %d\n", solution[keys[0]])
	var total = int64(0)
	for i := 0; i < 3; i++ {
		total += solution[keys[i]]
	}
	fmt.Printf("The total of the top 3 elves is: %d\n", total)
}
