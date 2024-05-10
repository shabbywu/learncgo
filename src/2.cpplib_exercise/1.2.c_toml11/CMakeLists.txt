cmake_minimum_required(VERSION 3.8)
find_package(toml11 CONFIG REQUIRED)

get_fullname(NAME)

add_library(${NAME}_toml11_lib STATIC toml11_lib.cpp)
target_link_libraries(${NAME}_toml11_lib PRIVATE toml11::toml11)

add_executable(${NAME} example.c)
target_link_libraries(${NAME} PUBLIC ${NAME}_toml11_lib)
