package basics

import (
	"fmt"
)

func Primitives() {
	fmt.Println("--------------BOOLEAN------------")
	// boolean
	var boolVar bool = true
	fmt.Printf("%v %T\n", boolVar, boolVar)

	n := 1 == 1
	m := 2 == 1
	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T\n", m, m)

	var defaultBool bool
	fmt.Printf("%v %T\n", defaultBool, defaultBool)

	fmt.Println("--------------INT------------")

	var int8Var int8 = 42   // -128 -> 127
	var int16Var int16 = 42 // -32768 -> 32767
	var int32Var int32 = 42 // -2147483648 -> 2147483648
	var int64Var int64 = 42 // -9 trillion -> 9trillion
	fmt.Printf("%v %T\n", int8Var, int8Var)
	fmt.Printf("%v %T\n", int16Var, int16Var)
	fmt.Printf("%v %T\n", int32Var, int32Var)
	fmt.Printf("%v %T\n", int64Var, int64Var)

	var uInt8 uint8 = 42   // 0 - 255
	var uInt16 uint16 = 42 // 0 - 65535
	fmt.Printf("%v %T\n", uInt8, uInt8)
	fmt.Printf("%v %T\n", uInt16, uInt16)

	fmt.Println("--------------Calc------------")

	a := 10
	b := 3
	fmt.Printf("%v %T\n", (a + b), (a + b))
	fmt.Printf("%v %T\n", (a - b), (a - b))
	fmt.Printf("%v %T\n", (a * b), (a * b))
	fmt.Printf("%v %T\n", (a / b), (a / b))
	fmt.Printf("%v %T\n", (a % b), (a % b))

	var c int8 = 4
	//fmt.Printf("%v %T\n", (a % c), (a % c)) // invalida operation becuase of different type int and int8
	fmt.Printf("%v %T\n", (a % int(c)), (a % int(c)))

	fmt.Println("--------------Binary Operator------------")
	fmt.Printf("%v %T\n", (a & b), (a & b))
	fmt.Printf("%v %T\n", (a | b), (a | b))
	fmt.Printf("%v %T\n", (a ^ b), (a ^ b))
	fmt.Printf("%v %T\n", (a &^ b), (a &^ b))
	fmt.Printf("%v %T\n", (a | ^b), (a | ^b))

	var d int = 8
	fmt.Printf("%v %T\n", (d << 3), (d << 3))
	fmt.Printf("%v %T\n", (d >> 3), (d >> 3))

	fmt.Println("--------------Float------------")
	var floatVar float32 = 3.14
	fmt.Printf("%v %T\n", floatVar, floatVar)
	floatVar = 13.7e10
	fmt.Printf("%v %T\n", floatVar, floatVar)
	floatVar = 2.14e12
	fmt.Printf("%v %T\n", floatVar, floatVar)
	// floatVar = 2.14e72 // overflow
	var floatVar64 float64 = 2.14e72
	fmt.Printf("%v %T\n", floatVar64, floatVar64)

	fmt.Println("--------------Complex------------")

	var ncomplex complex64 = 2 + 5i // https://en.wikipedia.org/wiki/Complex_number
	fmt.Printf("%v %T\n", ncomplex, ncomplex)
	fmt.Printf("%v %T\n", real(ncomplex), real(ncomplex))
	fmt.Printf("%v %T\n", imag(ncomplex), imag(ncomplex))

	fmt.Println("--------------String------------")
	var stringVar string = "this is a string"
	fmt.Printf("%v %T\n", stringVar, stringVar)
	fmt.Printf("%v %T\n", stringVar[2], stringVar[2])
	fmt.Printf("%v %T\n", string(stringVar[2]), string(stringVar[2]))
	var secondStringVar string = "another string"
	fmt.Printf("%v %T\n", stringVar+secondStringVar, stringVar+secondStringVar)

	byteVar := []byte(stringVar)
	fmt.Printf("%v %T\n", byteVar, byteVar)

	var rune rune = 'C'
	fmt.Printf("%v %T\n", rune, rune)

}
