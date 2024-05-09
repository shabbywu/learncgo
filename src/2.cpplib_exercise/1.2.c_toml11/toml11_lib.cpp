#include <toml.hpp>
#include <strstream>
#include <string>
#include <memory>
#include "toml11_lib.h"


template<typename V>
V& toml11_find(void* handler, std::string key) {
    toml::value* h =(toml::value*)(handler);
    V& v = toml::find<V>(*h, key);
    return v;
}


void* toml11_parse_content(char* content) {
    std::istrstream is(content);
    auto ret = new toml::value();
    *ret=toml::parse(is);
    return ret;
}

void toml11_free(void* handler) {
    delete (toml::value*)handler;
}

const char* toml11_find_string(void* handler, char* key) {
    auto& v = toml11_find<std::string>(handler, key);
    return v.c_str();
}
