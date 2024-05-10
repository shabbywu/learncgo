#include <pybind11/pybind11.h>
#include "libBash.h"
namespace py = pybind11;


namespace _5_final_practice_golib {
    std::string __BashCommandWithTokens(int nToken) {
        auto cgoResult = BashCommandWithTokens(nToken);
        std::string result = cgoResult;
        free(cgoResult);
        return result;
    }
}


PYBIND11_MODULE(libBash, m) {
    m.def("BashCommandWithTokens", &_5_final_practice_golib::__BashCommandWithTokens);
    m.attr("__author__") = "shabbywu<shabbywu@qq.com>";
}
