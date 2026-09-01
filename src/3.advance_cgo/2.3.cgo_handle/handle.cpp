#include "handle.h"

#include <new>

extern "C" void goHandleCallback(uintptr_t value);

namespace {
struct HandleHolder {
  uintptr_t value;
};
}

extern "C" void *handle_new(uintptr_t value) {
  return new (std::nothrow) HandleHolder{value};
}

extern "C" void handle_invoke(void *handle) {
  goHandleCallback(static_cast<HandleHolder *>(handle)->value);
}

extern "C" void handle_free(void *handle) {
  delete static_cast<HandleHolder *>(handle);
}
