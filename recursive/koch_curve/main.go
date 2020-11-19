package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x float64
	y float64
}

var Start = &Point{x: 0.0, y: 0.0}
var End = &Point{x: 100.0, y: 0.0}

func run(n int, startPoint *Point, endPoint *Point) (p1 *Point, p2 *Point) {
	p1, p2 = startPoint, endPoint
	fmt.Println(startPoint)
	if n == 0 {
		//return p1, p2 = startPoint, endPoint
	}

	p1, p2 = startPoint, endPoint
	if n == 1 {
		size := math.Pow(math.Pow(startPoint.x-endPoint.x, 2)+math.Pow(startPoint.y-endPoint.y, 2), 0.5)
		divideSize := size / 3

		s := &Point{
			x: startPoint.x + divideSize,
			y: startPoint.y + divideSize,
		}
		fmt.Printf("size = %f\n", divideSize)
		fmt.Printf("u.y = s.y + divideSize = %f + %f\n", s.y, divideSize)
		u := &Point{
			x: s.x + divideSize/2,
			y: s.y + divideSize,
		}
		t := &Point{
			x: startPoint.x + 2*divideSize,
			y: startPoint.y + 2*divideSize,
		}
		//fmt.Printf("s: %v t: %v, u: %v\n", s, u, t)
		fmt.Println(s)
		fmt.Println(u)
		fmt.Println(t)
	}
	fmt.Println(endPoint)

	return p1, p2
}

func main() {
	// 初期化
	//n := scanNumber()

	run(1, Start, End)
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
