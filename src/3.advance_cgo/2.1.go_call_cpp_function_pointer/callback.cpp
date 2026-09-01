#include "callback.h"

extern "C" int cpp_increment(int value) { return value + 1; }

extern "C" int invoke_int_callback(int_callback callback, int value) {
  return callback(value);
}
