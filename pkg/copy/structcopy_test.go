package copy

import (
	"fmt"
	"testing"
)

type X struct {
	A int
	B float32
	C string
}

type Y struct {
	A int
	B float32
	C string
	D string
}

func TestStructCopy(t *testing.T) {

	x1 := X{
		A: 0,
		B: 12.0,
		C: "i am source ",
	}
	y1 := Y{
		A: 5,
	}

	StructCopy(x1, &y1)
	fmt.Println(y1)
}

func TestStructArrayCopy(t *testing.T) {

	var x1 = []X{
		{
			A: 0,
			B: 12.0,
			C: "i am 1 ",
		},
		{
			A: 10,
			B: 11.0,
			C: "i am 2 ",
		},
	}
	var y1 = []Y{}

	StructSliceCopy(x1, &y1)

	fmt.Println("copy:")
	fmt.Println(y1)

}
