# (WIP) learncgo
Hand-crafted cgo examples


# RoadMap
- [x] 1.getting_started - basic knowledge of c/c++, such as how to link c/c++ with go, how to access struct/union/enum/class/pointer via go
  - [x] 1.1.hello_go - hello world example for go
  - [x] 2.1.hello_cgo_embeded - cgo example, where c code is embeded at go source code as comment
  - [x] 2.2.hello_cgo_lib - cgo example, where c code is compile as c static lib
  - [x] 2.3.hello_cppgo - cgo example, in this case we show that how to use go to access c++ **std::stirng**
  - [x] 3.1.hello_cstruct - cgo example, in this case we show that how to use go to access c++ **struct**
  - [x] 3.2.hello_cunion - cgo example, in this case we show that how to use go to access c++ **union** via cpp function
  - [x] 3.3.hello_cunion_ultra - cgo example, in this case we show that how to use go to access c++ **union** directly
  - [x] 3.4.hello_cenum - cgo example, in this case we show that how to use go to access c++ **enum**
  - [x] 3.5.hello_cpointer - cgo example, in this case we show that how to use go to access c++ **pointer**
  - [x] 3.6.hello_cppclass - cgo example, in this case we show that how to use go to access c++ **class**
- [x] 2.cpplib_exercise - a full example shows how to use go to access c++ libraries - [toml11](https://github.com/ToruNiina/toml11)
  - [x] 1.1.hello_toml11 - cpp example using toml11, in this case we show that how to export toml11 api from a cpp static library, and use it from another cpp executable
  - [x] 1.2.c_toml11 - c example using toml11, in this case we show that how to export toml11 api from a cpp static library, and use it from another c executable
  - [x] 2.cgo_toml11 - cgo example using toml11, in this case we show that how to export toml11 api from a cpp static library, and use it from another go executable
  - [x] 3.go_toml11 - go example using toml11, in this case we show that how to wrap toml11 cpp static library as a go library, and use it from another go file
- [ ] 3.advance_cgo - basic knowledge of cgo advance usage - such as how C references to Go, how to use c/c++ function pointer via go(callback)
  - [ ] 1.1.cpp_std_array
  - [ ] 1.2.cpp_std_vector
  - [ ] 1.3.cpp_std_set
  - [ ] 1.4.cpp_std_map
  - [ ] 2.1.go_call_cpp_function_pointer - cgo example, in this case we show that how to call c/c++ function pointer from go
  - [ ] 2.2.cpp_call_go_export_callback - cgo example, in this case we show that how to call go function from cpp
  - [ ] 2.3 cgo_handle - cgo example, in this case we show that how to use cpp maintain go pointers lifecycle
- [ ] 4.cgolib_exercise - a full example shows how to use c++ access go libraries - [docker/libtrust](https://github.com/distribution/distribution/tree/release/2.8/vendor/github.com/docker/libtrust)
- [ ] 5.final_practice - a full example shows how to use cpython to access go libraries through c++
