package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeCombination() {

}

var N int
var A []int

func run(index int, value int) bool {
	if value == 0 {
		return true
	}
	if index >= N {
		return false
	}

	return run(index+1, value) || run(index+1, value-A[index])
}

func main() {
	// 初期化
	N = scanNumber()
	A = scanNumberCollection()
	q := scanNumber()
	m := scanNumberCollection()
	fmt.Printf("n=%d, A=%v, q=%d, m=%v\n", N, A, q, m)

	for i := 0; i < q; i++ {
		if run(0, m[i]) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
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
