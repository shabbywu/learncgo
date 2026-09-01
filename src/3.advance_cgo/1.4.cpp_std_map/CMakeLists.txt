include(${GOLANG_CMAKE_LIB})

get_fullname(NAME)
add_go_executable(${NAME} OUTPUT ${CMAKE_SOURCE_DIR}/bin/${NAME} PACKAGES main.go)
set_tests_properties(${NAME} PROPERTIES PASS_REGULAR_EXPRESSION "OK ${NAME}")
