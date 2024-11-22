GO = go
bin = ./gopherconf

build:
	$(GO) build

install:
	mv $(bin) /usr/local/bin/$(bin)

clean:
	rm -rf $(bin)
