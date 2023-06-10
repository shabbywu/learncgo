typedef struct _Foo {
    union
    {
        int foo;
        float bar;
        double baz;
    } u;
} Foo;