MAKEFLAGS += --silent

BIN_DIR=usr/bin
BIN=tvshow-parse
BUILD_DIR=build
RELEASE_DIR=$(BUILD_DIR)/release
TMP_DIR=$(BUILD_DIR)/tmp
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

package: clean prepare cp $(BIN) control
	@echo Building package...
	mv $(BIN) $(TMP_DIR)/$(BIN_DIR)/
	chmod --quiet 0555 $(TMP_DIR)/DEBIAN/p* || true
	fakeroot dpkg-deb -b -z9 $(TMP_DIR) $(RELEASE_DIR)

install: $(BIN)
	install -D -g root -o root $(BIN) $(DESTDIR)/usr/bin
	rm $(BIN)

$(BIN): test
	go mod tidy
	go mod vendor > /dev/null 2>&1
	chown -R $(USER).$(USER) vendor
	GOOS=linux GOARCH=$(GOARCH) go build -o $(BIN) main.go

test:
	go clean -testcache
	go test ./...

clean:
	rm -rf $(TMP_DIR) $(RELEASE_DIR)

prepare:
	@echo Prepare...
	mkdir -p $(TMP_DIR)/$(BIN_DIR) $(RELEASE_DIR)

cp:
	cp -R deb/* $(TMP_DIR)

control:
	$(eval size=$(shell du -sbk $(TMP_DIR)/ | grep -o '[0-9]*'))
	@sed -i "s/==version==/$(VERSION)/g;s/==size==/$(size)/g;s/==architecture==/$(ARCH)/g" "$(TMP_DIR)/DEBIAN/control"

