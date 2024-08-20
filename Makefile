# Project-specific variables
BINARY_DIR = bin

# Build commands
.PHONY: all ls mv rm clean

all: ls mv rm

ls:
	go build -o $(BINARY_DIR)/ls.exe cmd/ls/ls.go

mv:
	go build -o $(BINARY_DIR)/mv.exe cmd/mv/mv.go

rm:
	go build -o $(BINARY_DIR)/rm.exe cmd/rm/rm.go

# Clean up build artifacts
clean:
	@if exist $(BINARY_DIR) (rmdir /S /Q $(BINARY_DIR))
