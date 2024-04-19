function multiply( a, b ) result( c ) bind(C, name="multiply")
        use iso_c_binding
        implicit none

        integer(c_int), intent(in), value :: a, b
        integer(c_int) :: c

        c = a * b
end function 

subroutine arrayByHundred( arr, n ) bind(C, name="arrayByHundred")
        use iso_c_binding
        implicit none

        integer(c_int), intent(inout) :: n
        real(c_double), dimension(n), intent(inout) :: arr

        arr(:) = arr(:) * 100.0d0
end subroutine
