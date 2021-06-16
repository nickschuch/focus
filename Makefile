#!/usr/bin/make -f

export CGO_ENABLED=0

PROJECT=github.com/nickschuch/focus

build:
	gox -os='linux darwin' -arch='amd64' -output='bin/focus_{{.OS}}_{{.Arch}}' -ldflags='-extldflags "-static"' $(PROJECT)

.PHONY: *
