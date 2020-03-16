GOPATH=$(shell pwd)
SVC_NAME=fin-receiver
GOPACKAGE_DIR=src/$(SVC_NAME)
EXE=$(SVC_NAME)
DOCKER_DIR=docker

.PHONY: vendor compile docker

vendor: $(GOPACKAGE_DIR)/vendor
	GOPATH=$(GOPATH) cd $(GOPACKAGE_DIR) && govendor fetch +missing

$(GOPACKAGE_DIR)/vendor:
	GOPATH=$(GOPATH) cd $(GOPACKAGE_DIR) && govendor init 

#$(EXE):
#	GOPATH=$(GOPATH) cd $(GOPACKAGE_DIR) && go build

$(EXE):
	go build .

build-nc: $(EXE)
	cp $(EXE) $(DOCKER_DIR)/$(EXE)
	make -C $(DOCKER_DIR)/ build-nc

compile: $(EXE)

clean-exe:
	rm -rf $(EXE)
