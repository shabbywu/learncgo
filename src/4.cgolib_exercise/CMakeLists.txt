include(${GOLANG_CMAKE_LIB})

get_fullname(NAME)
set(GOLIB_ARCHIVE ${CMAKE_CURRENT_BINARY_DIR}/libtrust_go.a)

add_go_static_library(${NAME}_golib OUTPUT ${GOLIB_ARCHIVE} PACKAGES
                      go/libtrust.go)

add_executable(${NAME} cpp/example.cpp)
add_dependencies(${NAME} ${NAME}_golib)
target_include_directories(${NAME} PRIVATE ${CMAKE_CURRENT_BINARY_DIR}
                                           ${CMAKE_CURRENT_SOURCE_DIR}/cpp)
target_link_libraries(${NAME} PRIVATE ${GOLIB_ARCHIVE})

if(BUILD_TESTING)
  add_test(NAME ${NAME} COMMAND $<TARGET_FILE:${NAME}>)
  set_tests_properties(${NAME} PROPERTIES ENVIRONMENT "GODEBUG=cgocheck=1"
                                          PASS_REGULAR_EXPRESSION "OK ${NAME}")
endif()
