BIN := terraform-provider-math
ENTRYPOINT := cmd/$(BIN)/main.go
DIST_DIR := dist

ACCTEST_COUNT ?= 1
ACCTEST_PARALLELISM ?= 20
ACCTEST_TIMEOUT ?= 240m

testacc:
	TF_ACC=1 \
	go test ./... -v $(TESTARGS) -count $(ACCTEST_COUNT) -timeout $(ACCTEST_TIMEOUT) -parallel=$(ACCTEST_PARALLELISM)

build:
	rm -rf $(DIST_DIR)
	mkdir $(DIST_DIR)
	go build -o $(DIST_DIR)/$(BIN) $(ENTRYPOINT)
