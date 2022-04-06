TEST    ?= $$(go list ./... | grep -v 'vendor')
NAME    := zookeeper
BINARY  := terraform-provider-${NAME}
GO      ?= go

# Acceptance tests zookeeper
ZOOKEEPER_HOST ?= localhost
ZOOKEEPER_PORT ?= 2181

build:
	go build -o $(BINARY)

test:
	$(GO) test $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 $(GO) test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	ZOOKEEPER_HOST=$(ZOOKEEPER_HOST) ZOOKEEPER_PORT=$(ZOOKEEPER_PORT) TF_ACC=1 $(GO) test $(TEST) -v $(TESTARGS) -timeout 120m

vet:
	$(GO) vet