package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//var A []float64

var Count = 0

func merge(A []float64, leftIndex int, middleIndex int, rightIndex int) {
	//fmt.Printf("merge(leftIndex: %d, middleIndex, %d, rightIndex: %d), ", leftIndex, middleIndex, rightIndex)
	sizeL := middleIndex - leftIndex
	sizeR := rightIndex - middleIndex

	var L []float64
	for i := 0; i < sizeL; i++ {
		L = append(L, A[leftIndex+i])
	}
	L = append(L, math.Inf(0))

	var R []float64
	for i := 0; i < sizeR; i++ {
		R = append(R, A[middleIndex+i])
	}
	R = append(R, math.Inf(0))

	fmt.Printf("run merge L=%v, R=%v => ", L, R)
	i := 0
	j := 0
	for k := leftIndex; k < rightIndex; k++ {
		Count += 1
		if L[i] <= R[j] {
			A[k] = L[i]
			i += 1
		} else {
			A[k] = R[j]
			j += 1
		}
	}
	fmt.Printf("%v\n", A[leftIndex:rightIndex])
	//fmt.Printf("{L=%v}(by merge)\n", A[:middleIndex], A[middleIndex:], middleIndex)
}

func mergeSort(A []float64, leftIndex int, rightIndex int) {
	fmt.Printf("mergeSort(A[leftIndex: %d, rightIndex: %d] => %v)", leftIndex, rightIndex, A[leftIndex:rightIndex+1])
	if leftIndex+1 < rightIndex {
		middleIndex := int(math.Ceil((float64(leftIndex) + float64(rightIndex)) / 2))
		//middleIndex := (leftIndex + rightIndex) / 2
		fmt.Printf(", middleIndex=%d\n", middleIndex)
		//fmt.Printf("{L=%v, R=%v}(middleIndex=%d)\n", A[:middleIndex], A[middleIndex:], middleIndex)

		fmt.Print("Divide left: ")
		mergeSort(A, leftIndex, middleIndex-1)
		fmt.Print("Divide right: ")
		mergeSort(A, middleIndex, rightIndex)
		merge(A, leftIndex, middleIndex, rightIndex+1)
	} else {
		Count += 1
		if A[leftIndex] > A[rightIndex] {
			A[leftIndex], A[rightIndex] = A[rightIndex], A[leftIndex]
		}
		fmt.Printf(" => sorted %v\n", A[leftIndex:rightIndex+1])
		//fmt.Printf("mergeSort skiped: leftIndex+1=%d, A[leftIndex+1]=%f\n", leftIndex+1, A[leftIndex+1])
	}
}

func main() {
	// 初期化
	skipScan()
	A := scanNumberCollection()

	fmt.Printf("intput A = %v\n\n", A)
	mergeSort(A, 0, len(A)-1)
	fmt.Printf("sorted A = %v\n", A)
	fmt.Println(Count)
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

func scanNumberCollection() []float64 {
	ssc := scanStringCollection()
	result := make([]float64, len(ssc))
	for i := 0; i < len(ssc); i++ {
		intValue, _ := strconv.Atoi(ssc[i])
		result[i] = float64(intValue)
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
