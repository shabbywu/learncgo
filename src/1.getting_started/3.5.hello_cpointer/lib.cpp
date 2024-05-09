#include "lib.h"

extern "C" {
    Foo Foo_init (int foo, float bar, double baz) {
        Foo obj = {foo, bar, baz};
        return obj;
    }

    Foo* Foo_new(int foo, float bar, double baz) {
        Foo* obj = new Foo{foo, bar, baz};
        return obj;
    }

    void Foo_free(Foo* p) {
        delete (p);
    }
}
