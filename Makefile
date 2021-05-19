CC = `which go`

SRC = $(shell (find ./src -maxdepth 1 \( -name "*.go" ! -name "*_test.go" \)))
EXE = transactions

all:
	$(CC) run $(SRC)

build:
	$(CC) build -o $(EXE) $(SRC)

docker-build:
	docker build -t pismo .

docker-run: docker-build
	docker run -d --network=host pismo

vars:
	@echo COMPILER........$(CC)
	@echo SOURCE FILES...$(SRC)