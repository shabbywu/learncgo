#pragma once
typedef struct _Foo {
    int foo;
    float bar;
    double baz;
    // case for go keyword
    // func will be rewrote to _func
    int(*func)();
    // _func will be rewrote to __func, for _func is defined
    int(*_func)();

    // field name mapping is independent of definition order
    int _type;
    int __type;
    int type;
} Foo;


#ifdef __cplusplus
extern "C" {
#endif

Foo Foo_init (int foo, float bar, double baz);
Foo* Foo_new(int foo, float bar, double baz);
void Foo_free(Foo* p);

#ifdef __cplusplus
}
#endif
