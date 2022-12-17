#!/usr/bin/env -S gmake -f

GOBIN := minica2
GOFLAGS := -ldflags '-X main.date=$(shell date +%Y-%m-%d)'
VERSION := 1.0.1

ifeq ($(PREFIX),)
    PREFIX := /opt
endif

.DEFAULT:
.PHONY:
all: $(GOBIN)

$(GOBIN).go:
	[ -f $@ ] && rm $@ || true;\
	go build $(GOFLAGS)

$(GOBIN): $(GOBIN).go
	mv $< $@

README.md:
	./$@.sh > $@

.PHONY:
install: $(GOBIN)
	install -d $(DESTDIR)$(PREFIX)/bin/
	install -m 777 $< $(DESTDIR)$(PREFIX)/bin/

.PHONY:
dist:
	rm -rf dist && mkdir dist && $(MAKE) -j$(shell nproc) bdist && \
		find dist -executable -type f | parallel tar --zstd -c -v -f {}.tar.zstd {}

.PHONY:
bdist:
	GOOS=darwin GOARCH=amd64 go build -o dist/$(GOBIN).v$(VERSION).x86_64.darwin $(GOFLAGS)
	GOOS=linux GOARCH=amd64 go build -o dist/$(GOBIN).v$(VERSION).x86_64.linux $(GOFLAGS)
	GOOS=windows GOARCH=amd64 go build -o dist/$(GOBIN).v$(VERSION).x86_64.windows.exe $(GOFLAGS)
	GOOS=linux GOARCH=386 go build -o dist/$(GOBIN).v$(VERSION).x86.linux $(GOFLAGS)
	GOOS=windows GOARCH=386 go build -o dist/$(GOBIN).v$(VERSION).x86.windows.exe $(GOFLAGS)
	GOOS=linux GOARCH=arm64 go build -o dist/$(GOBIN).v$(VERSION).arm64.linux $(GOFLAGS)
	GOOS=darwin GOARCH=arm64 go build -o dist/$(GOBIN).v$(VERSION).arm64.darwin $(GOFLAGS)
	GOOS=linux GOARCH=arm go build -o dist/$(GOBIN).v$(VERSION).arm.linux $(GOFLAGS)
