package main


type Rectangle struct {
	height float64
	width float64

	
}

func Perimeter(rect Rectangle) float64 {
	return  2*(rect.height+rect.width)
}

func Area(rect Rectangle)float64 {
	return rect.height*rect.width
}