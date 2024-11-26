GO = go
bin = gomgr

build:
	$(GO) build

install:
	mv $(bin) /usr/local/bin/$(bin)

clean:
	rm -rf $(bin)
