SRC?=main.go
ARGS?=

test:
	@go run $(SRC) $(ARGS)

deps:
	@go get
