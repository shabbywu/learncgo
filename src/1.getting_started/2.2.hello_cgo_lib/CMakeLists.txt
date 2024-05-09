cmake_minimum_required(VERSION 3.8)
include(${CMAKE_CURRENT_SOURCE_DIR}/../../../cmake/golang.cmake)

get_fullname(NAME)
set(LIBNAME _${NAME})
add_library(${LIBNAME} STATIC lib.c)
# 获取 LIB 的构建目录
get_target_property(LIB_BINARY_DIR ${LIBNAME} BINARY_DIR)

add_go_executable(
  ${NAME}
  OUTPUT
  ${CMAKE_SOURCE_DIR}/bin/${NAME}
  PACKAGES
  main.go
  CGO_LDFLAGS
  "-l${LIBNAME} -L${LIB_BINARY_DIR}"
  DEPENDS
  ${LIBNAME})
