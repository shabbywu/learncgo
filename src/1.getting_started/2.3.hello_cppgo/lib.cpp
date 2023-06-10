#include <iostream>
#include <string>


extern "C" {
    void* CppString (char* s, int size) {
        std::string* string = new std::string();
        string->assign(s, size);
        return string;
    }

    void CppPrint(std::string s) {
        std::cout << s;
    }

    void CppStringFree(void * p) {
        delete (std::string *)p;
    }
}
