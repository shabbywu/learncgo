#pragma once
#include <memory>
#include <string>

std::shared_ptr<void> toml11_parse_content(std::string content);

std::string toml11_find_string(std::shared_ptr<void> handler, std::string key);
