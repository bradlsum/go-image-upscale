all: gscale internal

internal:
	go build ./internal/*

gscale:
	go build -o gscale cmd/gscale/main.go

test:
	go test ./internal/*

clean:
	rm gscale
