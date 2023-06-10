#ifdef __cplusplus

class MyClass {
public:
    MyClass(int value);
    void SetValue(int value);
    int GetValue() const;
private:
    int m_value;
};

extern "C" {
#endif

typedef struct MyClass MyClass;

MyClass* MyClass_new(int value);
void MyClass_SetValue(MyClass* obj, int value);
int MyClass_GetValue(MyClass* obj);
void MyClass_delete(MyClass* obj);

#ifdef __cplusplus
}
#endif