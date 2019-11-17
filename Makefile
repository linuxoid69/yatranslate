.PHONY: build install uninstall

# .DEFAULT_GOAL := build
build:
	set GOPATH=%cd%
	go build

install:
	install -v -m 555 yatranslate /usr/local/bin/yatranslate

uninstall:
	rm /usr/local/bin/yatranslate
