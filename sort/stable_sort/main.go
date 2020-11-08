package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(A []string) {
	n := len(A)

	for i := 0; i <= n-1; i++ {
		for j := n - 1; j > i; j-- {
			numAj, _ := strconv.Atoi(A[j][1:])
			numAj1, _ := strconv.Atoi(A[j-1][1:])
			if numAj1 > numAj {
				A[j-1], A[j] = A[j], A[j-1]
			}
		}
	}

	printAll(A)
}

func selectionSort(A []string) {
	for i := 0; i < len(A); i++ {
		mini := i
		for j := i; j < len(A); j++ {
			//fmt.Printf("i=%d, j=%d, mini=%d, A[%d]=%d, A[%d]=%d\n", i, j, mini, j, A[j], mini, A[mini])
			numAj, _ := strconv.Atoi(A[j][1:])
			numAmini, _ := strconv.Atoi(A[mini][1:])
			if numAj < numAmini {
				mini = j
			}
		}
		if i != mini {
			A[i], A[mini] = A[mini], A[i]
		}
	}

	printAll(A)
}

func isStable(A []string, B []string) bool {
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			return false
		}
	}
	return true
}

func main() {
	// 初期化
	number, _ := strconv.Atoi(nextLine())
	inputs := strings.Split(nextLine(), " ")
	elements := make([]string, number)
	elements2 := make([]string, number)
	for i := 0; i < len(inputs); i++ {
		elements[i] = inputs[i]
	}
	copy(elements2, elements)

	bubbleSort(elements)
	fmt.Println("Stable")
	selectionSort(elements2)
	if isStable(elements, elements2) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

func printAll(elements []string) {
	result := ""
	for _, value := range elements {
		result += value + " "
	}
	fmt.Println(strings.Trim(result, " "))
}

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
