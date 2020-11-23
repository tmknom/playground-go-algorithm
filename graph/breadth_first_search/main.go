package main

import "fmt"

type Queue struct {
	names         []string
	searchedNames []string
}

func (q *Queue) enqueueSlice(names []string) {
	q.names = append(q.names, names...)
}

func (q *Queue) enqueue(name string) {
	q.names = append(q.names, name)
}

func (q *Queue) dequeue() string {
	result := q.names[0]
	q.names = q.names[1:]
	return result
}

func (q *Queue) isEmpty() bool {
	return len(q.names) == 0
}

func (q *Queue) search() bool {
	for {
		if q.isEmpty() {
			break
		}

		name := q.dequeue()
		if isSeller(name) {
			fmt.Printf("\nseller is %s!\n", name)
			return true
		}
		if q.isSearched(name) {
			continue
		}
		fmt.Printf("%s is not seller\n", name)
		q.enqueueSlice(graph[name])
		q.addSearched(name)
	}
	return false
}

func (q *Queue) addSearched(name string) {
	q.searchedNames = append(q.searchedNames, name)
}

func (q *Queue) isSearched(name string) bool {
	for _, searchedName := range q.searchedNames {
		if searchedName == name {
			return true
		}
	}
	return false
}

func isSeller(name string) bool {
	return name == "thom"
}

func main() {
	graph = initGraph()

	queue := &Queue{names: []string{}, searchedNames: []string{}}
	queue.enqueueSlice(graph["you"])
	queue.search()
	//fmt.Println(queue.dequeue())
	//fmt.Println(queue.dequeue())
	//fmt.Println(queue.dequeue())
}

var graph = map[string][]string{}

func initGraph() map[string][]string {
	graph := map[string][]string{}
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	return graph
}
