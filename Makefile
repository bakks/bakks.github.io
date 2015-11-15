files=$(wildcard *.go)

all: ico

ico: $(files)
	go build -o ico $(files)

clean:
	rm -rf ico

js: $(files)
	./pkg/darwin_amd64/github.com/gopherjs/bin/gopherjs build

.PHONY: all clean

