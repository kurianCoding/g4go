CC=go
FLAG=-v
OUT=bin

make:
	$(CC) build -o bin/g4g $(FLAG) main.go

test:
	$(CC) test $(FLAG)
