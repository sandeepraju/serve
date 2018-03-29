.PHONY: all clean run build release

all:
	clean

clean:
	rm -f ./serve
	rm -f ./serve.zip

run:
	go run serve.go

build:
	go build .

release:
	echo "zipping..."
	zip serve.zip serve
	openssl sha -sha256 serve.zip
