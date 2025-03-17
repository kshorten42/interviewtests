package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	a, b, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}

	if len(a) != len(b) {
		panic("different length arrays")
	}

	slices.Sort(a)
	slices.Sort(b)

	dist := 0

	for i := 0; i < len(a); i++ {
		d := absDiff(a[i], b[i])
		dist = dist + d
	}

	fmt.Println("Total distance:  ", dist)
}

func readFile(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var a []int
	var b []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "   ")
		aval, err := strconv.Atoi(s[0])
		if err != nil {
			return nil, nil, err
		}
		bval, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, nil, err
		}
		a = append(a, aval)
		b = append(b, bval)
	}
	return a, b, scanner.Err()

}

func absDiff(a, b int) int {

	d := a - b

	if d < 0 {
		return -d
	}
	return d
}
