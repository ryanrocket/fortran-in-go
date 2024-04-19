subroutine hello() bind(C)
        use iso_c_binding
        implicit none
        print *, "Hello from Fortran"
end subroutine hello
