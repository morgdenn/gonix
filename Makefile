# Project-specific variables
BINARY_DIR = bin

# Build commands
.PHONY: all ls mv rm clean

all: ls ll mv rm which

ls:
	go build -o $(BINARY_DIR)/ls.exe cmd/ls/ls.go

ll:
	go build -o $(BINARY_DIR)/ll.exe cmd/ll/ll.go

mv:
	go build -o $(BINARY_DIR)/mv.exe cmd/mv/mv.go

rm:
	go build -o $(BINARY_DIR)/rm.exe cmd/rm/rm.go

which:
	go build -o $(BINARY_DIR)/which.exe cmd/which/which.go

# Clean up build artifacts
clean:
	@if exist $(BINARY_DIR) (rmdir /S /Q $(BINARY_DIR))
