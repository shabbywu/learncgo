.DEFAULT_GOAL:=help
.PHONY: build
build: _build ## 构建所有案例

.PHONY: _build
ifeq ($(OS),Windows_NT)
_build:
	mkdir -p build
	cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_CONFIGURATION_TYPES="Release" -B build
	cmake --build build --config Release
else
_build:
	mkdir -p build
	cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_CONFIGURATION_TYPES="Release" -B build
	cmake --build build --config Release -v
endif


.PHONY: help
help:  ## Show this message
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)