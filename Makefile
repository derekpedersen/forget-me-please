build:
	rm -rf bin/forget-me-please
	go build -o bin/forget-me-please

run: build
	./bin/run.sh