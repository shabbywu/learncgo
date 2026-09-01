#ifndef LEARNCGO_CPP_FUNCTION_POINTER_H
#define LEARNCGO_CPP_FUNCTION_POINTER_H

#ifdef __cplusplus
extern "C" {
#endif

typedef int (*int_callback)(int);

int cpp_increment(int value);
int invoke_int_callback(int_callback callback, int value);

#ifdef __cplusplus
}
#endif

#endif
