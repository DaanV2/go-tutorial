package main

type Rect struct {
	width, height int
}

func (r *Rect) Area() int {
	return r.width * r.height
}

func (r Rect) Perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := Rect{width: 10, height: 5}
	area := r.Area()
	perim := r.Perim()
}