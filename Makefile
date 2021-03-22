all:
	CGO_ENABLED=0 go build -tags netgo -o snago

clean:
	rm -f snago
