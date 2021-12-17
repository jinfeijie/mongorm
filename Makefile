
all:mod

clean:
	rm -rf vendor

mod:
	go mod tidy -v
	go mod vendor -v