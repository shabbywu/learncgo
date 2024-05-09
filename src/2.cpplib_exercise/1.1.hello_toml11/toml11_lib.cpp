#include <toml.hpp>
#include <strstream>
#include <string>
#include <memory>
#include "toml11_lib.h"


std::shared_ptr<void> toml11_parse_content(std::string content) {
    std::istrstream is(content.c_str());
    auto ret = new toml::value();
    *ret=toml::parse(is);
    return std::shared_ptr<toml::value>(ret);
}


template<typename V> 
V toml11_find(std::shared_ptr<void> handler, std::string key) {
    std::shared_ptr<toml::value> h = std::static_pointer_cast<toml::value>(handler);
    V v = toml::find<V>(*h, key);
    return v;
}


std::string toml11_find_string(std::shared_ptr<void> handler, std::string key) {
    return toml11_find<std::string>(handler, key);
}