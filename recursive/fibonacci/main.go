package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1 1 2 3 5 8 13 21
func fibonacci(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func main() {
	// 初期化
	num := scanNumber()
	for i := 0; i <= num; i++ {
		result := fibonacci(i)
		fmt.Printf("%d ", result)
	}
	fmt.Println()
}

//
// 共通ユーティリティ
//

func printAll(elements []int) {
	result := ""
	for _, value := range elements {
		result += strconv.Itoa(value) + " "
	}
	fmt.Println(strings.Trim(result, " "))
}

func scanNumberCollection() []int {
	ssc := scanStringCollection()
	result := make([]int, len(ssc))
	for i := 0; i < len(ssc); i++ {
		intValue, _ := strconv.Atoi(ssc[i])
		result[i] = intValue
	}
	return result
}

func scanStringCollection() []string {
	return strings.Split(scanString(), " ")
}

func scanNumber() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

func skipScan() {
	scanString()
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

var sc = bufio.NewScanner(os.Stdin)
