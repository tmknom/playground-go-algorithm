package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Element struct {
	value int
	next  *Element
}

type LinkedList struct {
	head   *Element
	length int
}

func newLinkedList() *LinkedList {
	return &LinkedList{
		head:   &Element{},
		length: 0,
	}
}

// h -> nil
// h -> 1 -> nil
// h -> 1 ->2 -> nil
// h -> 1 -> 3 ->2 -> nil
// h -> 1 ->2 -> nil
// h -> 1 -> nil
// h -> nil
func (l *LinkedList) remove(index int) {
	prev := l.head
	if index > 0 {
		prev = l.get(index - 1)
	}

	current := prev.next
	next := current.next

	prev.next = next
	current.next = nil
	l.length -= 1
}

func (l *LinkedList) set(index int, value int) {
	current := l.get(index)
	current.value = value
}

func (l *LinkedList) add(index int, value int) {
	current := l.head
	l.length += 1

	if index == 0 {
		new := &Element{
			value: value,
			next:  current.next,
		}
		current.next = new
		return
	}

	prev := l.get(index - 1)
	next := prev.next

	new := &Element{
		value: value,
		next:  next,
	}
	prev.next = new
}

func (l *LinkedList) get(index int) *Element {
	current := l.head.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current
}

func (l *LinkedList) size() int {
	return l.length
}

func (l *LinkedList) isEmpty() bool {
	return l.head.next == nil
}

func (s *LinkedList) string() string {
	result := ""
	current := s.head
	for {
		if current.next == nil {
			return result
		}
		current = current.next
		result += fmt.Sprintf("%d\n", current.value)
	}
}

func main() {
	line := nextLine()
	inputs := strings.Split(line, " ")

	// 初期化
	list := newLinkedList()
	for _, value := range inputs {
		intValue, _ := strconv.Atoi(value)
		list.add(0, intValue)
	}
	list.set(5, 100)
	list.remove(2)

	fmt.Println(list.string())
	fmt.Printf("size = %d\n", list.size())
}

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
