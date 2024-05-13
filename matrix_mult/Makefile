build: multiply.f90 main.go 
	gfortran -shared -fPIC -o libmultiply.so multiply.f90
	mv libmultiply.so ./lib/
	go build -ldflags "-s -w -buildid=" main.go