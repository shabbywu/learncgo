// Package containers exposes selected C++ standard containers through a small,
// opaque C ABI. Values returned to callers are copied into Go memory.
package containers

// #cgo CXXFLAGS: -std=c++17
// #include "containers.h"
import "C"

import (
	"errors"
	"sync"
	"unsafe"
)

type native struct {
	mu     sync.RWMutex
	handle unsafe.Pointer
	free   func(unsafe.Pointer)
}

func newNative(handle unsafe.Pointer, free func(unsafe.Pointer)) (native, error) {
	if handle == nil {
		return native{}, errors.New("containers: native allocation failed")
	}
	return native{handle: handle, free: free}, nil
}

func (n *native) withHandle(fn func(unsafe.Pointer)) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	if n.handle == nil {
		panic("containers: use after Close")
	}
	fn(n.handle)
}

func (n *native) close() {
	n.mu.Lock()
	defer n.mu.Unlock()
	if n.handle != nil {
		n.free(n.handle)
		n.handle = nil
	}
}

// Array is a fixed-length (four element) C++ std::array<int, 4>.
type Array struct{ native }

func NewArray() (*Array, error) {
	native, err := newNative(C.array_new(), func(h unsafe.Pointer) { C.array_free(h) })
	if err != nil {
		return nil, err
	}
	return &Array{native}, nil
}

func (a *Array) Close() { a.close() }
func (a *Array) Len() int {
	var result int
	a.withHandle(func(h unsafe.Pointer) { result = int(C.array_len(h)) })
	return result
}
func (a *Array) Get(index int) int {
	if index < 0 || index >= 4 {
		panic("containers: array index out of range")
	}
	var result int
	a.withHandle(func(h unsafe.Pointer) { result = int(C.array_get(h, C.size_t(index))) })
	return result
}
func (a *Array) Set(index, value int) {
	if index < 0 || index >= 4 {
		panic("containers: array index out of range")
	}
	a.withHandle(func(h unsafe.Pointer) { C.array_set(h, C.size_t(index), C.int(value)) })
}
func (a *Array) Values() []int {
	values := make([]int, a.Len())
	for i := range values {
		values[i] = a.Get(i)
	}
	return values
}

// Vector is a C++ std::vector<int>.
type Vector struct{ native }

func NewVector() (*Vector, error) {
	native, err := newNative(C.vector_new(), func(h unsafe.Pointer) { C.vector_free(h) })
	if err != nil {
		return nil, err
	}
	return &Vector{native}, nil
}
func (v *Vector) Close() { v.close() }
func (v *Vector) Len() int {
	var result int
	v.withHandle(func(h unsafe.Pointer) { result = int(C.vector_len(h)) })
	return result
}
func (v *Vector) Push(value int) {
	v.withHandle(func(h unsafe.Pointer) { C.vector_push(h, C.int(value)) })
}
func (v *Vector) Get(index int) int {
	if index < 0 || index >= v.Len() {
		panic("containers: vector index out of range")
	}
	var result int
	v.withHandle(func(h unsafe.Pointer) { result = int(C.vector_get(h, C.size_t(index))) })
	return result
}
func (v *Vector) Values() []int {
	values := make([]int, v.Len())
	for i := range values {
		values[i] = v.Get(i)
	}
	return values
}

// Set is a C++ std::set<int>. Values returns ascending values.
type Set struct{ native }

func NewSet() (*Set, error) {
	native, err := newNative(C.set_new(), func(h unsafe.Pointer) { C.set_free(h) })
	if err != nil {
		return nil, err
	}
	return &Set{native}, nil
}
func (s *Set) Close() { s.close() }
func (s *Set) Len() int {
	var result int
	s.withHandle(func(h unsafe.Pointer) { result = int(C.set_len(h)) })
	return result
}
func (s *Set) Insert(value int) bool {
	var inserted bool
	s.withHandle(func(h unsafe.Pointer) { inserted = C.set_insert(h, C.int(value)) != 0 })
	return inserted
}
func (s *Set) Contains(value int) bool {
	var found bool
	s.withHandle(func(h unsafe.Pointer) { found = C.set_contains(h, C.int(value)) != 0 })
	return found
}
func (s *Set) Values() []int {
	values := make([]int, s.Len())
	s.withHandle(func(h unsafe.Pointer) {
		for i := range values {
			values[i] = int(C.set_get(h, C.size_t(i)))
		}
	})
	return values
}

// Map is a C++ std::map<string, int>. Keys are passed as pointer-and-length,
// so embedded NUL bytes are preserved and no Go pointer is retained by C++.
type Map struct{ native }

func NewMap() (*Map, error) {
	native, err := newNative(C.map_new(), func(h unsafe.Pointer) { C.map_free(h) })
	if err != nil {
		return nil, err
	}
	return &Map{native}, nil
}
func (m *Map) Close() { m.close() }
func (m *Map) Len() int {
	var result int
	m.withHandle(func(h unsafe.Pointer) { result = int(C.map_len(h)) })
	return result
}
func (m *Map) Set(key string, value int) {
	bytes := []byte(key)
	m.withHandle(func(h unsafe.Pointer) {
		var pointer *C.char
		if len(bytes) != 0 {
			pointer = (*C.char)(unsafe.Pointer(&bytes[0]))
		}
		C.map_set(h, pointer, C.size_t(len(bytes)), C.int(value))
	})
}
func (m *Map) Get(key string) (int, bool) {
	bytes := []byte(key)
	var value C.int
	var found bool
	m.withHandle(func(h unsafe.Pointer) {
		var pointer *C.char
		if len(bytes) != 0 {
			pointer = (*C.char)(unsafe.Pointer(&bytes[0]))
		}
		found = C.map_get(h, pointer, C.size_t(len(bytes)), &value) != 0
	})
	return int(value), found
}
