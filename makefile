# Go パラメータ
GOCMD = go
GOFMT = goimports
GOBUILD = $(GOCMD) build
GOCLEAN	= $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOGENERATE = ${GOCMD} generate
# ターゲットパラメータ
GOFILES	= $(shell find . -name "*.go")
BINARY_NAME = docbase

# タスク
.PHONY: fmt
fmt:
	$(GOFMT) -w ${GOFILES}

.PHONY: build
build:
	${GOGENERATE}
	$(GOBUILD) -o $(BINARY_NAME) main.go

.PHONY: clean
clean:
	$(GOCLEANN)

.PHONY: install
install:
	make build
	mv $(BINARY_NAME) /usr/local/bin/
