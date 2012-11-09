PROJECT=variant

build:
	GOPATH=$(PWD):$(PWD)/ext go build $(PROJECT)

test:
	GOPATH=$(PWD):$(PWD)/ext go test $(PROJECT)

clean:
	rm -f $(PROJECT)
	rm -rf pkg
	rm -rf ext/pkg
	rm -f TAGS

#setup:
#	GOPATH=$(PWD)/ext go get launchpad.net/gocheck

retag:
	rm -f TAGS
	find $(GOROOT)src/pkg -type f -name '*.go' | xargs ctags -e -a
	find src -type f -name '*.go' | xargs ctags -e -a
	find ext -type f -name '*.go' | xargs ctags -e -a

tags: retag

godoc:
	GOPATH=$(PWD):$(PWD)/ext godoc -http=:6060
