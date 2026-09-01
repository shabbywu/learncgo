include(${GOLANG_CMAKE_LIB})

get_fullname(NAME)
add_library(handle STATIC handle.cpp)
get_target_property(LIB_BINARY_DIR handle BINARY_DIR)
add_go_executable(
  ${NAME}
  OUTPUT ${CMAKE_SOURCE_DIR}/bin/${NAME}
  PACKAGES main.go
  CGO_LDFLAGS "-L${LIB_BINARY_DIR} -lhandle -lstdc++"
  DEPENDS handle)
set_tests_properties(
  ${NAME}
  PROPERTIES ENVIRONMENT "GODEBUG=cgocheck=1"
             PASS_REGULAR_EXPRESSION "OK ${NAME}")
