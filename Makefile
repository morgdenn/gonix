# Project-specific variables
BINARY_NAME = gonix

# Build commands
.PHONY: all ls mv rm clean

all: ls mv rm

ls:
	go build -o bin/ls cmd/ls/ls.go

mv:
	go build -o bin/mv cmd/mv/mv.go

rm:
	go build -o bin/rm cmd/rm/rm.go

# Clean up build artifacts
clean:
	rm -rf bin/*
