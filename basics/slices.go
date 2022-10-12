package basics

import "fmt"

func Slices() {
	fmt.Println("***************************Slices Section Start***************************")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("a %v\n", a)
	fmt.Printf("a Length %v\n", len(a))
	fmt.Printf("a Capacity %v\n", cap(a))

	b := a[:]   // slice of all element
	c := a[3:]  // slice from 3rd element (3 is inclusive)
	d := a[:6]  // slice first 6 element
	e := a[3:6] //slice from 3rd to 6th element
	fmt.Printf("b %v\n", b)
	fmt.Printf("c %v\n", c)
	fmt.Printf("d %v\n", d)
	fmt.Printf("e %v\n", e)

	f := make([]int, 3) // different method to create slice
	g := []int{}        // without initial length
	fmt.Printf("f %v\n", f)
	fmt.Printf("f Length %v\n", len(f))
	fmt.Printf("f Capacity %v\n", cap(f))

	fmt.Printf("g %v\n", g)
	fmt.Printf("g Length %v\n", len(g))
	fmt.Printf("g Capacity %v\n", cap(g))

	f = append(f, 3, 4, 5, 6)

	fmt.Printf("f %v\n", f)
	fmt.Printf("f Length %v\n", len(f))
	fmt.Printf("f Capacity %v\n", cap(f))

	fmt.Println("***************************Slices Section End***************************")

}
