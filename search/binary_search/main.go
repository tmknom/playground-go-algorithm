package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(S []int, key int) bool {
	if len(S) == 0 {
		return false
	}

	//fmt.Print("search : ")
	//printAll(S)

	// 真ん中の値を取り出す
	centerIndex := len(S) / 2
	centerValue := S[centerIndex]

	// 真ん中とkeyが等しいかチェック
	if centerValue == key {
		//fmt.Printf("matched value : %d\n\n", centerValue)
		return true
	}

	// 真ん中より小さいことが確定した場合はleftValuesに対してsearchを実行
	if centerValue > key {
		//fmt.Print("recursive search : ")
		//printAll(S[0:centerIndex])
		//fmt.Println("")
		return search(S[0:centerIndex], key)
	}

	// 真ん中より大きいことが確定した場合はleftValuesに対してsearchを実行
	if centerValue < key {
		//fmt.Print("recursive search : ")
		//printAll(S[centerIndex+1:])
		//fmt.Println("")
		return search(S[centerIndex+1:], key)
	}

	return false
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
