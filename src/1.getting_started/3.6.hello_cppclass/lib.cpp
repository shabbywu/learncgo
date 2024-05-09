#include "lib.h"
#ifdef __cplusplus

MyClass::MyClass(int value) : m_value(value) {}

void MyClass::SetValue(int value) {
    m_value = value;
}

int MyClass::GetValue() const {
    return m_value;
}

extern "C" {
#endif

MyClass* MyClass_new(int value) {
    MyClass* obj = new MyClass(value);
    return obj;
}

void MyClass_SetValue(MyClass* obj, int value) {
    obj->SetValue(value);
}

int MyClass_GetValue(MyClass* obj) {
    return obj->GetValue();
}

void MyClass_delete(MyClass* obj) {
    delete obj;
}

#ifdef __cplusplus
}
#endif
