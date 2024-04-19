build: fortran.f95 main.go
	gfortran -shared -fPIC -o libfortran.so fortran.f95
	go build -ldflags "-s -w -buildid=" main.go

build-noSoName: fortran.f95 main.go
	gfortran -shared -fPIC -o liba.so fortran.f95
	# sed -i 's/-lfortran/-la/g' ./main.go
	sed 's/-lfortran/-la/g' main.go > main.go.tmp && mv main.go.tmp main.go
	go build -ldflags "-s -w -buildid=" main.go
