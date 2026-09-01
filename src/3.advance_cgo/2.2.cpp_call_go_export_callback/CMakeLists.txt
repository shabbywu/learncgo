include(${GOLANG_CMAKE_LIB})

get_fullname(NAME)
add_go_static_library(${NAME}_golib OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/golib.a
                      PACKAGES go_callback.go)
add_executable(${NAME} callback.cpp)
add_dependencies(${NAME} ${NAME}_golib)
target_include_directories(${NAME} PRIVATE ${CMAKE_CURRENT_BINARY_DIR})
target_link_libraries(${NAME} PRIVATE ${CMAKE_CURRENT_BINARY_DIR}/golib.a)
if(BUILD_TESTING)
  add_test(NAME ${NAME} COMMAND $<TARGET_FILE:${NAME}>)
  set_tests_properties(${NAME} PROPERTIES PASS_REGULAR_EXPRESSION "OK ${NAME}")
endif()
