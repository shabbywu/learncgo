cmake_minimum_required(VERSION 3.8)
include(${CMAKE_CURRENT_SOURCE_DIR}/../../../cmake/golang.cmake)
set(CMAKE_CXX_STANDARD 17) # this does nothing for MSVC, use
                           # target_compile_options below
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS ON)

get_fullname(NAME)
add_library(_${NAME}_toml11_lib STATIC toml11_lib.cpp)
target_include_directories(_${NAME}_toml11_lib PUBLIC ${LIBS_DIR}/toml11)
# 获取 LIB 的构建目录
get_target_property(LIB_BINARY_DIR _${NAME}_toml11_lib BINARY_DIR)

add_go_executable(
  ${NAME}
  OUTPUT
  ${CMAKE_SOURCE_DIR}/bin/${NAME}
  PACKAGES
  main.go
  CGO_LDFLAGS
  "-l_${NAME}_toml11_lib -lstdc++ -L${LIB_BINARY_DIR}"
  DEPENDS
  _${NAME}_toml11_lib)
