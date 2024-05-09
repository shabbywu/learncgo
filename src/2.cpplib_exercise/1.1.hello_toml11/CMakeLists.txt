cmake_minimum_required(VERSION 3.8)

get_fullname(NAME)

add_library(${NAME}_toml11_lib STATIC toml11_lib.cpp)
target_include_directories(${NAME}_toml11_lib PUBLIC ${LIBS_DIR}/toml11)

add_executable(${NAME} example.cpp)
target_link_libraries(${NAME} PUBLIC ${NAME}_toml11_lib)