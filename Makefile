GO = gccgo
bin = ./sysconf

build:
	$(GO) -flto main.go -o $(bin)

clean:
	rm -rf $(bin)

