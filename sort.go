package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}
func (p People) Less(i, j int) bool {
	if p[i].birthDay == p[j].birthDay {
		if p[i].firstName == p[j].firstName {
			return p[i].lastName < p[j].lastName
		} else {
			return p[i].firstName < p[j].firstName
		}
	}
	return p[i].birthDay.After(p[j].birthDay)
}
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Sort interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func main() {
	ivanIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	ivanIvanovDate2, err := time.Parse("2006-Jan-02", "2003-Aug-05")
	artiomIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	artiomIvanovDate2, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	alexIvanovDate2, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p := People{
		{"Ivan", "Ivanov", ivanIvanovDate},
		{"Ivan", "Ivanov", ivanIvanovDate2},
		{"Artiom", "Ivanov", artiomIvanovDate},
		{"Artiom", "Ivanov", artiomIvanovDate2},
		{"Alexandr", "Ivanov", alexIvanovDate2},
	}
	// fmt.Println(p)
	sort.Sort(p)
	fmt.Println(p)
}
