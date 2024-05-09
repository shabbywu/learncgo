#include <string>
#include <iostream>
#include <memory>
#include "toml11_lib.h"


int main() {
    std::string content = R"(
title = "an example toml file"
nums  = [3, 1, 4, 1, 5]
    )";

    std::shared_ptr<void> h = toml11_parse_content(content);
    std::cout << "title: " << toml11_find_string(h, "title") << std::endl;;
}
