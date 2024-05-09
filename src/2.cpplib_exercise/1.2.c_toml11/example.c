#include <stdio.h>
#include "toml11_lib.h"


int main() {
    char* content = "title = \"an example toml file\"\nnums  = [3, 1, 4, 1, 5]";

    void* h = toml11_parse_content(content);
    const char* title = toml11_find_string(h, "title");
    printf("tilte: %s\n", title);

    toml11_free(h);
    printf("tilte: %s\n", title);
}
