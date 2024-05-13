module MatrixModule
    use iso_c_binding
    implicit none

    type, bind(C) :: Matrix
        integer(c_int) :: rows
        integer(c_int) :: cols
        type(c_ptr) :: data
    end type

contains

    function multiply( a, b ) result( c ) bind(C, name="multiply")        
        use iso_c_binding
        implicit none

        type(Matrix), intent(in) :: a, b
        type(Matrix) :: c
        integer(c_int) :: i, j, k
        real(c_double), dimension(:,:), pointer :: a_data, b_data, c_data
        real(c_double) :: sum

        c%rows = a%rows
        c%cols = b%cols
        allocate(c_data(c%rows, c%cols))

        call c_f_pointer(a%data, a_data, [a%rows, a%cols])
        call c_f_pointer(b%data, b_data, [b%rows, b%cols])
        call c_f_pointer(c%data, c_data, [c%rows, c%cols])

        do i = 1, c%rows
            do j = 1, c%cols
                sum = 0.0_c_double
                do k = 1, a%cols
                    sum = sum + a_data(i,k) * b_data(k,j)
                end do
                c_data(i,j) = sum
            end do
        end do
    end function

end module MatrixModule