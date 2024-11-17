GO = gccgo
CFLAGS = -O3 -flto -march=native -funroll-loops -ftree-vectorize
bin = ./sysconf

build:
	$(GO) $(CFLAGS) main.go -o $(bin)

clean:
	rm -rf $(bin)
	$(GO) $(CFLAGS) main.go -o $(bin)

