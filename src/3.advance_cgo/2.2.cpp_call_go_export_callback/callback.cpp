#include "golib.h"

#include <iostream>

int main() {
  if (GoDouble(21) != 42) return 1;
  std::cout << "OK 3.advance_cgo_2.2.cpp_call_go_export_callback" << std::endl;
  return 0;
}
