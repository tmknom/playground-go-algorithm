package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run(num int) int {
	if num == 0 {
		return 1
	}
	return num * run(num-1)
}

func main() {
	// 初期化
	n := scanNumber()
	for i := 0; i < n; i++ {
		num := scanNumber()
		result := run(num)
		fmt.Printf("num=%d, result=%d\n", num, result)
	}
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
