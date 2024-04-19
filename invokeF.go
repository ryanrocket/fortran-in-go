package main

// #cgo LDFLAGS: bruh.o -L/opt/homebrew/opt/gcc/lib/gcc/13 -lgfortran
// void hello();
import (
	"C"
)

func main() {
	C.hello()
}
