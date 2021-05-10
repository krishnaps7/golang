package main

import "fmt"

func ptr(pa *int) {
	*pa = 748
}

type employee struct {
	name  string
	empid int
}

func main() {
	var p = 30
	var x = &p
	fmt.Println("Value of p before passing to pointer function: ", *x)
	ptr(x)
	fmt.Println("Value of p after passing to pointer function: ", *x)
	var emp = employee{"krishna", 7}
	var e = &emp
	fmt.Println("Pointer to structure is: ", e)
	fmt.Println("the name and id of employree", e.name, e.empid)
	e.name = "Renuka"
	fmt.Println("employee after name changed ", e)
	// fmt.Println("the p is ", p)
	// fmt.Println("the p address is ", &p)
	// fmt.Println("the x is ", x)
	// fmt.Println("the value in x is ", *x)
	// fmt.Println("the x address is ", &x)
	var v = 100
	var p1 = &v
	var p2 = &p1
	fmt.Println("the v, &v : ", v, &v)
	fmt.Println("the p1, &p1 : ", p1, &p1)
	fmt.Println("the p2, &p2 : ", p2, &p2)
	fmt.Println("v,*p1,**p2", v, *p1, *(*p2))
	*p1 = 300
	fmt.Println("the v, &v : ", v, &v)
	**p2 = 500
	fmt.Println("the v, &v : ", v, &v)
	*p2 = 0xc0000b4010
	fmt.Println("the v, &v : ", v, &v)

}
