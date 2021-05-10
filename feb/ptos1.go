package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pwe", int(*l))
}

func main() {
	pew := laser(2)
	shout(&pew)
	var nowhere *int
	fmt.Println(nowhere)
	fmt.Println(*nowhere)

	// const layout = "Mon, 26 Jan, 2021"
	// day := time.Now()
	// fmt.Println(day.Format(layout))
	// tomorrow := day.Add(24 * time.Hour)
	// fmt.Println(day.Format(layout))
	// fmt.Println(tomorrow.Format(layout))

}
