package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//5 6 4 2 1 3
func selectionSort(A []int) int {
	counter := 0
	for i := 0; i < len(A); i++ {
		mini := i
		for j := i; j < len(A); j++ {
			//fmt.Printf("i=%d, j=%d, mini=%d, A[%d]=%d, A[%d]=%d\n", i, j, mini, j, A[j], mini, A[mini])
			if A[j] < A[mini] {
				mini = j
			}
		}
		if i != mini {
			A[i], A[mini] = A[mini], A[i]
			counter += 1
			//printAll(A)
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

	counter := selectionSort(elements)
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
