package main

// #cgo LDFLAGS: -L. -lfortran
// int multiply(int, int);
// void arrayByHundred(double*, int*);
import "C"
import "fmt"

func main() {
	a := 77
	b := 9
	fmt.Println("Multiply: ", C.multiply(C.int(a), C.int(b)))

	array := []C.double{1.0, 2.0, 3.0, 4.0, 5.0}
	size := C.int(len(array))
	C.arrayByHundred(&array[0], &size)
	fmt.Println("Array by hundred: ", array)
}
