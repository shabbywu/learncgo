cmake_minimum_required (VERSION 3.8)
include(${CMAKE_CURRENT_SOURCE_DIR}/../../../cmake/golang.cmake)

get_fullname(NAME)
add_go_executable(
    ${NAME}
    OUTPUT ${CMAKE_SOURCE_DIR}/bin/${NAME}
    PACKAGES main.go
)