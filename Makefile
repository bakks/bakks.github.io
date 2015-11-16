files=$(wildcard *.go)

all: js

ico: $(files)
	go build -o ico $(files)

clean:
	rm -rf ico icosahedron *.log *.js *.js.map

closure/compiler.jar:
	mkdir -p closure
	curl https://dl.google.com/closure-compiler/compiler-latest.zip > closure/closure.zip
	cd closure && unzip closure.zip

js: $(files) closure/compiler.jar
	gopherjs build
	java -jar closure/compiler.jar --js icosahedron.js --js_output_file ico-min.js

.PHONY: all clean

