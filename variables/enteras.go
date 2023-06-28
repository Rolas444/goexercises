package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Print("intcomun = ", intcomun)
	fmt.Print("intde32 = ", intde32)
	fmt.Print("intde64 = ", intde64)
}
