# CMakeList.txt: 顶层 CMake 项目文件，在此处执行全局配置
# 并包含子项目。
#
include(${CMAKE_CURRENT_SOURCE_DIR}/cmake/golang.cmake)
cmake_minimum_required (VERSION 3.8)

# Enable Hot Reload for MSVC compilers if supported.
if (POLICY CMP0141)
  cmake_policy(SET CMP0141 NEW)
  set(CMAKE_MSVC_DEBUG_INFORMATION_FORMAT "$<IF:$<AND:$<C_COMPILER_ID:MSVC>,$<CXX_COMPILER_ID:MSVC>>,$<$<CONFIG:Debug,RelWithDebInfo>:EditAndContinue>,$<$<CONFIG:Debug,RelWithDebInfo>:ProgramDatabase>>")
endif()

project ("learncgo")


set(CHAPTERS
    1.getting_started
)

set(1.getting_started
  1.1.hello_go
  2.1.hello_cgo_embed
  2.2.hello_cgo_lib
  2.3.hello_cppgo
  3.1.hello_cstruct
  3.2.hello_cunion
  3.3.hello_cunion_ultra
  3.4.hello_cenum
  3.5.hello_cppclass
)


# then create a project file per tutorial
foreach(CHAPTER ${CHAPTERS})
  foreach(DEMO ${${CHAPTER}})
    add_subdirectory (src/${CHAPTER}/${DEMO})
  endforeach(DEMO)
endforeach(DEMO)
