#include <iostream>
#include "libBash.h"

int main() {
    char* bashCommand = BashCommandWithTokens(4);
    std::cout << "BashCommandWithTokens(4): " << bashCommand << std::endl;
}
