package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func run(numbers []int, l int, r int) int {
	var result int
	m := (l + r) / 2
	fmt.Printf("(m=%d, l=%d, r=%d => ", m, l, r)
	if l == r-1 {
		fmt.Printf(" numbers[%d]=%d\n", l, numbers[l])
		result = numbers[l]
	} else {
		left := run(numbers, l, m)
		right := run(numbers, m, r)
		result = int(math.Max(float64(left), float64(right)))
		fmt.Printf("max(left=%d, right=%d) => result=%d)\n", left, right, result)
	}
	//fmt.Printf(")\n")
	return result
}

func main() {
	// 初期化
	inputs := scanNumberCollection()
	fmt.Println(inputs)
	max := run(inputs, 0, len(inputs)-1)
	fmt.Println(max)
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
