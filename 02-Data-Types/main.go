package main

//Numbers start

//8-bit signed integer (-128 to 127)
var xInt8 int8 = 8

//16-bit signed integer (-32,768 to 32,767)
var xInt16 int16 = 16

//32-bit signed integer (-2,147,483,648 to 2,147,483,647)
var xInt32 int32 = 32

//64-bit signed integer (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)
var xInt64 int64 = 64

//8-bit unsigned integer (0 to 255)
var xUInt8 uint8 = 8

//16-bit unsigned integer (0 to 65,535)
var xUInt16 uint16 = 16

//32-bit unsigned integer (0 to 4,294,967,295)
var xUInt32 uint32 = 32

//64-bit unsigned integer (0 to 18,446,744,073,709,551,615)
var xUInt64 uint64 = 64

//Both in and uint contain same size, either 32 or 64 bit (-2,147,483,648 to 2,147,483,647)
var xInt int = 1

//Both in and uint contain same size, either 32 or 64 bit. (0 to 4,294,967,295)
var xUInt uint = 1

//It is a synonym of int32 and also represent Unicode code points
var xRune rune = 1

//It is a synonym of int8.
var xByte byte = 3

//It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.
var xUIntPtr uintptr = 31

//32-bit IEEE 754 floating-point number (-3.4E+38 to +3.4E+38)
var xFloat32 float32 = 3.1

//64-bit IEEE 754 floating-point number (-1.7E+308 to +1.7E+308)
var xFloat64 float64 = 3.11

//Complex numbers which contain float32 as a real and imaginary component.

var XComplex64 complex64 = (1 + 1i)

//Complex numbers which contain float64 as a real and imaginary component.

var XComplex128 complex128 = (2 + 2i)

//Numbers End

var xString string = "string"

var xBool bool = true

func main() {

}
