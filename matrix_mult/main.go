package main

/*
#include <stdlib.h>
#cgo CFLAGS: -I${SRCDIR}/lib
#cgo LDFLAGS: -L${SRCDIR}/lib -Wl,-rpath=\$ORIGIN/lib -lmultiply
int tester(int, int);
void multiply(int, int, double*, int, int, double*, double*);
*/
import "C"
import (
	"fmt"
	"math/rand"
)

// Define matrix data type
type Matrix struct {
	rows int
	cols int
	data []C.double // Use C.double to match the C data type
}

// Create a new, empty matrix
func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		rows: rows,
		cols: cols,
		data: make([]C.double, rows*cols), // Use C.double to match the C data type
	}
}

// Populate a matrix with random integers between 0 and 9
func (m *Matrix) Randomize() {
	for i := 0; i < m.rows*m.cols; i++ {
		randInt := rand.Intn(10)
		m.data[i] = C.double(randInt % 10)
	}
}

// Print a matrix
func (m *Matrix) Print() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Print(int(m.data[i*m.cols+j]), " ")
		}
		fmt.Println()
	}
}

// Wrapper for C.multiply function
func Multiply(a, b *Matrix) *Matrix {
	// Create a new matrix to store the result
	result := NewMatrix(a.rows, b.cols)

	// Call C.multiply function
	C.multiply(C.int(a.rows), C.int(b.cols), &a.data[0], C.int(a.rows), C.int(b.cols), &b.data[0], &result.data[0])

	return result
}

// Create two matrices, populate them, and print them
func main() {
	// Create two matrices
	a := NewMatrix(2, 3)
	b := NewMatrix(3, 2)

	// Populate matrices
	a.Randomize()
	b.Randomize()

	// Print matrices
	fmt.Println("Matrix A:")
	a.Print()
	fmt.Println("Matrix B:")
	b.Print()

	// Test fortran implementation
	// out := C.tester(C.int(a.rows), C.int(b.cols))
	// fmt.Println("Fortran test: ", out)

	// Print address of matrix data
	addrA := fmt.Sprintf("%p", &a.data[0])
	addrB := fmt.Sprintf("%p", &b.data[0])

	fmt.Println("Matrix A data address: ", addrA)
	fmt.Println("Matrix B data address: ", addrB)

	// Multiply matrices
	result := Multiply(a, b)

	// Print result
	fmt.Println("Result:")
	result.Print()
}
