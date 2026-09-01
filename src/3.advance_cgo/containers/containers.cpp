#include "containers.h"

#include <array>
#include <map>
#include <new>
#include <set>
#include <string>
#include <vector>

namespace {
template <typename T> T *as(void *handle) { return static_cast<T *>(handle); }
}

extern "C" {
void *array_new(void) { return new (std::nothrow) std::array<int, 4>{}; }
void array_free(void *handle) { delete as<std::array<int, 4>>(handle); }
size_t array_len(void *) { return 4; }
int array_get(void *handle, size_t index) {
  return index < 4 ? (*as<std::array<int, 4>>(handle))[index] : 0;
}
void array_set(void *handle, size_t index, int value) {
  if (index < 4) (*as<std::array<int, 4>>(handle))[index] = value;
}

void *vector_new(void) { return new (std::nothrow) std::vector<int>(); }
void vector_free(void *handle) { delete as<std::vector<int>>(handle); }
size_t vector_len(void *handle) { return as<std::vector<int>>(handle)->size(); }
int vector_get(void *handle, size_t index) {
  const auto &values = *as<std::vector<int>>(handle);
  return index < values.size() ? values[index] : 0;
}
void vector_push(void *handle, int value) { as<std::vector<int>>(handle)->push_back(value); }

void *set_new(void) { return new (std::nothrow) std::set<int>(); }
void set_free(void *handle) { delete as<std::set<int>>(handle); }
size_t set_len(void *handle) { return as<std::set<int>>(handle)->size(); }
int set_get(void *handle, size_t index) {
  const auto &values = *as<std::set<int>>(handle);
  if (index >= values.size()) return 0;
  auto iterator = values.begin();
  std::advance(iterator, static_cast<long>(index));
  return *iterator;
}
int set_insert(void *handle, int value) { return as<std::set<int>>(handle)->insert(value).second; }
int set_contains(void *handle, int value) { return as<std::set<int>>(handle)->count(value) != 0; }

void *map_new(void) { return new (std::nothrow) std::map<std::string, int>(); }
void map_free(void *handle) { delete as<std::map<std::string, int>>(handle); }
size_t map_len(void *handle) { return as<std::map<std::string, int>>(handle)->size(); }
void map_set(void *handle, const char *key, size_t key_len, int value) {
  (*as<std::map<std::string, int>>(handle))[std::string(key ? key : "", key_len)] = value;
}
int map_get(void *handle, const char *key, size_t key_len, int *value) {
  const auto &values = *as<std::map<std::string, int>>(handle);
  const auto iterator = values.find(std::string(key ? key : "", key_len));
  if (iterator == values.end()) return 0;
  *value = iterator->second;
  return 1;
}
}
