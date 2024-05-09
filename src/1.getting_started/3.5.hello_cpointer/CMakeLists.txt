cmake_minimum_required(VERSION 3.8)
include(${CMAKE_CURRENT_SOURCE_DIR}/../../../cmake/golang.cmake)
set(CMAKE_CXX_STANDARD 17) # this does nothing for MSVC, use
                           # target_compile_options below
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS ON)

get_fullname(NAME)
set(LIBNAME _${NAME})
add_library(${LIBNAME} STATIC lib.cpp)
# 获取 LIB 的构建目录
get_target_property(LIB_BINARY_DIR ${LIBNAME} BINARY_DIR)

add_go_executable(
  ${NAME}
  OUTPUT
  ${CMAKE_SOURCE_DIR}/bin/${NAME}
  PACKAGES
  main.go
  CGO_LDFLAGS
  "-l${LIBNAME} -lstdc++ -L${LIB_BINARY_DIR}"
  DEPENDS
  ${LIBNAME})
