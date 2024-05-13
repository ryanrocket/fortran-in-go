module MatrixModule
    use iso_c_binding
    implicit none

contains

    function tester( a, b ) result( c ) bind(C, name="tester")
        use iso_c_binding
        implicit none

        integer(c_int), intent(in), value :: a, b
        integer(c_int) :: c

        print *, "Hello from tester"

        c = a + b
    end function tester

    subroutine multiply( a, b, m1, c, d, m2, mo ) bind(C, name="multiply")
        use iso_c_binding
        implicit none

        integer(c_int), intent(in), value :: a, b, c, d
        real(c_double), dimension(a*b), intent(in) :: m1
        real(c_double), dimension(c*d), intent(in) :: m2
        real(c_double), dimension(a, d), intent(out) :: mo
        integer :: i, j, k

        ! Print the address of the matrices
        print *, "Address of Matrix A: ", loc(m1)
        print *, "Address of Matrix B: ", loc(m2)
        print *, "Address of Matrix AxB: ", loc(mo)

        ! Perform the matrix multiplication
        mo = matmul(reshape(m1, (/a, b/)), reshape(m2, (/c, d/)))

    end subroutine multiply

end module MatrixModule