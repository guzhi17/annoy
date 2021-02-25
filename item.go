package annoy

import "math"

type itemId int


//n dimension
type Point []float64
type PointAble interface {
	Dimension() int
	Point() Point
}

type PointArray interface {
	Len()int
	At(int)PointAble
}

type item struct {
	id     itemId
	vector PointAble//[]float64
}



type PointN []float64

func (p PointN)Dimension() int {
	return len(p)
}
func (p PointN)Point() Point {
	return Point(p)
}

type PointNArray []PointN

func (a PointNArray)Len()int {
	return len(a)
}
func (a PointNArray)At(i int)PointAble {
	return a[i]
}



type PointR3 struct {
	pt Point
	lat, lon float64
}

func (p PointR3)Dimension() int {
	return 3
}
func (p PointR3)Point() Point {
	return p.pt
}

func NewPointR3(latitude, longitude float64) (r PointR3) {
	pt := R3OfLocation(latitude, longitude)
	return PointR3{
		pt: pt[:],
		lat: latitude,
		lon: longitude,
	}
}


const(
	RadFactor 		= 0.017453292519943295
)
const(
	AxisX = iota
	AxisY
	AxisZ
	AxisEnd
)

type Point3 [AxisEnd]float64
func R3OfLocation(latitude, longitude float64) (pt Point3) {
	latitude, longitude = latitude*RadFactor, longitude*RadFactor
	pt[AxisZ] = math.Sin(latitude)
	r := math.Cos(latitude)
	pt[AxisY] = r*math.Sin(longitude)
	pt[AxisX] = r*math.Cos(longitude)
	return
}


func PointR3OfLocation(latitude, longitude float64) (pt Point3) {
	latitude, longitude = latitude*RadFactor, longitude*RadFactor
	pt[AxisZ] = math.Sin(latitude)
	r := math.Cos(latitude)
	pt[AxisY] = r*math.Sin(longitude)
	pt[AxisX] = r*math.Cos(longitude)
	return
}




type PointR3Array []PointR3

func (a PointR3Array)Len()int {
	return len(a)
}
func (a PointR3Array)At(i int)PointAble {
	return a[i]
}