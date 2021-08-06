BIN := terraform-provider-math
ENTRYPOINT := cmd/$(BIN)/main.go

ACCTEST_COUNT ?= 1
ACCTEST_PARALLELISM ?= 20
ACCTEST_TIMEOUT ?= 240m

testacc:
	TF_ACC=1 \
	go test ./... -v $(TESTARGS) -count $(ACCTEST_COUNT) -timeout $(ACCTEST_TIMEOUT) -parallel=$(ACCTEST_PARALLELISM)

build:
	rm -f $(BIN)
	go build -o $(BIN) $(ENTRYPOINT)

