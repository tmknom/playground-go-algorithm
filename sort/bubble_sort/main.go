package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(A []int, n int) int {
	counter := 0

	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if A[j-1] > A[j] {
				A[j-1], A[j] = A[j], A[j-1]
				counter += 1
			}
		}
	}

	return counter
}

func main() {
	// 初期化
	number, _ := strconv.Atoi(nextLine())
	inputs := strings.Split(nextLine(), " ")
	elements := make([]int, number)
	for i := 0; i < len(inputs); i++ {
		intValue, _ := strconv.Atoi(inputs[i])
		elements[i] = intValue
	}

	counter := bubbleSort(elements, number)
	printAll(elements)
	fmt.Println(counter)
}

func printAll(elements []int) {
	result := ""
	for _, value := range elements {
		result += strconv.Itoa(value) + " "
	}
	fmt.Println(strings.Trim(result, " "))
}

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
