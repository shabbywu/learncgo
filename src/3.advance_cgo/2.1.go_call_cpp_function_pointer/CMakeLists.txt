include(${GOLANG_CMAKE_LIB})

get_fullname(NAME)
add_library(callback STATIC callback.cpp)
get_target_property(LIB_BINARY_DIR callback BINARY_DIR)
add_go_executable(
  ${NAME}
  OUTPUT ${CMAKE_SOURCE_DIR}/bin/${NAME}
  PACKAGES main.go
  CGO_LDFLAGS "-L${LIB_BINARY_DIR} -lcallback -lstdc++"
  DEPENDS callback)
set_tests_properties(${NAME} PROPERTIES PASS_REGULAR_EXPRESSION "OK ${NAME}")
