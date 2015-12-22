BASE := $(shell pwd)
export GOPATH=$(BASE)
#export GOBIN=$(BASE)/bin
export GO15VENDOREXPERIMENT=1

files=$(wildcard *.go)

all: js

clean:
	rm -rf target

vendor/closure/compiler.jar:
	mkdir -p vendor/closure
	curl https://dl.google.com/closure-compiler/compiler-latest.zip > vendor/closure/closure.zip
	cd vendor/closure && unzip closure.zip

js: vendor/closure/compiler.jar
	mkdir -p target
	gopherjs build -o target/icosahedron.js
	java -jar vendor/closure/compiler.jar --js target/icosahedron.js --js_output_file target/ico-min.js

.PHONY: all clean

