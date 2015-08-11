TAG ?= v1

all:
	# noop
release_osx:
	go build github.com/xlab/gitio/cmd/gitio
	zip -9 gitio_$(TAG)_osx.zip gitio
	rm gitio
release_linux:
	gobldock github.com/xlab/gitio/cmd/gitio
	zip -9 gitio_$(TAG)_linux.zip gitio
	rm gitio
release: release_osx release_linux
build:
	go build github.com/xlab/gitio/cmd/gitio
clean:
	rm -f gitio
	rm -f gitio_$(TAG)_linux.zip gitio_$(TAG)_osx.zip
