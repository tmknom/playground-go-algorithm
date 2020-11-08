package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printAll(elements []int) {
	result := ""
	for _, value := range elements {
		result += strconv.Itoa(value) + " "
	}
	fmt.Println(strings.Trim(result, " "))
}

//5 2 4 6 1 3
//2 5 4 6 1 3
//2 4 5 6 1 3
//2 4 5 6 1 3
//1 2 4 5 6 3
//1 2 3 4 5 6
func insertionSort(elements []int, number int) {
	for sortingElementIndex := 1; sortingElementIndex < number; sortingElementIndex++ {
		sortingElement := elements[sortingElementIndex]
		sortedIndex := sortingElementIndex - 1
		for {
			if sortedIndex < 0 {
				break
			}

			if elements[sortedIndex] <= sortingElement {
				break
			}

			elements[sortedIndex+1] = elements[sortedIndex]
			sortedIndex--
		}
		elements[sortedIndex+1] = sortingElement
		printAll(elements)
	}
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

	//fmt.Println(number)
	//fmt.Printf("inputs = %v\n\n", elements)

	printAll(elements)
	insertionSort(elements, number)
}

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
