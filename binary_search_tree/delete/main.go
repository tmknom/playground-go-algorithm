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
	var parent *Node
	x := t.Root
	for {
		if x == nil {
			break
		}
		parent = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = parent

	if parent == nil {
		t.Root = z
	} else if z.key < parent.key {
		parent.left = z
	} else {
		parent.right = z
	}
}

func (t *Tree) print() {
	inorder(t.Root)
	println("")
	preorder(t.Root)
}

func find(n *Node, key int) *Node {
	if n == nil {
		return nil
	}

	if n.key == key {
		return n
	}
	if key < n.key && n.left != nil {
		return find(n.left, key)
	}
	if key >= n.key && n.right != nil {
		return find(n.right, key)
	}
	return nil
}

func (t *Tree) delete(key int) {
	node := find(t.Root, key)
	parent := node.parent

	// 子を持たない場合は単純に指定されたキーのノードを削除
	if node.left == nil && node.right == nil {
		if node.key < parent.key {
			parent.left = nil
		} else {
			parent.right = nil
		}
		node = nil
		return
	}

	// 2つ子を持つ場合は指定されたキーのノードを削除して、子を昇格
	if node.left != nil && node.right != nil {
		if node.key < parent.key {
			parent.left = node.right
			node.right.parent = parent
			node.left.parent = node.right
			node.right.left = node.left
		} else {
			parent.right = node.right
			node.right.parent = parent
			node.left.parent = node.right
			node.right.left = node.left
		}
		node = nil
		return
	}

	if node.left != nil {
		parent.left = node.left
		node.left.parent = parent
	}

	if node.right != nil {
		parent.right = node.right
		node.right.parent = parent
	}
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
			fmt.Println("")
		} else if inputs[0] == "insert" {
			key, _ := strconv.Atoi(inputs[1])
			z := &Node{key: key}
			tree.insert(z)
		} else if inputs[0] == "delete" {
			key, _ := strconv.Atoi(inputs[1])
			tree.delete(key)
		} else {
			key, _ := strconv.Atoi(inputs[1])
			result := find(tree.Root, key)
			if result != nil {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
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
