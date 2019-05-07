GOPATH=$(shell pwd)
SVC_NAME=fin-receiver
GOPACKAGE_DIR=src/$(SVC_NAME)
EXE=$(GOPACKAGE_DIR)/$(SVC_NAME)

.PHONY: vendor compile

vendor: $(GOPACKAGE_DIR)/vendor
	GOPATH=$(GOPATH) cd $(GOPACKAGE_DIR) && govendor fetch +missing

$(GOPACKAGE_DIR)/vendor:
	GOPATH=$(GOPATH) cd $(GOPACKAGE_DIR) && govendor init 

$(EXE):
	GOPATH=$(GOPATH) cd $(GOPACKAGE_DIR) && go build

compile: $(EXE)

clean-exe:
	rm -rf $(EXE)
