package main

// #cgo CFLAGS: -I${SRCDIR}
// #cgo LDFLAGS: -L${SRCDIR} -Wl,-rpath=\$ORIGIN -lfortran
// int multiply(int, int);
// void arrayByHundred(double*, int*);
// void arrayWoBind(double*, int*);
import "C"
import "fmt"

func main() {
	a := 77
	b := 9
	fmt.Println("Multiply: ", C.multiply(C.int(a), C.int(b)))

	array := []C.double{1.0, 2.0, 3.0, 4.0, 5.0}
	size := C.int(len(array))

	// Uses iso_c_binding
	C.arrayByHundred(&array[0], &size)
	fmt.Println("Array by hundred: ", array)

	// Does not use iso_c_binding
	C.arrayWoBind(&array[0], &size)
	fmt.Println("Array without binding: ", array)
}
