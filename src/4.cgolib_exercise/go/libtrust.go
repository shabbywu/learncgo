package main

// #include <stdint.h>
// #include <stdlib.h>
import "C"

import (
	"encoding/pem"
	"errors"
	"fmt"
	"runtime/cgo"
	"sync"
	"unsafe"

	"github.com/docker/libtrust"
)

var keyHandles = struct {
	sync.RWMutex
	values map[uintptr]cgo.Handle
}{values: make(map[uintptr]cgo.Handle)}

func setError(destination **C.char, err error) C.int {
	if destination != nil {
		*destination = C.CString(err.Error())
	}
	return 1
}

func clearError(destination **C.char) {
	if destination != nil {
		*destination = nil
	}
}

func setBytes(destination **C.char, length *C.size_t, value []byte) error {
	clearBytes(destination, length)
	if destination == nil || length == nil {
		return errors.New("output buffer and length are required")
	}
	if len(value) == 0 {
		return nil
	}
	*destination = (*C.char)(C.CBytes(value))
	*length = C.size_t(len(value))
	return nil
}

func clearBytes(destination **C.char, length *C.size_t) {
	if destination != nil {
		*destination = nil
	}
	if length != nil {
		*length = 0
	}
}

func copyInput(input *C.char, length C.size_t) ([]byte, error) {
	if length == 0 {
		return []byte{}, nil
	}
	if input == nil {
		return nil, errors.New("input is nil with a non-zero length")
	}
	if uint64(length) > uint64(^uint(0)>>1) {
		return nil, errors.New("input is too large")
	}
	return append([]byte(nil), unsafe.Slice((*byte)(unsafe.Pointer(input)), int(length))...), nil
}

func storeKey(key libtrust.PrivateKey) uintptr {
	handle := cgo.NewHandle(key)
	id := uintptr(handle)
	keyHandles.Lock()
	keyHandles.values[id] = handle
	keyHandles.Unlock()
	return id
}

func loadKey(id C.uintptr_t) (libtrust.PrivateKey, error) {
	keyHandles.RLock()
	handle, found := keyHandles.values[uintptr(id)]
	if !found {
		keyHandles.RUnlock()
		return nil, errors.New("invalid or released key handle")
	}
	key, valid := handle.Value().(libtrust.PrivateKey)
	keyHandles.RUnlock()
	if !valid {
		return nil, errors.New("key handle has an unexpected value")
	}
	return key, nil
}

//export libtrust_generate_ec_p256
func libtrust_generate_ec_p256(handle *C.uintptr_t, errorMessage **C.char) C.int {
	clearError(errorMessage)
	if handle == nil {
		return setError(errorMessage, errors.New("key handle output is required"))
	}
	*handle = 0
	key, err := libtrust.GenerateECP256PrivateKey()
	if err != nil {
		return setError(errorMessage, err)
	}
	*handle = C.uintptr_t(storeKey(key))
	return 0
}

//export libtrust_key_free
func libtrust_key_free(id C.uintptr_t) {
	keyHandles.Lock()
	handle, found := keyHandles.values[uintptr(id)]
	if found {
		delete(keyHandles.values, uintptr(id))
	}
	keyHandles.Unlock()
	if found {
		handle.Delete()
	}
}

//export libtrust_buffer_free
func libtrust_buffer_free(value *C.char) {
	C.free(unsafe.Pointer(value))
}

//export libtrust_key_id
func libtrust_key_id(id C.uintptr_t, output **C.char, outputLength *C.size_t, errorMessage **C.char) C.int {
	clearError(errorMessage)
	clearBytes(output, outputLength)
	key, err := loadKey(id)
	if err != nil {
		return setError(errorMessage, err)
	}
	if err := setBytes(output, outputLength, []byte(key.KeyID())); err != nil {
		return setError(errorMessage, err)
	}
	return 0
}

//export libtrust_public_key_pem
func libtrust_public_key_pem(id C.uintptr_t, output **C.char, outputLength *C.size_t, errorMessage **C.char) C.int {
	clearError(errorMessage)
	clearBytes(output, outputLength)
	key, err := loadKey(id)
	if err != nil {
		return setError(errorMessage, err)
	}
	block, err := key.PublicKey().PEMBlock()
	if err != nil {
		return setError(errorMessage, err)
	}
	if err := setBytes(output, outputLength, pem.EncodeToMemory(block)); err != nil {
		return setError(errorMessage, err)
	}
	return 0
}

//export libtrust_sign_json
func libtrust_sign_json(id C.uintptr_t, input *C.char, inputLength C.size_t, output **C.char, outputLength *C.size_t, errorMessage **C.char) C.int {
	clearError(errorMessage)
	clearBytes(output, outputLength)
	key, err := loadKey(id)
	if err != nil {
		return setError(errorMessage, err)
	}
	content, err := copyInput(input, inputLength)
	if err != nil {
		return setError(errorMessage, err)
	}
	signature, err := libtrust.NewJSONSignature(content)
	if err == nil {
		err = signature.Sign(key)
	}
	if err != nil {
		return setError(errorMessage, err)
	}
	signedJSON, err := signature.JWS()
	if err != nil {
		return setError(errorMessage, err)
	}
	if err := setBytes(output, outputLength, signedJSON); err != nil {
		return setError(errorMessage, err)
	}
	return 0
}

//export libtrust_verify_jws
func libtrust_verify_jws(input *C.char, inputLength C.size_t, signerKeyID **C.char, signerKeyIDLength *C.size_t, errorMessage **C.char) C.int {
	clearError(errorMessage)
	clearBytes(signerKeyID, signerKeyIDLength)
	content, err := copyInput(input, inputLength)
	if err != nil {
		return setError(errorMessage, err)
	}
	signature, err := libtrust.ParseJWS(content)
	if err != nil {
		return setError(errorMessage, err)
	}
	keys, err := signature.Verify()
	if err != nil {
		return setError(errorMessage, err)
	}
	if len(keys) != 1 {
		return setError(errorMessage, fmt.Errorf("expected one signer, got %d", len(keys)))
	}
	if err := setBytes(signerKeyID, signerKeyIDLength, []byte(keys[0].KeyID())); err != nil {
		return setError(errorMessage, err)
	}
	return 0
}

func main() {}
