IPTB_ROOT ?=$(HOME)/testbed
NODES ?=5

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
	iptb run -- btfs config profile apply storage-host
	iptb start
	sleep 10
	iptb logs > iptb_logs.txt
	go run standalone/rewrite_config.go

start_dev:
	iptb auto -type localbtfs -count $(NODES)
	iptb run -- btfs config profile apply storage-host-dev
	iptb run -- btfs config optin
	iptb start
	sleep 10
	iptb logs > iptb_logs.txt
	go run standalone/rewrite_config.go

stop:
	iptb stop


.PHONY: all clean ipfslocal p2pdlocal ipfsdocker ipfsbrowser start start_dev stop
