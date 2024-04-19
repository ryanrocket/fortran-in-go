package main

// #cgo LDFLAGS: -L. -lfortran
// void hello();
import (
	"C"
)

func main() {
	C.hello()
}
