#ifndef LEARNCGO_HANDLE_H
#define LEARNCGO_HANDLE_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

void *handle_new(uintptr_t value);
void handle_invoke(void *handle);
void handle_free(void *handle);

#ifdef __cplusplus
}
#endif

#endif
