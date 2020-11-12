package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(S []int, key int) bool {
	result := false
	for _, valueS := range S {
		if key == valueS {
			result = true
			break
		}
	}
	return result
}

func run(S []int, T []int) int {
	result := 0

	for _, valueT := range T {
		if search(S, valueT) {
			result += 1
		}
	}

	return result
}

func main() {
	// 初期化
	skipScan()
	S := scanNumberCollection()
	skipScan()
	T := scanNumberCollection()

	// アルゴリズムの実行
	result := run(S, T)
	fmt.Println(result)
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
