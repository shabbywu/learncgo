include(${GOLANG_CMAKE_LIB})

if(NOT DEFINED Python_EXECUTABLE)
  find_package(
    Python3
    COMPONENTS Development
    REQUIRED)
endif()
find_package(pybind11 CONFIG)

get_fullname(NAME)
slugify(${NAME} SLUG_NAME)
add_go_static_library(
  ${NAME}_golib OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/${NAME}_golib.a PACKAGES
  go/libBash.go)

add_executable(${NAME}_cppexample cpp/example.cpp)
target_link_libraries(${NAME}_cppexample
                      PUBLIC ${CMAKE_CURRENT_BINARY_DIR}/${NAME}_golib.a)

pybind11_add_module(_${NAME}_golib cpp/pybind11.cpp)
target_link_libraries(
  _${NAME}_golib
  PUBLIC ${CMAKE_CURRENT_BINARY_DIR}/${NAME}_golib.a pybind11::pybind11
         $<IF:$<TARGET_EXISTS:Python3::Python>,Python3::Python,>)

set_target_properties(_${NAME}_golib PROPERTIES OUTPUT_NAME "libBash")
