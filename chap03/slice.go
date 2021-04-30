package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(a, "length:", len(a), "capacity:", cap(a))
	var b []int
	b = []int{20, 40, 60, 80}
	fmt.Println(b, "length:", len(b), "capacity:", cap(b))
	var c = []int{2, 3, 4, 5}
	c = append(c, 1)
	fmt.Println(c, "length:", len(c), "capacity:", cap(c))
	d := append(a, 100, 200, 300, 400)
	fmt.Println(d)
	e := append(a, b...)
	fmt.Println(e)
	a = append(a, 60)
	fmt.Println(a)
	fmt.Println("Length of a:", len(a))
	fmt.Println("Capacity of a:", cap(a))

	// Capacity of slice
	// var x []int
	// fmt.Println(x, "Length:", len(x), "Capacity:", cap(x))
	// x = append(x, 10)
	// fmt.Println(x, "Length:", len(x), "Capacity:", cap(x))
	// x = append(x, 20)
	// fmt.Println(x, "Length:", len(x), "Capacity:", cap(x))
	// x = append(x, 30)
	// fmt.Println(x, "Length:", len(x), "Capacity:", cap(x))
	// x = append(x, 40)
	// fmt.Println(x, "Length:", len(x), "Capacity:", cap(x))
	// x = append(x, 50)
	// fmt.Println(x, "Length:", len(x), "Capacity:", cap(x))

	x := []int{1, 2, 3, 4}
	fmt.Println("x slice:", x)
	y := x[:2]
	fmt.Println("y slice from index 0 upto 2:", y)
	z := x[1:]
	fmt.Println("z slice from index 1 to end:", z)
	g := x[1:3]
	fmt.Println("g slice from index 1 upto 3:", g)
	h := x[:]
	fmt.Println("h slice from start to end:", h)

	i := []int{1, 2, 3, 4}
	j := i[:2]
	k := i[1:]
	i[1] = 20
	j[0] = 10
	k[1] = 30
	fmt.Println("i:", i)
	fmt.Println("j:", j)
	fmt.Println("k:", k)

	// Confusing
	// Slicing slice while apending
	l := []int{1, 2, 3, 4}
	m := l[:2]
	fmt.Println("Cap:", cap(l), cap(m))
	m = append(m, 30)
	fmt.Println("l:", l)
	fmt.Println("m:", m)

	// More Confusing
	r := make([]int, 0, 5)
	r = append(r, 1, 2, 3, 4)
	s := r[:2] // 1, 2
	t := r[2:] // 3, 4
	fmt.Println("r:", cap(r), "s:", cap(s), "t:", cap(t))
	s = append(s, 30, 40, 50) // 1, 2, 30, 40, 50
	r = append(r, 60)         // 1, 2, 60, 40, 50
	t = append(t, 70)         // 30, 40, 70, 60
	fmt.Println("r:", r)
	fmt.Println("s:", s)
	fmt.Println("t:", t)
	fmt.Println("r after modifications:", r)

	u := r[:2:2]
	fmt.Println("u from r:", u)
	v := r[2:4:4]
	fmt.Println("v from r:", v)
	u = append(u, 3, 4, 5)
	fmt.Println("u:", u, "cap of u:", cap(u), "cap of r:", cap(r))
}
