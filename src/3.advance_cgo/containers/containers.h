#ifndef LEARNCGO_CONTAINERS_H
#define LEARNCGO_CONTAINERS_H

#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

void *array_new(void);
void array_free(void *handle);
size_t array_len(void *handle);
int array_get(void *handle, size_t index);
void array_set(void *handle, size_t index, int value);

void *vector_new(void);
void vector_free(void *handle);
size_t vector_len(void *handle);
int vector_get(void *handle, size_t index);
void vector_push(void *handle, int value);

void *set_new(void);
void set_free(void *handle);
size_t set_len(void *handle);
int set_get(void *handle, size_t index);
int set_insert(void *handle, int value);
int set_contains(void *handle, int value);

void *map_new(void);
void map_free(void *handle);
size_t map_len(void *handle);
void map_set(void *handle, const char *key, size_t key_len, int value);
int map_get(void *handle, const char *key, size_t key_len, int *value);

#ifdef __cplusplus
}
#endif

#endif
