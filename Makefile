IPTB_ROOT ?=$(HOME)/testbed
NODES ?=5
# set env vars ESCROW_PUB_KEYS and the GUARD_PUB_KEYS and the 'make start' will config for all nodes
GUARDPUBKEYS := $(GUARD_PUB_KEYS)
ESCROWPUBKEYS := $(ESCROW_PUB_KEYS)

all: iptb

deps:
	go mod download

iptb: deps
	(cd iptb; go build)
CLEAN += iptb/iptb

btfslocal: deps
	(cd localbtfs/plugin; go build -buildmode=plugin -o ../../build/localbtfs.so)
CLEAN += build/localbtfs.so

install: deps
	(cd iptb; go install)

clean:
	rm ${CLEAN}

start:
	iptb auto -type localbtfs -count $(NODES)
	iptb run -- btfs config --json Addresses.Announce  []
ifdef GUARDPUBKEYS
	iptb run -- btfs config --json Services.GuardPubKeys ['"$(GUARDPUBKEYS)"']
endif
ifdef ESCROWPUBKEYS
	iptb run -- btfs config --json Services.EscrowPubKeys ['"$(ESCROWPUBKEYS)"']
endif
	iptb start

start_dev:
	iptb auto -type localbtfs -count $(NODES)
	iptb run -- btfs config --json Addresses.Announce  []
	iptb run -- btfs config Services.StatusServerDomain 'https://status-dev.btfs.io'
	iptb run -- btfs config Services.EscrowDomain 'https://escrow-dev.btfs.io'
	iptb run -- btfs config Services.GuardDomain 'https://guard-dev.btfs.io'
	iptb run -- btfs config Services.HubDomain 'https://hub-dev.btfs.io'
ifdef GUARDPUBKEYS
	iptb run -- btfs config --json Services.GuardPubKeys ['"$(GUARDPUBKEYS)"']
endif
ifdef ESCROWPUBKEYS
	iptb run -- btfs config --json Services.EscrowPubKeys ['"$(ESCROWPUBKEYS)"']
endif
	iptb run -- btfs config optin
	iptb start

stop:
	iptb stop


.PHONY: all clean ipfslocal p2pdlocal ipfsdocker ipfsbrowser start start_dev stop
