package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	Root *Node
}

type Node struct {
	key    int
	parent *Node
	left   *Node
	right  *Node
}

func (t *Tree) insert(z *Node) {
	fmt.Printf("\ninsert %d\n", z.key)
	var parent *Node
	x := t.Root
	for {
		if x == nil {
			break
		}
		parent = x
		if z.key < x.key {
			fmt.Printf("x = x.left = %v\n", x.left)
			x = x.left
		} else {
			fmt.Printf("x = x.right = %v\n", x.right)
			x = x.right
		}
		fmt.Printf("parent=%d, x=%d\n", parent.key, z.key)
	}
	z.parent = parent
	if parent != nil {
		fmt.Printf("z.parent = %d\n", z.parent.key)
	}

	if parent == nil {
		t.Root = z
		fmt.Printf("root = %d\n", z.key)
	} else if z.key < parent.key {
		parent.left = z
		fmt.Printf("z.key < parent.key => parent.left = %d\n", z.key)
	} else {
		parent.right = z
		fmt.Printf("z.key >= parent.key => parent.right = %d\n", z.key)
	}
}

func (t *Tree) print() {
	inorder(t.Root)
	println("")
	preorder(t.Root)
}

func preorder(n *Node) {
	fmt.Printf("%d ", n.key)
	if n.left != nil {
		preorder(n.left)
	}
	if n.right != nil {
		preorder(n.right)
	}
}

func inorder(n *Node) {
	if n.left != nil {
		inorder(n.left)
	}
	fmt.Printf("%d ", n.key)
	if n.right != nil {
		inorder(n.right)
	}
}

func main() {
	// 初期化
	m := scanNumber()

	tree := &Tree{}
	for i := 0; i < m; i++ {
		inputs := scanStringCollection()
		if len(inputs) == 1 {
			tree.print()
		} else {
			key, _ := strconv.Atoi(inputs[1])
			z := &Node{key: key}
			tree.insert(z)
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
