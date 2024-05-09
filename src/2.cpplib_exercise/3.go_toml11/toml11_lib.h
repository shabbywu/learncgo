#pragma once
#ifdef __cplusplus
extern "C" {
#endif
void* toml11_parse_content(char* content);
void toml11_free(void* handler);
const char* toml11_find_string(void* handler, char* key);
#ifdef __cplusplus
}
#endif
