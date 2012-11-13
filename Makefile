PROJECT=variant

build:
	GOPATH=$(PWD) go build

test:
	GOPATH=$(PWD) go test

clean:
	rm -rf pkg

godoc:
	GOPATH=$(PWD) godoc -http=:6060
