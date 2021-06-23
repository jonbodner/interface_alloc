bench:
	go test -bench=. -benchmem

bench_noinline:
	go test -bench=. -benchmem -gcflags=-l
