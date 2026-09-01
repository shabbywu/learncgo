#ifndef LEARNCGO_LIBTRUST_HPP
#define LEARNCGO_LIBTRUST_HPP

#include "libtrust_go.h"

#include <cstddef>
#include <cstdint>
#include <stdexcept>
#include <string>
#include <utility>

namespace learncgo::libtrust {

class Buffer {
 public:
  Buffer() = default;
  Buffer(char* data, std::size_t size) : data_(data), size_(size) {}
  ~Buffer() { reset(); }

  Buffer(const Buffer&) = delete;
  Buffer& operator=(const Buffer&) = delete;

  Buffer(Buffer&& other) noexcept
      : data_(std::exchange(other.data_, nullptr)),
        size_(std::exchange(other.size_, 0)) {}

  Buffer& operator=(Buffer&& other) noexcept {
    if (this != &other) {
      reset();
      data_ = std::exchange(other.data_, nullptr);
      size_ = std::exchange(other.size_, 0);
    }
    return *this;
  }

  std::string string() const { return std::string(data_, size_); }

 private:
  void reset() {
    if (data_ != nullptr) {
      libtrust_buffer_free(data_);
      data_ = nullptr;
      size_ = 0;
    }
  }

  char* data_ = nullptr;
  std::size_t size_ = 0;
};

inline void check(int status, char* error) {
  if (status == 0) {
    if (error != nullptr) {
      libtrust_buffer_free(error);
    }
    return;
  }
  std::string message = error == nullptr ? "libtrust operation failed" : error;
  libtrust_buffer_free(error);
  throw std::runtime_error(message);
}

inline Buffer getBuffer(int status, char* data, std::size_t size, char* error) {
  check(status, error);
  return Buffer(data, size);
}

class Key {
 public:
  static Key generateP256() {
    uintptr_t handle = 0;
    char* error = nullptr;
    check(libtrust_generate_ec_p256(&handle, &error), error);
    return Key(handle);
  }

  ~Key() { reset(); }

  Key(const Key&) = delete;
  Key& operator=(const Key&) = delete;

  Key(Key&& other) noexcept : handle_(std::exchange(other.handle_, 0)) {}

  Key& operator=(Key&& other) noexcept {
    if (this != &other) {
      reset();
      handle_ = std::exchange(other.handle_, 0);
    }
    return *this;
  }

  std::string id() const {
    char* data = nullptr;
    std::size_t size = 0;
    char* error = nullptr;
    return getBuffer(libtrust_key_id(handle_, &data, &size, &error), data, size, error).string();
  }

  std::string publicKeyPEM() const {
    char* data = nullptr;
    std::size_t size = 0;
    char* error = nullptr;
    return getBuffer(libtrust_public_key_pem(handle_, &data, &size, &error), data, size, error).string();
  }

  std::string signJSON(const std::string& json) const {
    char* data = nullptr;
    std::size_t size = 0;
    char* error = nullptr;
    return getBuffer(
               libtrust_sign_json(handle_, const_cast<char*>(json.data()), json.size(), &data, &size, &error),
               data, size, error)
        .string();
  }

 private:
  explicit Key(uintptr_t handle) : handle_(handle) {}

  void reset() {
    if (handle_ != 0) {
      libtrust_key_free(handle_);
      handle_ = 0;
    }
  }

  uintptr_t handle_ = 0;
};

inline std::string verifyJWS(const std::string& jws) {
  char* data = nullptr;
  std::size_t size = 0;
  char* error = nullptr;
  return getBuffer(libtrust_verify_jws(const_cast<char*>(jws.data()), jws.size(), &data, &size, &error),
                   data, size, error)
      .string();
}

}  // namespace learncgo::libtrust

#endif
