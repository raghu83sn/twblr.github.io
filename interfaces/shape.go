package interfaces

type Shape interface {
	Area() int
}

type Square struct {
	side int
}

func (sqr Square) Area() int {
	return sqr.side * sqr.side
}

type Rectangle struct {
	length, breadth int
}

func (rec Rectangle) Area() int {
	return rec.length * rec.breadth
}

type Hybrid struct {
	shapes []Shape
}

func (hyb Hybrid) Area() int {
	areaSum := 0
	for _, each_shape := range hyb.shapes {
		areaSum = areaSum + each_shape.Area()
	}
	return areaSum
}
