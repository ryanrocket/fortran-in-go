package main

/*
#include <stdlib.h>
#cgo CFLAGS: -I${SRCDIR}/lib
#cgo LDFLAGS: -L${SRCDIR}/lib -Wl,-rpath=\$ORIGIN/lib -lmultiply
struct Matrix {
	int rows;
	int cols;
	double *data;
};
struct Matrix multiply(struct Matrix, struct Matrix);
*/
import "C"
import (
	"fmt"
	"math/rand"
	"unsafe"
)

// Define matrix data type
type Matrix struct {
	rows int
	cols int
	data []float64 // Use C.double to match the C data type
}

// Create a new, empty matrix
func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		rows: rows,
		cols: cols,
		data: make([]float64, rows*cols), // Use C.double to match the C data type
	}
}

// Populate a matrix with random integers between 0 and 9
func (m *Matrix) Randomize() {
	for i := 0; i < m.rows*m.cols; i++ {
		randInt := rand.Intn(10)
		m.data[i] = float64(randInt % 10)
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
	if a.cols != b.rows {
		panic("Invalid matrix dimensions for multiplication!")
	}

	c := C.multiply(C.struct_Matrix{
		rows: C.int(a.rows),
		cols: C.int(b.cols),
		data: (*C.double)(unsafe.Pointer(&a.data[0])),
	}, C.struct_Matrix{
		rows: C.int(b.rows),
		cols: C.int(b.cols),
		data: (*C.double)(unsafe.Pointer(&b.data[0])),
	})

	return &Matrix{
		rows: int(c.rows),
		cols: int(c.cols),
		data: (*[1 << 30]float64)(unsafe.Pointer(c.data))[: c.rows*c.cols : c.rows*c.cols], // Convert C array to Go slice
	}
}

// Free allocated memory
func (m *Matrix) Free() {
	C.free(unsafe.Pointer(&m.data[0]))
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

	// Multiply matrices
	c := Multiply(a, b)

	// Print result
	fmt.Println("Matrix C (AxB):")
	c.Print()
	c.Free()
}
