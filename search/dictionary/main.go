package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var Dictionary = [2000000000]bool{}

func hash(key string) int {
	result := 0
	chars := strings.Split(key, "")

	for i := len(chars) - 1; i >= 0; i-- {
		result += hashChar(chars[i]) + (4 * i) - i
		//fmt.Printf("i=%d, char[i]=%s, hash=%d, result=%d\n", i, chars[i], hashChar(chars[i]), result)
	}

	return result
}

func hashChar(char string) int {
	if char == "A" {
		return 1
	}
	if char == "C" {
		return 2
	}
	if char == "G" {
		return 3
	}
	if char == "T" {
		return 4
	}
	return math.MinInt32
}

func insert(key string) {
	//fmt.Printf("insert => %s\n", key)
	hashKey := hash(key)
	Dictionary[hashKey] = true
}

func find(key string) {
	//fmt.Printf("find => %s\n", key)
	hashKey := hash(key)
	if Dictionary[hashKey] {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func run(command string, key string) {
	if command == "insert" {
		insert(key)
	} else {
		find(key)
	}
}

func main() {
	// 初期化
	n := scanNumber()
	for i := 0; i < n; i++ {
		line := scanStringCollection()
		command := line[0]
		key := line[1]
		run(command, key)
	}

	// アルゴリズムの実行
	//result := run(S, T)
	//fmt.Println(result)
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
