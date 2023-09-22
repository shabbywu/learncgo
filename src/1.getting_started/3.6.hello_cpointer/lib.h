typedef struct _Foo {
    int foo;
    float bar;
    double baz;
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