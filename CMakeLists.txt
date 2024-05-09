set(GOLANG_CMAKE_LIB ${CMAKE_CURRENT_SOURCE_DIR}/cmake/golang.cmake)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/bin)
include(${GOLANG_CMAKE_LIB})
cmake_minimum_required(VERSION 3.8)

# Enable Hot Reload for MSVC compilers if supported.
if(POLICY CMP0141)
  cmake_policy(SET CMP0141 NEW)
  set(CMAKE_MSVC_DEBUG_INFORMATION_FORMAT
      "$<IF:$<AND:$<C_COMPILER_ID:MSVC>,$<CXX_COMPILER_ID:MSVC>>,$<$<CONFIG:Debug,RelWithDebInfo>:EditAndContinue>,$<$<CONFIG:Debug,RelWithDebInfo>:ProgramDatabase>>"
  )
endif()

project("learncgo")
set(CMAKE_CXX_STANDARD 20)
set(LIBS_DIR ${CMAKE_CURRENT_SOURCE_DIR}/libs)

set(CHAPTERS 1.getting_started 2.cpplib_exercise)

set(1.getting_started
    1.1.hello_go
    2.1.hello_cgo_embeded
    2.2.hello_cgo_lib
    2.3.hello_cppgo
    3.1.hello_cstruct
    3.2.hello_cunion
    3.3.hello_cunion_ultra
    3.4.hello_cenum
    3.5.hello_cpointer
    3.6.hello_cppclass)

set(2.cpplib_exercise 1.1.hello_toml11 1.2.c_toml11 2.cgo_toml11 3.go_toml11)

# then create a project file per tutorial
foreach(CHAPTER ${CHAPTERS})
  foreach(DEMO ${${CHAPTER}})
    add_subdirectory(src/${CHAPTER}/${DEMO})
  endforeach(DEMO)
endforeach(DEMO)
