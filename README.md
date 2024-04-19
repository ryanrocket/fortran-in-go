# FORTRAN in Golang
Ryan Wans

---

## Summary
We can invoke FORTRAN code, such as subroutines and functions, from Golang using the `cgo` package. This method is by far the cleanest to do so, but is still somewhat tricky and messy to get working. The idea is to compile our FORTRAN code into a shared object with relative memory positioning (`-fPIC`) and allow Go to dynamically link it at compilation time. The method follows. 

## Steps
1. Write the FORTRAN code in a file, say `fortran.f90`. We need our subroutines and functions to use `bind(C)` to make them compatible with C. Optionally, we can use `iso_c_binding` to make the code more readable for C types. 
2. Compile the FORTRAN and obtain an object file: `gfortran -c fortran.f90`
3. In our Goland code, we need to import the `C` package and include the necessary comments to link the FORTRAN code. This helps `cgo` to link our object to the Golang code. For example, we could include `// #cgo LDFLAGS: fortran.o -L/path/to/gcc/lib/gcc/version -lgfortran`. The start of the path was obtained via `brew --prefix gcc`.
4. We need to add an additional commend for each subroutine of function we want to use. For example, we could include `// void hello();` to use the `hello` subroutine.
5. To use the subroutine of function, we simply call it as if it were a C function. For example, we could call `C.hello()`.
6. We can now build and run our Golang code.

## Alternative (Better) Linking
Instead of linking to the absolute path of the GCC library, we can create a shared object file from the fortran source using `gfortran -shared -fPIC -o libfortran.so fortran.f90`. We can then link to this shared object file using `// #cgo LDFLAGS: -L. -lfortran`.

This method is cleaner and more portable, but requires at least the relative path to the shared object file. If we are using multiple fortran files, we can compile them all into a single shared object file using `gfortran -shared -fPIC -o libfortran.so fortran1.f90 fortran2.f90 ...`.
