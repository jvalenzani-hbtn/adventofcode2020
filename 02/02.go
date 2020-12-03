package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	load()

}

func load() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	var min int
	var max int
	var key string
	var pass string

	var totalValid1 = 0
	var totalValid2 = 0

	for {
		_, err := fmt.Fscanln(f, &min, &max, &key, &pass)
		if err != nil {

			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		max = max * -1
		letter := rune(key[0])
		count := 0
		for _, l := range pass {
			if l == letter {
				count++
			}
		}
		if count >= min && count <= max {
			totalValid1++
		}
		if (rune(pass[min-1]) == letter || rune(pass[max-1]) == letter) && !(rune(pass[min-1]) == letter && rune(pass[max-1]) == letter) { // XOR
			totalValid2++
		}
	}
	fmt.Printf("*** Total Valid Part 1: %d\n", totalValid1)
	fmt.Printf("*** Total Valid Part 2: %d\n", totalValid2)
}
