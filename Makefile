GO = gccgo 
CFLAGS = -O3 -flto -march=native -funroll-loops -ftree-vectorize -fno-exceptions
CFLAGS_EXT = -fomit-frame-pointer -fno-stack-protector -g0
bin = ./gopherconf

debug:
	$(GO) $(CFLAGS) -g main.go backend.go -o $(bin)

build:
	$(GO) $(CFLAGS) $(CFLAGS_EXT) main.go backend.go -o $(bin)

clean:
	rm -rf $(bin)
	$(GO) $(CFLAGS) $(CFLAGS_EXT) main.go backend.go -o $(bin)