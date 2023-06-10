typedef struct _Foo {
    union
    {
        int foo;
        float bar;
        double baz;
    } u;
} Foo;

typedef union _u {
    int foo;
    float bar;
    double baz;
} U;


#ifdef __cplusplus
extern "C" {
#endif

Foo Foo_init_foo (int foo);
Foo Foo_init_bar (float bar);
Foo Foo_init_baz (double baz);

// golang can not access a union value directly, so we must defind a reader to access that union fields.
int read_u_foo(U o);
float read_u_bar(U o);
double read_u_baz(U o);

#ifdef __cplusplus
}
#endif