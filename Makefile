PYTHON_EXECUTABLE := $(shell which python)


.DEFAULT_GOAL:=help
.PHONY: build
build: ## 构建所有案例
	mkdir -p build
	cmake --preset=default -DPython_EXECUTABLE=${PYTHON_EXECUTABLE} -DPYTHON_EXECUTABLE=${PYTHON_EXECUTABLE}
	cmake --build build --config Release -v


.PHONY: help
help:  ## Show this message
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
