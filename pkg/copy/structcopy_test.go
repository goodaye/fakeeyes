package copy

import (
	"fmt"
	"testing"
)

func TestStructCopy(t *testing.T) {

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
