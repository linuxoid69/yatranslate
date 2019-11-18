.PHONY: build install uninstall

# .DEFAULT_GOAL := build
build:
	set GOPATH=%cd%
	GO111MODULE="on" go mod vendor
	go build

install:
	install -v -m 555 yatranslate /usr/local/bin/yatranslate

uninstall:
	rm /usr/local/bin/yatranslate
