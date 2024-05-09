#include "lib.h"

extern "C" {
    Foo Foo_init_foo (int foo) {
        Foo o = Foo{.u = {foo}};
        return o;
    }
    Foo Foo_init_bar (float bar) {
        Foo o = Foo{.u = { .bar = bar}};
        return o;
    }
    Foo Foo_init_baz (double baz) {
        Foo o = Foo{.u = { .baz = baz}};
        return o;
    }

    int read_u_foo(U o) {
        return o.foo;
    }

    float read_u_bar(U o) {
        return o.bar;
    }

    double read_u_baz(U o) {
        return o.baz;
    }
}
