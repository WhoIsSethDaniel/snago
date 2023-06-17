all:
	CGO_ENABLED=0 go build -tags netgo

install:
	CGO_ENABLED=0 go install -tags netgo

clean:
	rm -f snago
