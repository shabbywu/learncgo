set(GOLANG_CMAKE_LIB ${CMAKE_CURRENT_SOURCE_DIR}/cmake/golang.cmake)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/bin)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/bin)
include(${GOLANG_CMAKE_LIB})
cmake_minimum_required(VERSION 3.8)

# Enable VCPKG if VCPKG_ROOT defined.
if(DEFINED ENV{VCPKG_ROOT})
  set(CMAKE_TOOLCHAIN_FILE
      "$ENV{VCPKG_ROOT}/scripts/buildsystems/vcpkg.cmake"
      CACHE FILEPATH "")
endif()

option(INSTALL_PYTHON "Install Python with vcpkg?" OFF)
if(INSTALL_PYTHON)
  list(APPEND VCPKG_MANIFEST_FEATURES "with-python3")
else()
  list(APPEND VCPKG_MANIFEST_FEATURES "without-python3")
endif()

project("learncgo")
set(CMAKE_CXX_STANDARD 20)

set(CHAPTERS 1.getting_started 2.cpplib_exercise 5.final_practice)

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
  if(DEFINED ${CHAPTER})
    foreach(DEMO ${${CHAPTER}})
      add_subdirectory(src/${CHAPTER}/${DEMO})
    endforeach(DEMO)
  else()
    add_subdirectory(src/${CHAPTER})
  endif()
endforeach(DEMO)
