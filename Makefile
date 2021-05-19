CC = `which go`

SRC = $(shell (find ./src \( -name "*.go" ! -name "*_test.go" \)))
EXE = transactions

all:
	$(CC) run $(SRC)

build:
	$(CC) build -o $(EXE) $(SRC)

docker-build:
	docker build -t pismo .

docker-run: docker-build
	docker run -d -p 8090:8090 5432:5432 pismo

vars:
	@echo COMPILER........$(CC)
	@echo SOURCE FILES...$(SRC)