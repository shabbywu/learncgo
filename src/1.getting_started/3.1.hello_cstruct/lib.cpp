#include "lib.h"


int func() {
    return 100;
}


extern "C" {
    Foo Foo_init (int foo, float bar, double baz) {
        Foo obj = {foo, bar, baz, &func, &func, 1, 2, 3};
        return obj;
    }

    Foo* Foo_new(int foo, float bar, double baz) {
        Foo* obj = new Foo{foo, bar, baz, nullptr, nullptr, 1, 2, 3};
        return obj;
    }

    void Foo_free(Foo* p) {
        delete (p);
    }
}
