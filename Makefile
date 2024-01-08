MAKEFLAGS += --silent

BIN_DIR=/usr/bin
BIN=tvshow-parse
BUILD_DIR=build-debian
RELEASE_DIR := $(realpath $(CURDIR)/..)

VERSION := $(shell cat VERSION)
PLATFORM := $(shell uname -m)
USER := $(shell stat -c %U Makefile)

ARCH :=
	ifeq ($(PLATFORM),x86_64)
		ARCH = amd64
	endif
	ifeq ($(PLATFORM),aarch64)
		ARCH = arm64
	endif
	ifeq ($(PLATFORM),armv7l)
		ARCH = armhf
	endif
GOARCH :=
	ifeq ($(ARCH),amd64)
		GOARCH = amd64
	endif
	ifeq ($(ARCH),arm64)
		GOARCH = arm64
	endif
	ifeq ($(ARCH),armhf)
		GOARCH = arm
	endif

ifeq ($(GOARCH),)
	$(error Invalid ARCH: $(ARCH))
endif

.PHONY: all debian clean build tidy vendor install uninstall

all: build

debian: clean $(BUILD_DIR)/DEBIAN
	@echo Building package...
	cp $(BIN) $(BUILD_DIR)$(BIN_DIR)
	chmod --quiet 0555 $(BUILD_DIR)/DEBIAN/p* || true
	fakeroot dpkg-deb -b -z9 $(BUILD_DIR) $(RELEASE_DIR)

clean:
	@echo Clean...
	rm -rf $(BUILD_DIR)

$(BUILD_DIR)/DEBIAN: $(BUILD_DIR)
	@echo Prapare package...
	cp -R deb/DEBIAN $(BUILD_DIR)
	$(MAKE) install DESTDIR=$(BUILD_DIR)
	$(eval SIZE := $(shell du -sbk $(BUILD_DIR) | grep -o '[0-9]*'))
	@sed -i "s/==version==/$(VERSION)/g;s/==size==/$(size)/g;s/==architecture==/$(ARCH)/g" "$(BUILD_DIR)/DEBIAN/control"

$(BUILD_DIR):
	mkdir $(BUILD_DIR)

build:
	GOOS=linux GOARCH=$(GOARCH) go build -o $(BIN) .

tidy:
	go mod tidy

vendor: tidy
	go mod vendor

install:
	install -Dm755 $(BIN) $(DESTDIR)$(BIN_DIR)/$(BIN)

uninstall:
	rm -f $(DESTDIR)$(BIN_DIR)/$(BIN)

test:
	go test ./... -race -cover

release:
	git tag v$(VERSION)
	git push origin v$(VERSION)