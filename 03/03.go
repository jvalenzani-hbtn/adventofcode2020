package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var lines []string

func main() {
	file := "input.txt"
	loadLines(file)
	fmt.Printf("Final result: %d\n", treeFind(1, 1)*treeFind(3, 1)*treeFind(5, 1)*treeFind(7, 1)*treeFind(1, 2))
}

func loadLines(file string) {
	var line string
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, err := fmt.Fscanf(f, "%s\n", &line)
		if err != nil {

			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		lines = append(lines, line)
	}
}

func treeFind(incX int, incY int) int {

	var trees = 0
	var x = 0 + incX
	var y = 0 + incY

	for y < len(lines) {
		checkPos := x % len(lines[y])
		if lines[y][checkPos] == '#' {
			trees++
		}
		x += incX
		y += incY
	}
	fmt.Printf("*** Slope (%d,%d) - Total tree hits: %d\n", incX, incY, trees)
	return trees
}
