files=$(wildcard *.go)

all: ico

ico: $(files)
	go build -o ico $(files)

clean:
	rm -rf ico

js: $(files)
	gopherjs build

.PHONY: all clean

