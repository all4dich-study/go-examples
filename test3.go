package main

import (
	"fmt"
	"sort"
	//"os"
)

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) String() string {
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]" + ":mecha"
}

func typeName(v interface{}) string {

	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
func main() {
	/*
		a := true
		switch t := a.(type) {
		default:
			fmt.Println("default , %T\n", t)
		case int:
			fmt.Println("Integer\n")
		}
	*/
	var a Sequence = Sequence{5, 4, 3, 2, 1}
	fmt.Println(a)
	fmt.Println(typeName(a))
	/*
		switch a := a.(type) {
		default:
			fmt.Println("default")
		case Sequence:
			fmt.Println("Catch Type")
		}
	*/

}
