MARKDOWN_FILE?=README.md
SRC?=main.go
ARGS?=

test:
	@go run $(SRC) $(ARGS)

deps:
	@go get

toc:
	markdown-toc --indent "    " -i $(MARKDOWN_FILE)
