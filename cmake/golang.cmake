# add_go_executable: add target that build as go executable
function(add_go_executable)
  set(options OPTIONAL FAST)
  set(oneValueArgs
      OUTPUT
      CGO_CFLAGS
      CGO_CPPFLAGS
      CGO_CXXFLAGS
      CGO_FFLAGS
      CGO_LDFLAGS
      PKG_CONFIG)
  set(multiValueArgs PACKAGES DEPENDS)

  cmake_parse_arguments(GOBUILD "${options}" "${oneValueArgs}"
                        "${multiValueArgs}" ${ARGN})

  set(GOBUILD_ENV)
  set(GOBUILD_ENV_KEYS)
  # 处理 GO build 的环境变量
  set(CGO_ENV GOBUILD_CGO_CFLAGS GOBUILD_CGO_CPPFLAGS GOBUILD_CGO_CXXFLAGS
              GOBUILD_CGO_FFLAGS GOBUILD_CGO_LDFLAGS)
  list(APPEND GOBUILD_ENV_KEYS ${CGO_ENV})

  # then create a project file per tutorial
  foreach(ENV_KEY ${GOBUILD_ENV_KEYS})
    if(DEFINED ${ENV_KEY})
      string(REPLACE "GOBUILD_" "" REAL_ENV_KEY ${ENV_KEY})
      list(APPEND GOBUILD_ENV ${REAL_ENV_KEY}=${${ENV_KEY}})
    endif()
  endforeach(ENV_KEY)

  if(NOT DEFINED GOBUILD_OUTPUT)
    set(GOBUILD_OUTPUT ${NAME})
  endif()

  if(NOT DEFINED GOBUILD_PACKAGES)
    message(FATAL_ERROR "PACKAGES is required!")
  endif()

  message("${GOBUILD_ENV} go build -o ${GOBUILD_OUTPUT} ${GOBUILD_PACKAGES}")
  add_custom_target(${NAME} ALL)
  add_custom_command(
    TARGET ${NAME}
    COMMAND env ${GOBUILD_ENV} go build -o ${GOBUILD_OUTPUT} ${GOBUILD_PACKAGES}
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    DEPENDS ${DEPENDS}
    COMMENT "Building ${NAME}")
endfunction()

# add_go_static_library: add target that build as go c-archive
function(add_go_static_library)
  set(options OPTIONAL FAST)
  set(oneValueArgs
      OUTPUT
      CGO_CFLAGS
      CGO_CPPFLAGS
      CGO_CXXFLAGS
      CGO_FFLAGS
      CGO_LDFLAGS
      PKG_CONFIG)
  set(multiValueArgs PACKAGES DEPENDS)

  cmake_parse_arguments(GOBUILD "${options}" "${oneValueArgs}"
                        "${multiValueArgs}" ${ARGN})

  set(GOBUILD_ENV)
  set(GOBUILD_ENV_KEYS)
  # 处理 GO build 的环境变量
  set(CGO_ENV GOBUILD_CGO_CFLAGS GOBUILD_CGO_CPPFLAGS GOBUILD_CGO_CXXFLAGS
              GOBUILD_CGO_FFLAGS GOBUILD_CGO_LDFLAGS)
  list(APPEND GOBUILD_ENV_KEYS ${CGO_ENV})

  # then create a project file per tutorial
  foreach(ENV_KEY ${GOBUILD_ENV_KEYS})
    if(DEFINED ${ENV_KEY})
      string(REPLACE "GOBUILD_" "" REAL_ENV_KEY ${ENV_KEY})
      list(APPEND GOBUILD_ENV ${REAL_ENV_KEY}=${${ENV_KEY}})
    endif()
  endforeach(ENV_KEY)

  if(NOT DEFINED GOBUILD_OUTPUT)
    set(GOBUILD_OUTPUT ${NAME})
  endif()

  if(NOT DEFINED GOBUILD_PACKAGES)
    message(FATAL_ERROR "PACKAGES is required!")
  endif()

  message("${GOBUILD_ENV} go build -o ${GOBUILD_OUTPUT} ${GOBUILD_PACKAGES}")
  add_custom_target(${NAME} ALL)
  add_custom_command(
    TARGET ${NAME}
    COMMAND env ${GOBUILD_ENV} go build -o ${GOBUILD_OUTPUT}
            -buildmode=c-archive ${GOBUILD_PACKAGES}
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    DEPENDS ${DEPENDS}
    COMMENT "Building ${NAME}")
endfunction()

function(get_fullname result)
  if(CMAKE_SOURCE_DIR STREQUAL CMAKE_CURRENT_SOURCE_DIR)
    get_filename_component(parent_dir "${CMAKE_SOURCE_DIR}" DIRECTORY)
    string(REPLACE ${parent_dir}/ "" ${result} ${CMAKE_CURRENT_SOURCE_DIR})
  else()
    message("CMAKE_SOURCE_DIR: ${CMAKE_SOURCE_DIR}")
    message("CMAKE_CURRENT_SOURCE_DIR: ${CMAKE_CURRENT_SOURCE_DIR}")
    string(REPLACE ${CMAKE_SOURCE_DIR}/src/ "" ${result}
                   ${CMAKE_CURRENT_SOURCE_DIR})
  endif()
  string(REPLACE "/" "_" ${result} ${${result}})
  set(${result}
      ${${result}}
      PARENT_SCOPE)
endfunction()

function(slugify input_string output_variable)
  string(TOLOWER "${input_string}" slug)
  string(REGEX REPLACE "[^a-z0-9]+" "_" slug "${slug}")
  string(REGEX REPLACE "^-+|-+$" "" slug "${slug}")
  set(${output_variable}
      "${slug}"
      PARENT_SCOPE)
endfunction()
