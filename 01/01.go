package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	input := load()
	partOne(input)
	partTwo(input)

}

func partOne(input []int) {
	check := 2020
	for i, v1 := range input[:len(input)-1] {
		for _, v2 := range input[i+1:] {
			if v1+v2 == check {
				fmt.Printf("Found %d + %d = %d | %d * %d = %d\n", v1, v2, check, v1, v2, v1*v2)
				return
			}
		}
	}
}

func partTwo(input []int) {
	check := 2020
	for i, v1 := range input[:len(input)-2] {
		for _, v2 := range input[i+1 : len(input)-2] {
			for _, v3 := range input[i+2:] {
				if v1+v2+v3 == check {
					fmt.Printf("Found %d + %d + %d = %d | %d * %d * %d = %d\n", v1, v2, v3, check, v1, v2, v3, v1*v2*v3)
					return
				}
			}
		}
	}
}

func load() []int {
	var t []int
	f, err := os.Open("input.txt") // os.OpenFile has more options if you need them
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	var num int

	for {
		_, err := fmt.Fscanf(f, "%d\n", &num) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		t = append(t, num)
	}
	return t
}
