#include "libtrust.hpp"

#include <iostream>
#include <stdexcept>
#include <string>

int main() {
  try {
    auto key = learncgo::libtrust::Key::generateP256();
    const std::string keyID = key.id();
    if (keyID.empty() || key.publicKeyPEM().find("BEGIN PUBLIC KEY") == std::string::npos) {
      return 1;
    }

    const std::string payload = R"({"lesson":"4.cgolib_exercise","signed":true})";
    const std::string signedJWS = key.signJSON(payload);
    if (learncgo::libtrust::verifyJWS(signedJWS) != keyID) {
      return 1;
    }
    std::cout << "JWS: " << signedJWS << std::endl;

    std::string tamperedJWS = signedJWS;
    const std::string payloadField = "\"payload\": \"";
    const std::size_t payloadStart = tamperedJWS.find(payloadField);
    if (payloadStart == std::string::npos) {
      return 1;
    }
    const std::size_t payloadByte = payloadStart + payloadField.size();
    tamperedJWS[payloadByte] = tamperedJWS[payloadByte] == 'A' ? 'B' : 'A';

    bool tamperingRejected = false;
    try {
      static_cast<void>(learncgo::libtrust::verifyJWS(tamperedJWS));
    } catch (const std::runtime_error&) {
      tamperingRejected = true;
    }
    if (!tamperingRejected) {
      return 1;
    }

    std::cout << "OK 4.cgolib_exercise" << std::endl;
    return 0;
  } catch (const std::exception& error) {
    std::cerr << error.what() << std::endl;
    return 1;
  }
}
