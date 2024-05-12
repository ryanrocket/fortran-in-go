package main

// #cgo CFLAGS: -I${SRCDIR}
// #cgo LDFLAGS: -L${SRCDIR} -Wl,-rpath=\$ORIGIN -lfortran
// void hello();
import (
	"C"
)

func main() {
	C.hello()
}
